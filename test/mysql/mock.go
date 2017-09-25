package mysql

import (
	"github.com/bouk/monkey"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql/iface"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func GetDatabaseMock() (sqlmock.Sqlmock, error) {
	mysql.GetConnection()
	db, mock, err := sqlmock.New()
	monkey.Patch(mysql.GetConnection, func() iface.DBIface {
		return db
	})
	return mock, err
}

func UnmockDatabase() {
	monkey.Unpatch(mysql.GetConnection())
}
