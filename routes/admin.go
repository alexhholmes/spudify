package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func createArtist(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}
func createAlbum(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}
func createSong(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}
func deleteArtist(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}
func deleteAlbum(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}
func deleteSong(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}