package controllers

import (
	. "example/spudify/models"
	"github.com/georgysavva/scany/pgxscan"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GetUserPassword(username string) (string, error) {
	var password string
	query := `
SELECT password
FROM users
WHERE id = $1`
	err := db.QueryRow(ctx, query, username).Scan(&password)
	return password, err
}

func UpdateUserSession(username string, session string) error {
	query := `
UPDATE users
SET session = $1
WHERE username = $2`
	_, err := db.Exec(ctx, query, session, username)
	return err
}

func GetCurrentUserPlaylists(sessionID string) ([]*Playlist, error) {
	var playlists []*Playlist
	query := `
SELECT id, name, description, owner_id
FROM playlists
WHERE owner_id IN (
	SELECT id
	FROM users
	WHERE session = $1
	)`
	err := pgxscan.Select(ctx, db, &playlists, query, sessionID)
	return playlists, err
}

func GetCurrentUserFollowing(sessionID string) ([]*Artist, error) {
	var artists []*Artist
	query := `
SELECT id, name, bio
FROM artists
WHERE id IN (
	SELECT artist_id
	FROM artist_followers
	WHERE user_id IN (
		SELECT id
		FROM users
		WHERE session = $1
		)
	)`
	err := pgxscan.Select(ctx, db, &artists, query, sessionID)
	return artists, err
}

func FollowArtist(sessionID, artistID string) error {
	userID, err := getUserID(sessionID)
	if err != nil {
		return err
	}

	query := `
INSERT INTO artist_followers(user_id, artist_id)
VALUES($1, $2)`
	_, err = db.Exec(ctx, query, userID, artistID)

	return err
}

func GetPlaylistItems(id string) ([]*Song, error) {
	var songs []*Song
	query := `
SELECT id, name, genre, plays, duration, artist_id, album_id
FROM songs
WHERE id IN (
	SELECT song_id
	FROM songs_playlists
	WHERE playlist_id = $1
	)`
	err := pgxscan.Select(ctx, db, &songs, query, id)
	return songs, err
}

func AddSongToPlaylist(playlistID, songID string) error {
	query := `
INSERT INTO songs_playlists(song_id, playlist_id)
VALUES($1, $2)`
	_, err := db.Exec(ctx, query, songID, playlistID)
	return err
}

func CreatePlaylist(session, name, description string) error {
	ownerID, err := getUserID(session)
	if err != nil {
		return err
	}

	id, _ := gonanoid.New()
	query := `
INSERT INTO playlists(id, name, description, owner_id)
VALUES($1, $2, $3, $4)`
	_, err = db.Exec(ctx, query, id, name, description, ownerID)
	return err
}

func getUserID(sessionID string) (string, error) {
	var userID string
	query := `
SELECT id
FROM users
WHERE session = $1`
	err := db.QueryRow(ctx, query, sessionID).Scan(&userID)
	return userID, err
}
