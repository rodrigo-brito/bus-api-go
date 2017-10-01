package api

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/rodrigo-brito/bus-api-go/lib/mail"
)

func ContactHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	email := r.FormValue("email")
	name := r.FormValue("name")
	message := r.FormValue("message")

	go func() {
		err := mail.SendMessage(email, "Contato - Onibus Sabara (WEB)", name, message)
		if err != nil {
			glog.Error("api/contact: ", err)
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
}
