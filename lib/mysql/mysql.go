package mysql

import (
	"database/sql"

	"sync"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql/iface"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

type contextKey struct{}

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
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf["user"], conf["password"], conf["address"], conf["port"], conf["database"]))
	if err != nil {
		glog.Error(err)
	}
	if err := db.Ping(); err != nil {
		glog.Error(err)
	}
}

func NewContext(parent context.Context) context.Context {
	return context.WithValue(parent, contextKey{}, GetConnection())
}

func FromContext(ctx context.Context) iface.DBIface {
	return ctx.Value(contextKey{}).(iface.DBIface)
}
