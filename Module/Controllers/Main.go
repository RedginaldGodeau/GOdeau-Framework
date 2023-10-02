package Controllers

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

var Controller map[string]Handler

func InitController(name string, h Handler) {
	Controller[name] = h
}
