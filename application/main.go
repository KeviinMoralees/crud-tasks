package main

import (
	routes "crud-tasks/infrastructure/entrypoints/rest"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	panic(http.ListenAndServe("localhost:8080", &routes.HttpServer{}))
}
