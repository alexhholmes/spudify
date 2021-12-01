package routes

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Frontend Routes
	router.Use(static.Serve("/", static.LocalFile("./views/js", true)))

	// API Routes
	api := router.Group("/api")
	{
		api.GET("/health_check", healthCheck)
		api.GET("/artist/:id", getArtistByID)
		api.GET("/album/:id", getAlbumByID)
		api.GET("/song/:id", getSongByID)
	}
}

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func getArtistByID(c *gin.Context) {
	id := c.Param("id")
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, id)
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