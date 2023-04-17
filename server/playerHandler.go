package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type playerHandler struct{}

func showPlayerData(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) {
	var player Player

	player, err := getPlayerAllStats(db, uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"incorrect specified ID\"}")
		return
	}
	player.Number = getPlayerRankedNumber(db, uint(id))

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{\"error\": true, \"message\": \"failed to encode json\"}")
		return
	}

	fmt.Fprint(w, buf.String())
}

func (p *playerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["num"])
	db := connectDB()

	showPlayerData(db, id, w, r)
	db.Close()
}
