package routes

import (
	"example/spudify/controllers"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"net/http"
)

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

func createPlaylist(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	session, _ := c.Cookie("session")

	err := controllers.CreatePlaylist(session, c.PostForm("name"), c.PostForm("description"))
	if err != nil {
		c.JSON(http.StatusOK, "")
		return
	}
	c.JSON(http.StatusBadRequest, "")
}

func deletePlaylist(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}

func getPlaylistSongs(c *gin.Context) {
	id := c.Param("id")
	songs, err := controllers.GetPlaylistItems(id)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(http.StatusOK, songs)
	}
}

func addSongToPlaylist(c *gin.Context) {
	id := c.Param("id")
	songID := c.Param("song_id")
	err := controllers.AddSongToPlaylist(id, songID)

	c.Header("Content-Type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	} else {
		c.JSON(http.StatusOK, "")
	}
}

func deleteSongFromPlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, "")
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

func followArtist(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	artistID := c.Param("id")
	session, _ := c.Cookie("session")
	err := controllers.FollowArtist(session, artistID)
	if err != nil {
		c.JSON(http.StatusOK, "")
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session"})
}

func unfollowArtist(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "")
}
