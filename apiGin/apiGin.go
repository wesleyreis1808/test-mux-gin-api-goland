package apigin

import (
	"github.com/gin-gonic/gin"
)

var albums []album

func ApiGIN() {
	albums = getAlbumData()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
