package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	dr "github.com/rodrigo-brito/bus-api-go/domain/daytype/repository"
)

func GetDayTypeScheduleHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	busID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		glog.Error(err)
		return
	}
	bus, err := dr.GetByBus(busID, true)
	if err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(bus); err != nil {
		glog.Error(err)
	}
}
