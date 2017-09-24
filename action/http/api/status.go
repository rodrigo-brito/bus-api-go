package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func StatusHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OK")
}
