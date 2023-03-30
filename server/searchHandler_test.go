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

func doSearchTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/search?q=mock", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	setDummyPlayer(db, "dummy5", 1500)
	setDummyPlayer(db, "dummy1", 1900)
	setDummyPlayer(db, "dummy3", 1700)
	setDummyPlayer(db, "dummy2", 1800)
	setDummyPlayer(db, "dummy4", 1600)
	setDummyPlayer(db, "mock10", 1500)
	setDummyPlayer(db, "mock6", 1900)
	setDummyPlayer(db, "mock8", 1700)
	setDummyPlayer(db, "mock7", 1800)
	setDummyPlayer(db, "mock9", 1600)

	search(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var r SearchResponse
	err := json.Unmarshal(raw, &r)

	if err != nil {
		panic("failed to unmarshal response")
	}

	assert.Equal(http.StatusOK, resp.StatusCode)

	assert.Equal(r.Players[0].Name, "mock6")
	assert.Equal(r.Players[1].Name, "mock7")
	assert.Equal(r.Players[2].Name, "mock8")
	assert.Equal(r.Players[3].Name, "mock9")
	assert.Equal(r.Players[4].Name, "mock10")
}

func doSearchExceptionTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/search", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	search(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(http.StatusBadRequest, resp.StatusCode)

	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)
	assert.Equal("{\"error\": true, \"message\": \"incorrect search query\"}", body)
}

func doSearchPagingTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/ranking?p=1&q=dummy", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	for i := 0; i < 101; i++ {
		setDummyPlayer(db, "dummy", 1500)
	}

	search(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var r SearchResponse
	err := json.Unmarshal(raw, &r)

	if err != nil {
		panic("failed to unmarshal response")
	}

	assert.Equal(http.StatusOK, resp.StatusCode)
	assert.Equal(r.Players[0].Name, "dummy")
	assert.Equal(len(r.Players), 1)
}

func doSearchPagingExceptionTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/ranking?p=1", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	for i := 0; i < 100; i++ {
		setDummyPlayer(db, "dummy", 1500)
	}

	ranking(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(http.StatusBadRequest, resp.StatusCode)

	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)
	assert.Equal("{\"error\": true, \"message\": \"incorrect specified page number\"}", body)
}

func TestSearch(t *testing.T) {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings_test.json")
	if err != nil {
		panic("failed to load settings_test.json")
	}

	err = json.Unmarshal(raw, &s)

	if err != nil {
		panic("failed to unmarshal settings")
	}

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	if err != nil {
		panic("failed to connect test db, please exec `docker-compose up -d`")
	}

	doSearchTest(t, db)
	doSearchExceptionTest(t, db)
	doSearchPagingTest(t, db)
	doSearchPagingExceptionTest(t, db)
}
