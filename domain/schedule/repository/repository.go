package repository

import (
	"database/sql"

	"path/filepath"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/model"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path, err := filepath.Abs("./domain/schedule/repository/queries.sql")
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	queries = goyesql.MustParseFile(path)
}

func FetchManyByBus(busID int64) ([]*model.Schedule, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["by-bus"], busID)
	if err != nil {
		return nil, err
	}
	return parseRows(rows)
}

func FetchManyByBusDayType(busID int64, dayTypeID int64) ([]*model.Schedule, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["by-bus-daytype"], busID, dayTypeID)
	if err != nil {
		return nil, err
	}
	return parseRows(rows)
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
