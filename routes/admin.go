package routes

import (
	"example/spudify/controllers"
	"example/spudify/models"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
	"strconv"
)

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func createArtist(c *gin.Context) {
	id, _ := gonanoid.New()
	name := c.Param("name")
	bio := c.Param("bio")
	artist := models.Artist{
		ID:   id,
		Name: name,
		Bio:  bio,
	}

	err := controllers.CreateAlbum(artist)
	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}

func createAlbum(c *gin.Context) {
	id, _ := gonanoid.New()
	title := c.Param("title")
	genre := c.Param("genre")
	artistID := c.Param("artist_id")
	album := models.Album{
		ID:       id,
		Title:    title,
		Genre:    genre,
		ArtistID: artistID,
	}


	err := controllers.CreateAlbum(album)
	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}

func createSong(c *gin.Context) {
	id, _ := gonanoid.New()
	name := c.Param("name")
	genre := c.Param("genre")
	duration, err := strconv.Atoi(c.Param("duration"))
	artistID := c.Param("artist_id")
	albumID := c.Param("album_id")
	song := models.Song{
		ID:       id,
		Name:     name,
		Genre:    genre,
		Plays:    0,
		Duration: duration,
		ArtistID: artistID,
		AlbumID:  albumID,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	err := controllers.CreateSong(song)
	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}