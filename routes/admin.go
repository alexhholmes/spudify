package routes

import (
	"net/http"
	"os"
	"strconv"

	"example/spudify/controllers"
	"example/spudify/models"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3

func Init_s3() {
	key := os.Getenv("OIM5L4D7FVTTHNNOVB4R")
	secret := os.Getenv("634ishVkot21vbhJ6AAd75TwXYtbKgtmM2Pofxyd/K0")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"),
		Region:      aws.String("us-east-1"),
	}

	newSession := session.New(s3Config)
	s3Client = s3.New(newSession)
}

func verifyAdmin(token string) bool {
	return token == "password"
}

func createArtist(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if !verifyAdmin(c.PostForm("token")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
		return
	}

	id, _ := gonanoid.New()
	name := c.PostForm("name")
	bio := c.PostForm("bio")
	artist := models.Artist{
		ID:   id,
		Name: name,
		Bio:  bio,
		TotalPlays: 0,
	}

	artistID, err := controllers.CreateArtist(artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"artist_id":artistID})
}

func createAlbum(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if !verifyAdmin(c.PostForm("token")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
		return
	}

	id, _ := gonanoid.New()
	title := c.PostForm("title")
	genre := c.PostForm("genre")
	artistID := c.PostForm("artist_id")
	album := models.Album{
		ID:       id,
		Title:    title,
		Genre:    genre,
		ArtistID: artistID,
	}


	albumID, artistID, err := controllers.CreateAlbum(album)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"album_id":albumID, "artist_id":artistID})
}

func createSong(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	if !verifyAdmin(c.PostForm("token")) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad token"})
		return
	}

	id, _ := gonanoid.New()
	name := c.PostForm("name")
	genre := c.PostForm("genre")
	duration, err := strconv.Atoi(c.PostForm("duration"))
	artistID := c.PostForm("artist_id")
	albumID := c.PostForm("album_id")
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
		c.JSON(http.StatusBadRequest, "")
		return
	}

	songID, err := controllers.CreateSong(song)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"song_id":songID})
}