package Controllers

import (
	"GOdeau/godeau/Controllers/Methods"
	"fmt"
	"net/http"
)

type controller struct {
	name    string
	pattern string
	methods []Methods.Method
	handler Handler
}
type Handler func(w http.ResponseWriter, r *http.Request)

func New(name string, pattern string, methods []Methods.Method, handler Handler) {
	var controller = controller{name: name, handler: handler, pattern: pattern, methods: methods}
	controller.Init()
}

func (c *controller) Init() {
	http.HandleFunc(c.pattern, func(w http.ResponseWriter, r *http.Request) {
		if !canAccess(c.methods, r.Method) {
			fmt.Fprintf(w, "Error 404 Not Found")
			return
		}
		c.handler(w, r)
	})
}

func canAccess(methods []Methods.Method, method string) bool {
	accept := false

	for _, met := range methods {
		if method == string(met) {
			accept = true
		}
	}

	return accept
}
