package routes

import (
	"example/spudify/controllers"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Frontend routes
	router.Use(static.Serve("/", static.LocalFile("./views/js", true)))

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health_check", healthCheck)
		api.GET("/artists/:id", getArtistByID)
		api.GET("/albums/:id", getAlbumByID)
		api.GET("/songs/:id", getSongByID)
	}
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func getArtistByID(c *gin.Context) {
	id := c.Param("id")
	artist, err := controllers.GetArtistByID(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, artist)
	}
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, id)
}

func getSongByID(c *gin.Context) {
	id := c.Param("id")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, id)
}