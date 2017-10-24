package repository

import (
	"database/sql"

	"fmt"

	"time"

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

func GetAll() ([]*model.Bus, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["all"])
	if err != nil {
		return nil, err
	}
	return parseRows(rows)
}

func getBusCacheKey(ID int64) string {
	return fmt.Sprintf("bus-%d", ID)
}

func Get(ID int64, injectSchedule bool) (*model.Bus, error) {
	conn := mysql.GetConnection()

	bus := new(model.Bus)
	key := getBusCacheKey(ID)
	hit, err := memcached.Get(key, bus)
	if err != nil {
		return nil, err
	}
	if hit {
		fmt.Println("From cache!")
		return bus, nil
	}

	fmt.Println("From DB")
	rows, err := conn.Query(queries["by-id"], ID)
	if err != nil {
		return nil, err
	}

	var result []*model.Bus

	if result, err = parseRows(rows); err != nil {
		return nil, err
	} else if len(result) > 0 {
		bus := result[0]
		if injectSchedule {
			injectSchedules(bus)
		}
		return bus, nil
	}

	memcached.Set(key, bus, 48*time.Hour)
	return nil, nil
}

func GetByCompany(companyID int64) ([]*model.Bus, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["by-company"], companyID)
	if err != nil {
		return nil, err
	}
	return parseRows(rows)
}

func injectSchedules(bus *model.Bus) {
	schedules, err := repository.FetchManyByBus(bus.ID)
	if err != nil {
		glog.Error(err)
	}
	bus.Schedules = schedules
}

func parseRows(rows *sql.Rows) (result []*model.Bus, err error) {
	for rows.Next() {
		bus := new(model.Bus)
		err := rows.Scan(&bus.ID, &bus.Number, &bus.Name, &bus.Fare)
		if err != nil {
			return nil, err
		}
		result = append(result, bus)
	}
	return
}
