package main

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"bytes"

	"github.com/gorilla/mux"
)

type playerHandler struct{}

func (p *playerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["num"])
	db := ConnectDB()

	var player Player

	player, err := getPlayerAllStats(db, uint(id))

	if err != nil {
		panic("failed to get player")
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&player); err != nil {
		panic("encode error")
	}
	fmt.Fprint(w, buf.String())
}
