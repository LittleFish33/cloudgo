package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
