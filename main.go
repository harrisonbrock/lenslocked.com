package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Welcome to my awesome site!</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3030", nil)
}
