package controllers

import (
	. "example/spudify/models"
)

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

	user := new(User)
	err := db.Model(user).
		Column("id").
		Where("session = ?", sessionID).
		Select()

	model := ArtistFollowers {
		UserID:   user.ID,
		ArtistID: artistID,
	}
	_, err = db.Model(model).Insert()
	return err
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

func AddSongToPlaylist(playlistID, songID string) error {
	type SongsPlaylists struct {
		SongID      string    `json:"song_id"`
		PlaylistID  string    `json:"playlist_id"`
	}
	songPlaylist := SongsPlaylists{
		SongID:     songID,
		PlaylistID: playlistID,
	}

	_, err := db.Model(songPlaylist).Insert()
	return err
}

func CreatePlaylist(session string) error {

}