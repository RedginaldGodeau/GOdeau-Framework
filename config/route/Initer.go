package route

import (
	"GOdeau/godeau/Controllers"
	"GOdeau/godeau/Controllers/Methods"
	"GOdeau/src/controller"
	"fmt"
	"net/http"
)

func Init() {

	Controllers.New("hello_world", "/", Methods.List(Methods.GET), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Word")
	})

	Controllers.New("ping_get", "/ping", Methods.List(Methods.GET), controller.PingGet)
}
