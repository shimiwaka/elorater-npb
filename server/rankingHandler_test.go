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

func doRankingTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/ranking", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	setDummyPlayer(db, "dummy5", 1500)
	setDummyPlayer(db, "dummy1", 1900)
	setDummyPlayer(db, "dummy3", 1700)
	setDummyPlayer(db, "dummy2", 1800)
	setDummyPlayer(db, "dummy4", 1600)

	ranking(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var r RankingResponse
	json.Unmarshal(raw, &r)

	assert.Equal(http.StatusOK, resp.StatusCode)

	assert.Equal(r.Players[0].Name, "dummy1")
	assert.Equal(r.Players[1].Name, "dummy2")
	assert.Equal(r.Players[2].Name, "dummy3")
	assert.Equal(r.Players[3].Name, "dummy4")
	assert.Equal(r.Players[4].Name, "dummy5")
}

func doRankingExceptionTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/ranking", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	ranking(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(http.StatusBadRequest, resp.StatusCode)

	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)
	assert.Equal("{\"error\": true, \"message\": \"incorrect specified page number\"}", body)
}

func doRankingPagingTest(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/player/ranking?p=1", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	for i := 0; i < 101; i++ {
		setDummyPlayer(db, "dummy", 1500)
	}

	ranking(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	var r RankingResponse
	json.Unmarshal(raw, &r)

	assert.Equal(http.StatusOK, resp.StatusCode)
	assert.Equal(r.Players[0].Name, "dummy")
	assert.Equal(len(r.Players), 1)
}

func doRankingPagingExceptionTest(t *testing.T, db *gorm.DB) {
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

func TestRanking(t *testing.T) {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings_test.json")
	if err != nil {
		panic("failed to load settings_test.json")
	}

	json.Unmarshal(raw, &s)

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	doRankingTest(t, db)
	doRankingExceptionTest(t, db)
	doRankingPagingTest(t, db)
	doRankingPagingExceptionTest(t, db)
}
