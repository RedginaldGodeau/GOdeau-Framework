package Kernel

import (
	"GOdeau/config/route"
	"net/http"
)

type Kernel struct {
}

func ListenAndServe() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func loader() Kernel {
	route.Init()

	return Kernel{}
}

func Run() {
	loader()
	ListenAndServe()
}
