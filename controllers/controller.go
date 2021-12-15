package controllers

import (
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

func UpdatePlays(id string) error {
	_, err := db.Exec("CALL increment_plays($1)", id)
	return err
}