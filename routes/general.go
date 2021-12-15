package routes

import (
	"net/http"

	"example/spudify/controllers"
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func getAllArtists(c *gin.Context) {
	artists, err := controllers.GetAllArtists()
	// TODO
	c.SetCookie("session", "bruh", 0, "/", "localhost:8000", false, true)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, artists)
	}
}

func getArtistByID(c *gin.Context) {
	id := c.Param("id")
	artist, err := controllers.GetArtistByID(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, artist)
	}
}

func getArtistAlbums(c *gin.Context) {
	id := c.Param("id")
	albums, err := controllers.GetArtistAlbums(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, albums)
	}
}

func getArtistPopularSongs(c *gin.Context) {
	id := c.Param("id")
	songs, err := controllers.GetArtistPopularSongs(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := controllers.GetAlbumByID(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, album)
	}
}

func getAlbumSongs(c *gin.Context) {
	id := c.Param("id")
	songs, err := controllers.GetAlbumSongs(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func getSongByID(c *gin.Context) {
	id := c.Param("id")
	song, err := controllers.GetSongByID(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, song)
	}
}

func updatePlays(c *gin.Context) {
	id := c.Param("id")
	err := controllers.UpdatePlays(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	} else {
		c.JSON(http.StatusOK, "")
	}
}