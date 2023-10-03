package Route

import (
	"GOdeau/Module/Controllers"
	"net/http"
)

type icontroller struct {
	Package string `yaml:"package"`
}

func invokeHandler(pattern string, name string, methods []string, redirect string) {

	controller := Controllers.Controller

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if !haveMethod(methods, r.Method) {
			http.Redirect(w, r, redirect, http.StatusSeeOther)
			return
		}

		for s := range controller {
			println("func : ", s)
		}

		//controller[name](w, r)
	})
}
