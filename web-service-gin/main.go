package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// route handler
func main() {
	DB, err = InitDB("webServiceGin.sqlite")
	if err != nil {
		panic("Couldn't connect to DB.")
	}
	DB.AutoMigrate(&Album{})

	router := setupRouter()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	var albums Album
	allAlbums := DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, allAlbums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if error := c.BindJSON(&newAlbum); error != nil {
		return
	}
	// Add the new album to the data slice
	DB.Create(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var result Album
	// Look for the id in albums
	err := DB.First(&result, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
