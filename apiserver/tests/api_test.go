package mytest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"short_url/apiserver/api"
	"short_url/apiserver/pkg/database"
	"short_url/apiserver/pkg/net"
	"short_url/apiserver/pkg/shortener"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	r := net.RouteInit()
	w := performRequest(r, "GET", "/ping")

	assert.Equal(t, http.StatusOK, w.Code)

	response := api.Response{}
	// err := json.Unmarshal([]byte(w.Body.String()), &response)
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, "pong", response.Message)
}

func TestGetLinks(t *testing.T) {
	db := database.TestDBInit()
	defer database.TestDBFree(db)
	r := net.RouteInit()

	// mock data in db
	link1 := database.Link{}
	link1.Url = "https://www.google.com"
	link1.ID = uuid.New()
	link1.Code = shortener.GenerateCode(link1.ID)
	db.Create(&link1)

	link2 := database.Link{}
	link2.Url = "https://www.yahoo.com"
	link2.ID = uuid.New()
	link2.Code = shortener.GenerateCode(link2.ID)
	db.Create(&link2)

	w := performRequest(r, "GET", "/api/links")

	// var response map[string]string
	// err := json.Unmarshal([]byte(w.Body.String()), &response)

	response := api.LinksResponse{}
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.Nil(t, err)

	var payload []database.Link = response.Payload
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(payload))
	assert.Equal(t, "https://www.yahoo.com", payload[1].Url)
}
