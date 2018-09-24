package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handlerFunc(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "<h1>Welcome to my awesome site!</h1>")
}

// home
func home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>Welcome to my awesome site!</h1>")

}

func conact(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")

}

func faq(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>FAQ --- Coming Soon</h1>")

}

func notFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	fmt.Fprint(writer, "<h1>404 NOT FOUND</h1>")

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", conact)
	router.HandleFunc("/faq", faq)

	//custom 404
	var handler http.Handler = http.HandlerFunc(notFound)
	router.NotFoundHandler = handler
	http.ListenAndServe(":3030", router)
}
