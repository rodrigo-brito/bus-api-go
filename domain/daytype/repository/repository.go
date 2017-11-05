package repository

import (
	"database/sql"

	"fmt"
	"time"

	"context"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/daytype/model"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/repository"
	"github.com/rodrigo-brito/bus-api-go/lib/environment"
	"github.com/rodrigo-brito/bus-api-go/lib/memcached"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path := environment.AbsPath("domain/daytype/repository/queries.sql")
	queries = goyesql.MustParseFile(path)
}

func getDayTyppeBusCacheKey(busID int64) string {
	return fmt.Sprintf("daytype-bus-%d", busID)
}

func GetByBus(ctx context.Context, busID int64, injectSchedule bool) ([]*model.DayType, error) {
	db := mysql.FromContext(ctx)
	cache := memcached.FromContext(ctx)

	var dayTypes []*model.DayType
	key := getDayTyppeBusCacheKey(busID)

	err := cache.GetSet(key, &dayTypes, func() (interface{}, error) {
		rows, err := db.Query(queries["by-bus"], busID)
		if err != nil {
			return nil, err
		}
		res, err := parseRows(rows)
		return &res, err
	}, time.Hour*48)
	if err != nil {
		return nil, err
	}

	if injectSchedule {
		injectSchedules(ctx, dayTypes, busID)
	}

	return dayTypes, nil
}

func injectSchedules(ctx context.Context, dayTypes []*model.DayType, busID int64) {
	for _, dayType := range dayTypes {
		schedules, err := repository.FetchManyByBusDayType(ctx, busID, dayType.ID)
		if err != nil {
			glog.Error(err)
		}
		dayType.Schedules = schedules
	}
	return
}

func parseRows(rows *sql.Rows) (result []*model.DayType, err error) {
	for rows.Next() {
		dayType := new(model.DayType)
		err := rows.Scan(&dayType.ID, &dayType.Name, &dayType.Active)
		if err != nil {
			return nil, err
		}
		result = append(result, dayType)
	}
	return
}
