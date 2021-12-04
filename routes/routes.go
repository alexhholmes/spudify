package routes

import (
	"example/spudify/controllers"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

		api.POST("/login", login)
		api.GET("/logout", logout)

		api.GET("/artists/:id", getArtistByID)
		api.GET("/artists/:id/albums", getArtistAlbums)
		api.GET("/artists/:id/popular", getArtistPopularSongs)
		api.GET("/albums/:id", getAlbumByID)
		api.GET("/albums/:id/songs", getAlbumSongs)
		api.GET("/songs/:id", getSongByID)
	}
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verify username and password
	expectedPassword, err := controllers.GetUserPassword(username)
	if err != nil || password != expectedPassword {
		c.String(http.StatusUnauthorized, "Incorrect username or password")
		return
	}

	// Login user, start new session
	sessionID, _ := gonanoid.New()
	controllers.UpdateUserSession(username, sessionID)
	c.SetCookie("session", sessionID, 0, "/", "localhost", false, true)

	c.String(http.StatusOK, "Login successful")
}

func logout(c *gin.Context) {
	// Get current cookie and clear it
	cookie, _ := c.Cookie("session")
	c.SetCookie("session", "", -1, "/", "localhost", false, true)

	// Clear cookie from datastore
	err := controllers.ClearUserSession(cookie)
	if err != nil {
		c.String(http.StatusOK, "Logout successful")
		return
	}

	c.String(http.StatusUnauthorized, "Invalid session")
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

func getArtistAlbums(c *gin.Context) {
	id := c.Param("id")
	albums, err := controllers.GetArtistAlbums(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, albums)
	}
}

func getArtistPopularSongs(c *gin.Context) {
	id := c.Param("id")
	songs, err := controllers.GetArtistPopularSongs(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := controllers.GetAlbumByID(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, album)
	}
}

func getAlbumSongs(c *gin.Context) {
	id := c.Param("id")
	songs, err := controllers.GetAlbumSongs(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func getSongByID(c *gin.Context) {
	id := c.Param("id")
	song, err := controllers.GetSongByID(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, song)
	}
}