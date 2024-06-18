package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupGetRequest(path string) *httptest.ResponseRecorder {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%v", path), nil)
	router.ServeHTTP(w, req)
	return w
}

func setupPostRequest(path string, payload io.Reader) *httptest.ResponseRecorder {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/%v", path), payload)
	router.ServeHTTP(w, req)
	return w
}

func TestPing(t *testing.T) {
	w := setupGetRequest("ping")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetAlbums(t *testing.T) {
	w := setupGetRequest("albums")
	assert.Equal(t, 200, w.Code)
}
func TestPostAlbums(t *testing.T) {
	test_album := Album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}
	album_json, _ := json.Marshal(test_album)
	body_content := strings.NewReader(string(album_json))
	w := setupPostRequest("albums", body_content)
	assert.Equal(t, 200, w.Code)
}
func TestGetAlbumByID(t *testing.T) {
	w := setupGetRequest("albums")
	assert.Equal(t, 200, w.Code)
}
