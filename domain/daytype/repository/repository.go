package repository

import (
	"database/sql"

	"path/filepath"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/daytype/model"
	"github.com/rodrigo-brito/bus-api-go/domain/schedule/repository"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path, err := filepath.Abs("./domain/daytype/repository/queries.sql")
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	queries = goyesql.MustParseFile(path)
}

func GetByBus(busID int64, injectSchedule bool) ([]*model.DayType, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["by-bus"], busID)
	if err != nil {
		return nil, err
	}

	dayTypes, err := parseRows(rows)
	if err != nil {
		return nil, err
	}

	if injectSchedule {
		injectSchedules(dayTypes, busID)
	}

	return dayTypes, nil
}

func injectSchedules(dayTypes []*model.DayType, busID int64) {
	for _, dayType := range dayTypes {
		schedules, err := repository.FetchManyByBusDayType(busID, dayType.ID)
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
