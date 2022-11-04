package main

import (
	"net/http"
	// "net/http/cgi"
	"fmt"
	"os"
	muxtrace "github.com/gorilla/mux"
)

type root struct{}

func (p *root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "root")
}

func main() {
	router := muxtrace.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")

	router.Handle(rootPath + "/ping", &ping{})
	router.Handle(rootPath + "/", &root{})

	server := &http.Server{
		Addr:         ":9999",
		Handler:      router,
	}

	server.ListenAndServe()  
	// cgi.Serve(mux)
}