package main

import (
	"fmt"
	"net/http"

	"go-frond-end/router"
)

func main() {
	r := router.Router()

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
