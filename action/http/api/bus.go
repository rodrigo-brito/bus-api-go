package api

import (
	"net/http"

	"encoding/json"

	"strconv"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	br "github.com/rodrigo-brito/bus-api-go/domain/bus/repository"
)

func BusHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bus, err := br.GetAll(r.Context())
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

func GetBusHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		glog.Error(err)
		return
	}
	bus, err := br.Get(r.Context(), ID, false)
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

func GetBusScheduleHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		glog.Error(err)
		return
	}
	bus, err := br.Get(r.Context(), ID, true)
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

func BusByCompanyHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		glog.Error(err)
		return
	}
	bus, err := br.GetByCompany(r.Context(), ID)
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
