package api

import (
	"net/http"

	"encoding/json"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/rodrigo-brito/bus-api-go/lib/mail"
)

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func ContactHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := new(Message)
	err = json.Unmarshal(b, message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	go func() {
		err := mail.SendMessage(message.Email, "Contato - Onibus Sabara (WEB)",
			message.Name, message.Message)
		if err != nil {
			glog.Error("api/contact: ", err)
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json; charset=utf-8")
}
