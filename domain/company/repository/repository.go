package repository

import (
	"database/sql"

	"github.com/golang/glog"
	"github.com/nleof/goyesql"
	"github.com/rodrigo-brito/bus-api-go/domain/bus/repository"
	"github.com/rodrigo-brito/bus-api-go/domain/company/model"
	"github.com/rodrigo-brito/bus-api-go/lib/environment"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
)

var queries goyesql.Queries

func init() {
	path := environment.AbsPath("domain/company/repository/queries.sql")
	queries = goyesql.MustParseFile(path)
}

func GetAll() ([]*model.Company, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["all"])
	if err != nil {
		return nil, err
	}
	return parseRows(rows)
}

func Get(ID int64, injectBus bool) (*model.Company, error) {
	conn := mysql.GetConnection()
	rows, err := conn.Query(queries["by-id"], ID)
	if err != nil {
		return nil, err
	}

	result, err := parseRows(rows)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		if injectBus {
			bus, err := repository.GetByCompany(ID)
			if err != nil {
				glog.Error(err)
			}
			result[0].Bus = bus
		}
		return result[0], nil
	}

	return nil, nil
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
