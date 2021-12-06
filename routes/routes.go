package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Frontend static files
	router.Use(static.Serve("/", static.LocalFile("spudify-vue/dist/", true)))

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health_check", healthCheck)

		api.POST("/login", login)

		api.GET("/me/playlists", getCurrentUserPlaylists)
		api.POST("/me/playlists/:id", createPlaylist)
		api.DELETE("/me/playlists/:id", deletePlaylist)
		api.GET("/playlist/:id", getPlaylistSongs)
		api.POST("/playlist/:id/song/:song_id", addSongToPlaylist)
		api.DELETE("/playlist/:id/song/:song_id", deleteSongFromPlaylist)

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

		// Admin access only
		api.POST("/artist", createArtist)
		api.POST("/album", createAlbum)
		api.POST("/song", createSong)
	}
}
