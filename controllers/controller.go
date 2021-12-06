package controllers

import (
	"database/sql"
	. "example/spudify/models"
	"github.com/go-pg/pg/v10"
)

// DB connection reference
var db *pg.DB
func InitDB(dbConnection *pg.DB) {
	if dbConnection == nil {
		panic("InitDB() nil connection")
	}
	db = dbConnection
}

func GetUserPassword(username string) (string, error) {
	var password string
	err := db.Model((*User)(nil)).
		Column("password").
		Where("username = ?", username).
		Select(&password)
	return password, err
}

func UpdateUserSession(username string, session string) error {
	_, err := db.Model((*User)(nil)).
		Set("session = ?", session).
		Where("username = ?", username).
		Update()
	return err
}

func ClearUserSession(session string) error {
	_, err := db.Model((*User)(nil)).
		Set("session = ?", sql.NullString{}).
		Where("session = ?", session).
		Update()
	return err
}

func GetAllArtists() ([]Artist, error) {
	var artists []Artist
	err := db.Model(&artists).
		Select()
	return artists, err
}

func GetArtistByID(id string) (*Artist, error) {
	artist := new(Artist)
	err := db.Model(artist).
		Where("id = ?", id).
		Select()
	return artist, err
}

func GetArtistAlbums(artistID string) ([]Album, error) {
	var albums []Album
	err := db.Model(&albums).
		Where("artist_id = ?", artistID).
		Select()
	return albums, err
}

func GetArtistPopularSongs(id string) ([]Song, error) {
	var songs []Song
	err := db.Model(&songs).
		Where("artist_id = ?", id).
		Order("plays DESC").
		Limit(5).
		Select()
	return songs, err
}

func GetAlbumByID(id string) (*Album, error) {
	album := new(Album)
	err := db.Model(album).
		Where("id = ?", id).
		Select()
	return album, err
}

func GetAlbumSongs(albumID string) ([]Song, error) {
	var songs []Song
	err := db.Model(&songs).
		Where("album_id = ?", albumID).
		Select()
	return songs, err
}

func GetSongByID(id string) (*Song, error) {
	song := new(Song)
	err := db.Model(song).
		Where("id = ?", id).
		Select()
	return song, err
}

func GetCurrentUserPlaylists(sessionID string) ([]Playlist, error) {
	userID := db.Model((*User)(nil)).
		Column("id").
		Where("session = ?", sessionID)

	var playlists []Playlist
	err := db.Model(&playlists).
		Where("owner_id IN (?)", userID).
		Select()

	return playlists, err
}

func GetCurrentUserFollowing(sessionID string) ([]Artist, error) {
	userID := db.Model((*User)(nil)).
		Column("id").
		Where("session = ?", sessionID)

	var artists []Artist
	err := db.Model(&artists).
		Where("owner_id IN (?)", userID).
		Select()

	return artists, err
}

func FollowArtist(sessionID, artistID string) error {
	type ArtistFollowers struct {
		UserID      string    `json:"user_id"`
		ArtistID    string    `json:"artist_id"`
	}

	userID := db.Model((*User)(nil)).
		Column("id").
		Where("session = ?", sessionID).
		Select()

	model = ArtistFollowers{
		UserID:   userID,
		ArtistID: artistID,
	}
	err := db.Model().Insert()

}

func GetPlaylistItems(id string) ([]Song, error) {
	// For songs_playlists table
	type SongsPlaylists struct {
		SongID      string    `json:"song_id"`
		PlaylistID  string    `json:"playlist_id"`
	}

	var songs []Song
	// Select song IDs that match the playlist ID
	songIDs := db.Model((*SongsPlaylists)(nil)).
		ColumnExpr("song_id").
		Where("playlist_id = ?", id)
	// Select songs that are in the subquery
	err := db.Model(&songs).
		Where("song_id IN (?)", songIDs).
		Select()

	return songs, err
}

func AddSongsToPlaylist(userID, playlistID string, songIDs []string) {

}

func DeleteSongsFromPlaylist(userID, playlistID string, songIDs []string) {

}