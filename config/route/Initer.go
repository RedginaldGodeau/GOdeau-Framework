package route

import (
	"GOdeau/godeau/Controllers"
	"GOdeau/godeau/Controllers/Methods"
	"fmt"
	"net/http"
)

func Init() {

	Controllers.New("hello_world", "/", Methods.List(Methods.GET), func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Word")
	})

}
