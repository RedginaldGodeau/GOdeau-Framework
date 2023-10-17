package controller

import (
	"GOdeau/godeau/Controllers"
	"fmt"
	"net/http"
)

var PingGet = Controllers.Handler(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word")
})
