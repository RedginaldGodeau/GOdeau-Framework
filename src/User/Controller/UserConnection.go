package Controller

import (
	"Hooter/Module/Route"
	"fmt"
	"net/http"
)

func (c *Route.Controllers.Controller) Ping(w http.ResponseWriter) {
	fmt.Fprintln(w, "PONG !")
}
