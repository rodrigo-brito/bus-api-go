package mysql

import (
	"database/sql"

	"sync"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
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

func GetConnection() *sql.DB {
	once.Do(func() {
		initConnection()
	})
	return db
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
