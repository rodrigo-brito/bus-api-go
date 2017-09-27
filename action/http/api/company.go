package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/rodrigo-brito/bus-api-go/domain/company/repository"
)

func CompanyHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	company, err := repository.GetAll()
	if err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		glog.Error(err)
	}
}

func GetCompanyHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		glog.Error(err)
		return
	}
	company, err := repository.Get(ID, false)
	if err != nil {
		glog.Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		glog.Error(err)
	}
}
