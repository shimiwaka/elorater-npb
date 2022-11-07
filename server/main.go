package main

import (
	"net/http"
	// "net/http/cgi"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	rootPath := os.Getenv("SCRIPT_NAME")

	router.Handle(rootPath + "/player/{num}", setAccessControlAllowHeader(&playerHandler{}))
	router.Handle(rootPath + "/pick-up", setAccessControlAllowHeader(&pickUpHandler{}))
	router.Handle(rootPath + "/vote", setAccessControlAllowHeader(&voteHandler{}))
	router.Handle(rootPath + "/ranking", setAccessControlAllowHeader(&rankingHandler{}))
	router.Handle(rootPath + "/ping", &pingHandler{})

	server := &http.Server{
		Addr:         ":9999",
		Handler:      router,
	}

	server.ListenAndServe()  
	// cgi.Serve(router)
}