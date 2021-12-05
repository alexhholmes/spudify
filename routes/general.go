package routes

import (
	"log"
	"net/http"

	"example/spudify/controllers"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func healthCheck(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func getCurrentUserPlaylists(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	session, _ := c.Cookie("session")
	content, err := controllers.GetCurrentUserPlaylists(session)
	if err != nil {
		c.JSON(http.StatusOK, content)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session"})
}

func getCurrentUserFollowing(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	session, _ := c.Cookie("session")
	content, err := controllers.GetCurrentUserFollowing(session)
	if err != nil {
		c.JSON(http.StatusOK, content)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session"})
}

func getPlaylistByID(c *gin.Context) {
	id := c.Param("id")
	songs, err := controllers.GetPlaylistItems(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusOK, "")
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func login(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verify username and password
	expectedPassword, err := controllers.GetUserPassword(username)
	if err != nil || password != expectedPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect username or password"})
		log.Printf("Failed login attempt from username:%s", username)
		return
	}

	// Login user, start new session
	sessionID, _ := gonanoid.New()
	controllers.UpdateUserSession(username, sessionID)
	c.SetCookie("session", sessionID, 0, "/", "localhost", false, true)
	log.Printf("Successful login attempt from username:%s", username)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func logout(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Get current cookie and clear it
	cookie, _ := c.Cookie("session")
	c.SetCookie("session", "", -1, "/", "localhost", false, true)

	// Clear cookie from datastore
	err := controllers.ClearUserSession(cookie)
	if err != nil {
		log.Printf("Successful logout attempt from session:%s", cookie)
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
		return
	}

	log.Printf("Failed logout attempt from session:%s", cookie)
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session"})
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
