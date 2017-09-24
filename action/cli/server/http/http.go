package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rodrigo-brito/bus-api-go/action/http/api"
)

func InjectAPIRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/status", api.StatusHandle)
	router.GET("/api/v1/bus", api.BusHandle)
	router.GET("/api/v1/bus/:id", api.GetBusHandle)
	router.GET("/api/v1/bus/:id/schedule", api.GetBusScheduleHandle)
	router.GET("/api/v1/bus/:id/schedule/daytype", api.GetDayTypeScheduleHandle)
	return router
}
