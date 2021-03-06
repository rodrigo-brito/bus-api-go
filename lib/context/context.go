package context

import (
	"github.com/rodrigo-brito/bus-api-go/lib/memcached"
	"github.com/rodrigo-brito/bus-api-go/lib/mysql"
	"golang.org/x/net/context"
)

func DefaultContext(BPC bool) context.Context {
	ctx := memcached.NewContext(context.Background(), BPC)
	ctx = mysql.NewContext(ctx)
	return ctx
}
