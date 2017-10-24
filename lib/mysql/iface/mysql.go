package iface

import "database/sql"

type DBIface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}
