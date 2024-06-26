package main

import (
	"example/web-service-gin/initializers"
	"example/web-service-gin/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func init() {
	if os.Getenv("USER") != "appuser" {
		gin.SetMode(gin.DebugMode)
	}
	gin.SetMode(gin.ReleaseMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}

// route handler
func main() {
	var err error
	DB, err = initializers.InitDB("webServiceGin.sqlite")
	if err != nil {
		log.Println("Couldn't connect to DB.")
		panic("Couldn't connect to DB.")
	}
	DB.AutoMigrate(&models.Album{})

	router := SetupRouter()

	router = getAlbums(router)
	router = postAlbum(router)
	router = postAlbums(router)
	router = getAlbumByID(router)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(r *gin.Engine) *gin.Engine {
	r.GET("/albums", func(c *gin.Context) {
		var albums []models.Album
		result := DB.Find(&albums)
		if result.Error != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "couldn't get albums"})
			return
		}
		c.IndentedJSON(http.StatusOK, albums)
	})
	return r
}

// postAlbum adds an album from JSON received in the request body.
func postAlbum(r *gin.Engine) *gin.Engine {
	r.POST("/albums", func(c *gin.Context) {
		var newAlbum *models.Album
		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if error := c.BindJSON(&newAlbum); error != nil {
			return
		}
		// Add the new album to the data slice
		DB.Create(&newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	})
	return r
}

func postAlbums(r *gin.Engine) *gin.Engine {
	r.POST("/addAlbums", func(c *gin.Context) {
		var payload []models.Album
		if error := c.BindJSON(&payload); error != nil {
			return
		}
		fmt.Println(&payload)
		var results []models.Album
		for a := range len(payload) {
			item := models.Album{Title: payload[a].Title, Artist: payload[a].Artist, Price: payload[a].Price}
			result := DB.Create(&item)
			if result.Error != nil {
				log.Println("unable to create Album")
			}
			results = append(results, item)
		}
		c.IndentedJSON(http.StatusCreated, results)
	})
	return r
}

func getAlbumByID(r *gin.Engine) *gin.Engine {
	r.GET("/albums/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
		var result models.Album
		// Look for the id in albums
		err := DB.Where("id = ?", id).First(&result)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			log.Printf("unable to find Album %v", id)
			return
		}
		c.IndentedJSON(http.StatusOK, result)
	})
	return r
}
