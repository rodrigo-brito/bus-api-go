package repository

import (
	"database/sql"

	"path/filepath"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/bus/model"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/repository"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path, err := filepath.Abs("./domain/bus/repository/queries.sql")
	if err != nil {
		glog.Error(err)
		panic(err)
	}
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

func Get(ID int64, injectSchedule bool) (*model.Bus, error) {
	conn := mysql.GetConnection()
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

	return nil, nil
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
