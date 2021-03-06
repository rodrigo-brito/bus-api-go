package repository

import (
	"database/sql"

	"fmt"

	"time"

	"context"

	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/model"
	"github.com/rodrigo-brito/bus-api-go/lib/environment"
	"github.com/rodrigo-brito/bus-api-go/lib/memcached"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path := environment.AbsPath("domain/schedule/repository/queries.sql")
	queries = goyesql.MustParseFile(path)
}

func getScheduleBusCacheKey(busID int64) string {
	return fmt.Sprintf("bus-schedule-%d", busID)
}

func getScheduleBusDayTypeCacheKey(busID int64, dayType int64) string {
	return fmt.Sprintf("bus-daytype-schedule-%d-%d", busID, dayType)
}

func FetchManyByBus(ctx context.Context, busID int64) ([]*model.Schedule, error) {
	db := mysql.FromContext(ctx)
	cache := memcached.FromContext(ctx)

	var schedules []*model.Schedule
	key := getScheduleBusCacheKey(busID)

	err := cache.GetSet(key, &schedules, func() (interface{}, error) {
		rows, err := db.Query(queries["by-bus"], busID)
		if err != nil {
			return nil, err
		}
		res, err := parseRows(rows)
		return &res, err
	}, time.Hour*48)

	return schedules, err
}

func FetchManyByBusDayType(ctx context.Context, busID int64, dayTypeID int64) ([]*model.Schedule, error) {
	db := mysql.FromContext(ctx)
	cache := memcached.FromContext(ctx)

	var schedules []*model.Schedule
	key := getScheduleBusDayTypeCacheKey(busID, dayTypeID)

	err := cache.GetSet(key, &schedules, func() (interface{}, error) {
		rows, err := db.Query(queries["by-bus-daytype"], busID, dayTypeID)
		if err != nil {
			return nil, err
		}
		res, err := parseRows(rows)
		return &res, err
	}, time.Hour*48)

	return schedules, err
}

func parseRows(rows *sql.Rows) (result []*model.Schedule, err error) {
	for rows.Next() {
		schedule := new(model.Schedule)
		err = rows.Scan(&schedule.ID, &schedule.Origin, &schedule.Destiny, &schedule.Observation, &schedule.Time)
		if err != nil {
			return nil, err
		}
		result = append(result, schedule)
	}
	return
}
