package main

import (
	"net/http"
	// "net/http/cgi"
	"fmt"
	"os"
	"github.com/gorilla/mux"
)

type rootHandler struct{}

func (p *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "root")
}

func main() {
	router := mux.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")

	router.Handle(rootPath + "/player/{num}", &playerHandler{})
	router.Handle(rootPath + "/pick-up", &pickUpHandler{})
	router.Handle(rootPath + "/vote", &voteHandler{})

	router.Handle(rootPath + "/ranking", &rankingHandler{})
	router.Handle(rootPath + "/ping", &pingHandler{})
	router.Handle(rootPath + "/", &rootHandler{})

	server := &http.Server{
		Addr:         ":9999",
		Handler:      router,
	}

	server.ListenAndServe()  
	// cgi.Serve(mux)
}