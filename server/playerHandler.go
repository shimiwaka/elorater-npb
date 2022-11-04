package main

import (
	"net/http"
	"fmt"
    "strconv"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

type playerHandler struct{}

func (p *playerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Show Player %s\n", vars["num"])

	id, _ := strconv.Atoi(vars["num"])
	db := ConnectDB()

	var player Player

	db.First(&player, id)

	fmt.Fprintf(w, "%s", player)

}
