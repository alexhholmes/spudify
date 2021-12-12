package controllers

import (
	. "example/spudify/models"
)

func CreateArtist(artist Artist) error {
	query := `
INSERT INTO artists(id, name, bio)
VALUES($1, $2, $3)`
	_, err := db.Exec(ctx, query, artist.ID, artist.Name, artist.Bio)
	return err
}

func CreateAlbum(album Album) error {
	query := `
INSERT INTO albums(id, title, genre, artist_id)
VALUES($1, $2, $3, $4)`
	_, err := db.Exec(ctx, query, album.ID, album.Title, album.Genre, album.ArtistID)
	return err
}

func CreateSong(song Song) error {
	query := `
INSERT INTO songs(id, name, genre, plays, duration, artist_id, album_id)
VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(ctx, query, song.ID, song.Name, song.Genre, 0, song.Duration, song.ArtistID, song.AlbumID)
	return err
}
