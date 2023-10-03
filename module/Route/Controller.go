package Route

import (
	"GOdeau/module/Controllers"
	"net/http"
)

func invokeHandler(pattern string, name string, methods []string, redirect string) {

	controller := Controllers.Controller

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if !haveMethod(methods, r.Method) {
			http.Redirect(w, r, redirect, http.StatusSeeOther)
			return
		}

		for s := range controller {
			controller[s](w, r)
		}
	})
}
