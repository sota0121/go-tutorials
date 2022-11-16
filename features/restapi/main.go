package restapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var (
	domain = "localhost"
	port   = "8080"
)

// Main is the entrypoint for the restapi package.
func Main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	serverURL := fmt.Sprintf("%s:%s", domain, port)
	router.Run(serverURL)
}

// getAlbums returns the list of albums as JSON.
// [Note]
// ref: https://go.dev/doc/tutorial/web-service-gin
// > gin.Context is the most important part of Gin.
// > It carries request details, validates and serializes JSON, and more.
// > (Despite the similar name, this is different from Go’s built-in context package.)
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
	// c.IndentedJSON(http.StatusCreated, newAlbum) // this is for debug 'cause pretty print is heavy
}
