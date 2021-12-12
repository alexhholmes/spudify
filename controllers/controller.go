package controllers

import (
	"context"
	. "example/spudify/models"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/georgysavva/scany/pgxscan"
)

// DB connection reference
var db *pgxpool.Pool
var ctx context.Context

func InitDB(dbConnection *pgxpool.Pool) {
	if dbConnection == nil {
		panic("InitDB() nil connection")
	}
	db = dbConnection
	ctx = context.Background()
}

func GetAllArtists() ([]*Artist, error) {
	var artists []*Artist
	query := "SELECT id, name, bio FROM artists"
	err := pgxscan.Select(ctx, db, &artists, query)
	return artists, err
}

func GetArtistByID(id string) ([]*Artist, error) {
	var artists []*Artist
	query := fmt.Sprintf("SELECT id, name, bio FROM artists WHERE id = %s", id)
	err := pgxscan.Select(ctx, db, &artists, query)
	return artists, err
}

func GetArtistAlbums(artistID string) ([]*Album, error) {
	var albums []*Album
	query := fmt.Sprintf("SELECT id, title, genre, artist_id FROM album WHERE artist_id = %s", artistID)
	err := pgxscan.Select(ctx, db, &albums, query)
	return albums, err
}

func GetArtistPopularSongs(id string) ([]*Song, error) {
	var songs []*Song
	query := fmt.Sprintf("SELECT id, name, genre, plays, duration, artist_id, album_id FROM songs WHERE artist_id = %s ORDER BY plays DESC LIMIT 5", id)
	err := pgxscan.Select(ctx, db, &songs, query)
	return songs, err
}

func GetAlbumByID(id string) ([]*Album, error) {
	var albums []*Album
	query := fmt.Sprintf("SELECT id, title, genre, artist_id FROM album WHERE id = %s", id)
	err := pgxscan.Select(ctx, db, &albums, query)
	return albums, err
}

func GetAlbumSongs(albumID string) ([]*Song, error) {
	var songs []*Song
	query := fmt.Sprintf("SELECT id, name, genre, plays, duration, artist_id, album_id FROM songs WHERE album_id = %s", albumID)
	err := pgxscan.Select(ctx, db, &songs, query)
	return songs, err
}

func GetSongByID(id string) ([]*Song, error) {
	var songs []*Song
	query := fmt.Sprintf("SELECT id, name, genre, plays, duration, artist_id, album_id FROM songs WHERE id = %s", id)
	err := pgxscan.Select(ctx, db, &songs, query)
	return songs, err
}

