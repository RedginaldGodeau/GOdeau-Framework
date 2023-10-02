package Route

import (
	"fmt"
	"net/http"
	"reflect"
)

type Controllers struct {
	Controller Controller
}

type Controller struct{}

func invokeHandler(pattern string, fu string, methods []string, redirect string) {
	var c Controllers

	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if !haveMethod(methods, req.Method) {
			http.Redirect(w, req, redirect, http.StatusSeeOther)
			return
		}

		args := make([]reflect.Value, 2)
		args[0] = reflect.ValueOf(w)
		args[1] = reflect.ValueOf(req)

		fmt.Println("test")

		reflect.ValueOf(&c).MethodByName(fu).Call(args)
	})
}
