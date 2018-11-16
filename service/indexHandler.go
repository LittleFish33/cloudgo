package service

import (
	"net/http"
	"time"

	"github.com/unrolled/render"
)

func indexHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			Name string `json:"name"`
			Time string `json:"time"`
		}{Name: "LittleFish", Time: time.Now().Format("2006-01-02 15:04:05")})
	}
}
