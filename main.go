package main

import (
	"fmt"
	"github.com/rishabhdeepsingh/go_simple_server/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
