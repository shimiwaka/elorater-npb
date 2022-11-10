package main

import (
	"encoding/json"
	"testing"
	"os"
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type VoteTestCase struct {
	Name string
	Vote int
	Token string
	ExpectStatus int
	ExpectResponseBody string
	ExpectSortedName []string
	ExpectSortedRate []int
}

func mockRanking(db *gorm.DB) RankingResponse {
	resp := RankingResponse{}
	players := []Player{}

	db.Limit(100).Order("rate desc").Find(&players)

	for _, player := range players {
		resp.Players = append(resp.Players, RankedPlayer{Name: player.Name, Rate: player.Rate, Id: player.ID})
	}

	return resp
}

func doVoteTest(t *testing.T, db *gorm.DB, tc VoteTestCase) {
	req := httptest.NewRequest(http.MethodGet,
		 fmt.Sprintf("http://example.com/vote?c=%d&token=%s",tc.Vote, tc.Token), nil)
	w := httptest.NewRecorder()

	assert := assert.New(t)
	initializeDB(db)
	setDummyToken(db)

	vote(db, w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(tc.ExpectStatus, resp.StatusCode)

	if tc.ExpectResponseBody != "" {
		raw, _ := io.ReadAll(resp.Body)
		body := string(raw)
		assert.Equal(tc.ExpectResponseBody, body)
	}
	ranking := mockRanking(db)

	for i, player := range ranking.Players {
		assert.Equal(tc.ExpectSortedName[i], player.Name)
		assert.Equal(tc.ExpectSortedRate[i], player.Rate)
	}
}

func TestVote(t *testing.T) {
	s := Settings{}
	raw, err := os.ReadFile("./config/settings_test.json")
	if err != nil {
		panic("failed to load settings_test.json")
	}

	json.Unmarshal(raw, &s)

	connectQuery := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
					s.DB_username, s.DB_pass, s.DB_host, s.DB_port, s.DB_name)

	db, err := gorm.Open("mysql", connectQuery)

	if err != nil {
		panic("failed to connect test db, please exec `docker-compose up -d`")
	}

	tests := []VoteTestCase {
			{	Name: "vote player 1",
				Vote: 1,
				Token: "DUMMY",
				ExpectStatus: http.StatusOK,
				ExpectResponseBody: "{\"error\": false}",
				ExpectSortedName: []string { "dummy1", "dummy2" },
				ExpectSortedRate: []int { 1516, 1485 },
			},
			{	Name: "vote player 2",
				Vote: 2,
				Token: "DUMMY",
				ExpectStatus: http.StatusOK,
				ExpectResponseBody: "{\"error\": false}",
				ExpectSortedName: []string { "dummy2", "dummy1" },
				ExpectSortedRate: []int { 1516, 1485 },
			},
		}

	for _, tc := range tests {
		doVoteTest(t, db, tc)
	}
}