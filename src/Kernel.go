package src

import (
	"Hooter/Module/Route"
	"fmt"
	"net/http"
)

func Kernel() {

	err := Route.InitRoute("config/routes/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
