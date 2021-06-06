package controllers

import (
	"net/http"

	"github.io/taserbeat/golang-todo-app/modules/setting"
)

func StartMainServer() (err error) {
	files := http.FileServer(http.Dir(setting.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+setting.Config.Port, nil)
}
