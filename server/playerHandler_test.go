package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func doShowPlayerDataTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/0", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	id := setDummyPlayer(db, "dummy", 1500)

	showPlayerData(db, int(id), w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(http.StatusOK, resp.StatusCode)

	raw, _ := io.ReadAll(resp.Body)
	var player Player
	json.Unmarshal(raw, &player)

	assert.Equal("dummy", player.Name)
	assert.Equal(1500, player.Rate)
	assert.Equal("2000", player.Pitching[0].Year)
	assert.Equal(2.00, player.Pitching[0].Era)
	assert.Equal("2000", player.Batting[0].Year)
	assert.Equal(0.300, player.Batting[0].Avg)
}

func doShowPlayerDataExceptionTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/0", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	showPlayerData(db, 0, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(http.StatusBadRequest, resp.StatusCode)
	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)
	assert.Equal("{\"error\": true, \"message\": \"incorrect specified ID\"}", body)
}

func TestShowPlayerData(t *testing.T) {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings_test.json")
	if err != nil {
		panic("failed to load settings_test.json")
	}

	json.Unmarshal(raw, &s)

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	doShowPlayerDataTest(t, db)
	doShowPlayerDataExceptionTest(t, db)
}
