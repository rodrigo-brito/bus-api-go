package repository

import (
	"database/sql"

	"fmt"

	"time"

	"context"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/bus/model"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/repository"
	"github.com/rodrigo-brito/bus-api-go/lib/environment"
	"github.com/rodrigo-brito/bus-api-go/lib/memcached"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path := environment.AbsPath("domain/bus/repository/queries.sql")
	queries = goyesql.MustParseFile(path)
}

func GetAll(ctx context.Context) ([]*model.Bus, error) {
	db := mysql.FromContext(ctx)

	rows, err := db.Query(queries["all"])
	if err != nil {
		return nil, err
	}
	return parseRows(rows)
}

func getBusCacheKey(ID int64) string {
	return fmt.Sprintf("bus-%d", ID)
}

func getBusCompanyCacheKey(companyID int64) string {
	return fmt.Sprintf("bus-company-%d", companyID)
}

func Get(ctx context.Context, ID int64, injectSchedule bool) (*model.Bus, error) {
	db := mysql.FromContext(ctx)
	cache := memcached.FromContext(ctx)

	bus := new(model.Bus)
	key := getBusCacheKey(ID)

	err := cache.GetSet(key, bus, func() (interface{}, error) {
		rows, err := db.Query(queries["by-id"], ID)
		if err != nil {
			return nil, err
		}

		var result []*model.Bus

		if result, err = parseRows(rows); err != nil {
			return nil, err
		} else if len(result) > 0 {
			bus = result[0]
		}
		return bus, nil
	}, 48*time.Hour)

	if err != nil {
		return nil, err
	}

	if !bus.IsEmpty() && injectSchedule {
		injectSchedules(ctx, bus)
	}

	return bus, nil
}

func GetByCompany(ctx context.Context, companyID int64) ([]*model.Bus, error) {
	db := mysql.FromContext(ctx)
	cache := memcached.FromContext(ctx)

	var bus []*model.Bus
	key := getBusCompanyCacheKey(companyID)
	err := cache.GetSet(key, &bus, func() (interface{}, error) {
		rows, err := db.Query(queries["by-company"], companyID)
		if err != nil {
			return nil, err
		}
		return parseRows(rows)
	}, time.Hour*48)

	return bus, err
}

func injectSchedules(ctx context.Context, bus *model.Bus) {
	schedules, err := repository.FetchManyByBus(ctx, bus.ID)
	if err != nil {
		glog.Error(err)
	}
	bus.Schedules = schedules
}

func parseRows(rows *sql.Rows) (result []*model.Bus, err error) {
	for rows.Next() {
		bus := new(model.Bus)
		err := rows.Scan(&bus.ID, &bus.Number, &bus.Name, &bus.Fare, &bus.LastUpdate)
		if err != nil {
			return nil, err
		}
		result = append(result, bus)
	}
	return
}
