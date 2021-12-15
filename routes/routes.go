package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Dev env stuff
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Frontend static files
	router.Use(static.Serve("/", static.LocalFile("spudify-vue/dist/", true)))
	router.Use(static.Serve("/admin", static.LocalFile("admin/", true)))

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health_check", healthCheck)

		api.POST("/login", login)

		api.GET("/me/playlists", getCurrentUserPlaylists)
		api.POST("/me/playlists", createPlaylist)
		api.DELETE("/me/playlists/:id", deletePlaylist)
		api.GET("/playlists/:id", getPlaylistSongs)
		api.POST("/playlists/:id/song/:song_id", addSongToPlaylist)
		api.DELETE("/playlists/:id/song/:song_id", deleteSongFromPlaylist)

		api.GET("/me/following", getCurrentUserFollowing)
		api.POST("/me/following/:id", followArtist)
		api.DELETE("/me/following/:id", unfollowArtist)

		api.GET("/artists", getAllArtists)
		api.GET("/artists/:id", getArtistByID)
		api.GET("/artists/:id/albums", getArtistAlbums)
		api.GET("/artists/:id/popular", getArtistPopularSongs)
		api.GET("/albums/:id", getAlbumByID)
		api.GET("/albums/:id/songs", getAlbumSongs)
		api.GET("/songs/:id", getSongByID)
		api.PUT("/songs/:id", updatePlays)

		// Admin access only
		api.POST("/artist", createArtist)
		api.POST("/album", createAlbum)
		api.POST("/song", createSong)
	}
}
