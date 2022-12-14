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

type PickUpTestCase struct {
	Name               string
	DummyPlayerNum     int
	ExpectStatus       int
	ExpectResponseBody string
	RatedDummy         []DummyPlayer
}

func doPickUpTest(t *testing.T, db *gorm.DB, tc PickUpTestCase) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/pick-up", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	for i := 0; i < tc.DummyPlayerNum; i++ {
		setDummyPlayer(db, "dummy", 1500)
	}
	for _, dummy := range tc.RatedDummy {
		setDummyPlayer(db, dummy.Name, dummy.Rate)
	}

	pickUp(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(tc.ExpectStatus, resp.StatusCode)

	raw, _ := io.ReadAll(resp.Body)
	if tc.ExpectResponseBody != "" {
		body := string(raw)
		assert.Equal(tc.ExpectResponseBody, body)
	} else {
		r := PickUpResponse{}
		token := Token{}
		err := json.Unmarshal(raw, &r)

		if err != nil {
			panic("failed to unmarshal response")
		}

		db.First(&token, "token = ?", r.Token)
		assert.Equal(r.Player1.ID, token.Player1_id)
		assert.Equal(r.Player2.ID, token.Player2_id)
	}
}

func doPickUpTestUnusual(t *testing.T, db *gorm.DB) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/pick-up", nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)

	setDummyPlayer(db, "dummy", 1500)
	setDummyPlayer(db, "dummy", 1500)
	setDummyPlayer(db, "high_rate_dummy", 2000)
	setDummyPlayer(db, "high_rate_dummy", 2000)

	pickUp(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	r := PickUpResponse{}
	err := json.Unmarshal(raw, &r)

	if err != nil {
		panic("failed to unmarshal response")
	}

	assert.Equal(r.Player1.Name, r.Player2.Name)
}

func TestPickUp(t *testing.T) {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings_test.json")
	if err != nil {
		panic("failed to load settings_test.json")
	}

	err = json.Unmarshal(raw, &s)
	if err != nil {
		panic("failed to unmarshall settings")
	}

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	if err != nil {
		panic("failed to connect test db, please exec `docker-compose up -d`")
	}

	tests := []PickUpTestCase{
		{Name: "there are no players",
			DummyPlayerNum:     0,
			ExpectStatus:       http.StatusInternalServerError,
			ExpectResponseBody: "{\"error\": true, \"message\": \"no player data in database\"}",
		},
		{Name: "there is 1 player",
			DummyPlayerNum:     1,
			ExpectStatus:       http.StatusInternalServerError,
			ExpectResponseBody: "{\"error\": true, \"message\": \"no player data in database\"}",
		},
		{Name: "there are 2 player",
			DummyPlayerNum:     2,
			ExpectStatus:       http.StatusOK,
			ExpectResponseBody: "",
		},
		{Name: "there is high rate player",
			DummyPlayerNum:     1,
			ExpectStatus:       http.StatusOK,
			ExpectResponseBody: "",
			RatedDummy:         []DummyPlayer{{Name: "high_rate_dummy", Rate: 2000}},
		},
		{Name: "there is low rate player",
			DummyPlayerNum:     1,
			ExpectStatus:       http.StatusOK,
			ExpectResponseBody: "",
			RatedDummy:         []DummyPlayer{{Name: "low_rate_dummy", Rate: 1000}},
		},
	}

	for _, tc := range tests {
		doPickUpTest(t, db, tc)
	}

	for i := 0; i < 10; i++ {
		doPickUpTestUnusual(t, db)
	}
}
