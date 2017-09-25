package mysql

import (
	"database/sql"

	"sync"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql/iface"
	"github.com/spf13/viper"
)

type contextKey struct{}

type bus struct {
	id   int64
	name string
}

var (
	db   *sql.DB
	once sync.Once
)

func GetConnection() iface.DBIface {
	once.Do(func() {
		initConnection()
	})
	return db
}

func CloseConnection() {
	db.Close()
}

func initConnection() {
	var err error
	conf := viper.GetStringMapString("mysql")
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		conf["user"], conf["password"], conf["database"]))
	if err != nil {
		glog.Error(err)
	}
}
