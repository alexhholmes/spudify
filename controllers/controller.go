package controllers

import (
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

type Artist struct {
	ID        string    `json:"id"'`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Bio       string    `json:"bio"`
}

func GetArtistByID(id string) (*Artist, error) {
	artist := new(Artist)
	err := db.Model(artist).
		Where("id = ?", id).
		Select()
	return artist, err
}

/*
type Song struct {
	ID        string    `json:"id"'`
	Name      string    `json:"name"`
	Genre     string    `json:"genre"`
	Duration  int       `json:"duration"`
	ArtistID  string    `json:"artist_id"`
	Artist    string    `json:"artist"`
	AlbumID   string    `json:"album_id"`
	Album     string    `json:"album"`
}

func getSongByID(db *pg.DB) error {
	opts := orm.Query{

	}
}
 */