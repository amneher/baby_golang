package main

import (
	"encoding/json"
	"example/web-service-gin/initializers"
	"example/web-service-gin/models"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func setupGetRequest(path string) (*httptest.ResponseRecorder, *http.Request) {
	var err error
	DB, err = initializers.InitDB("webServiceGin.sqlite")
	if err != nil {
		panic("Couldn't connect to DB.")
	}
	DB.AutoMigrate(&models.Album{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%v", path), nil)
	return w, req
}

func setupPostRequest(path string, payload io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	var err error
	DB, err = initializers.InitDB("webServiceGin.sqlite")
	if err != nil {
		panic("Couldn't connect to DB.")
	}
	DB.AutoMigrate(&models.Album{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/%v", path), payload)
	return w, req
}

func TestPing(t *testing.T) {
	router := SetupRouter()
	// call route
	w, req := setupGetRequest("ping")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGetAlbums(t *testing.T) {
	router := SetupRouter()
	// call route
	router = getAlbums(router)
	w, req := setupGetRequest("albums")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPostAlbum(t *testing.T) {
	router := SetupRouter()
	// call route
	router = postAlbum(router)
	testAlbum := models.Album{Title: "Ah Um", Artist: "Charles Mingus", Price: 49.99}
	albumJson, _ := json.Marshal(testAlbum)
	body_content := strings.NewReader(string(albumJson))
	w, req := setupPostRequest("albums", body_content)
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestPostAlbums(t *testing.T) {
	router := SetupRouter()
	// call route
	router = postAlbums(router)
	var test_albums []models.Album
	test_albums = append(test_albums, models.Album{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	test_albums = append(test_albums, models.Album{Title: "Ah Um", Artist: "Charles Mingus", Price: 49.99})
	album_json, _ := json.Marshal(test_albums)
	body_content := strings.NewReader(string(album_json))
	w, req := setupPostRequest("addAlbums", body_content)
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}
func TestGetAlbumByID(t *testing.T) {
	router := SetupRouter()
	var err error
	DB, err = initializers.InitDB("webServiceGin.sqlite")
	if err != nil {
		panic("Couldn't connect to DB.")
	}
	DB.AutoMigrate(&models.Album{})
	testValues := map[string]string{"id": uuid.New().String(), "title": "Ah Um", "artist": "Charles Mingus", "price": "49.99"}
	testAlbum := models.Album{Title: testValues["title"], Artist: testValues["artist"], Price: 49.99}
	testAlbumRes := DB.Create(&testAlbum)
	if testAlbumRes.Error != nil {
		fmt.Printf(
			"ID: %v, CreatedAt: %v, UpdatedAt: %v, DeletedAt: %v, Title: %v, Artist: %v, Price: %v",
			reflect.TypeOf(testAlbum.ID),
			reflect.TypeOf(testAlbum.CreatedAt),
			reflect.TypeOf(testAlbum.UpdatedAt),
			reflect.TypeOf(testAlbum.DeletedAt),
			reflect.TypeOf(testAlbum.Title),
			reflect.TypeOf(testAlbum.Artist),
			reflect.TypeOf(testAlbum.Price),
		)
		panic(fmt.Sprintf("Couldn't create test album %v", testAlbumRes.Error))
	}
	// call route
	router = getAlbumByID(router)

	var result models.Album
	error := DB.Find(&result)
	if error.Error != nil {
		panic("couldn't find an Album")
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%v", result.ID), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
