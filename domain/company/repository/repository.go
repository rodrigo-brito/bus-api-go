package repository

import (
	"database/sql"

	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/bus/repository"
	"github.com/rodrigo-brito/bus-api-go/domain/company/model"
	"github.com/rodrigo-brito/bus-api-go/lib/environment"
	"github.com/rodrigo-brito/bus-api-go/lib/memcached"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path := environment.AbsPath("domain/company/repository/queries.sql")
	queries = goyesql.MustParseFile(path)
}

func getCompanyCacheKey(ID int64) string {
	return fmt.Sprintf("company-%d", ID)
}

func GetAll() ([]*model.Company, error) {
	conn := mysql.GetConnection()
	var companies []*model.Company

	err := memcached.GetSet("all-companies", &companies, func() (interface{}, error) {
		rows, err := conn.Query(queries["all"])
		if err != nil {
			return nil, err
		}
		return parseRows(rows)
	}, time.Hour*48)

	return companies, err
}

func Get(ID int64, injectBus bool) (*model.Company, error) {
	conn := mysql.GetConnection()

	company := new(model.Company)
	key := getCompanyCacheKey(ID)

	err := memcached.GetSet(key, company, func() (interface{}, error) {
		rows, err := conn.Query(queries["by-id"], ID)
		if err != nil {
			return nil, err
		}

		result, err := parseRows(rows)
		if err != nil {
			return nil, err
		}
		return result[0], nil
	}, time.Hour*48)

	if err != nil {
		return nil, err
	}

	if !company.IsEmpty() && injectBus {
		company.Bus, err = repository.GetByCompany(ID)
		if err != nil {
			glog.Error(err)
		}
	}

	return company, err
}

func parseRows(rows *sql.Rows) (result []*model.Company, err error) {
	for rows.Next() {
		company := new(model.Company)
		err := rows.Scan(&company.ID, &company.Name, &company.ImageURL, &company.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, company)
	}
	return
}
