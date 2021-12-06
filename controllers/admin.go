package controllers

import (
	. "example/spudify/models"
)

func CreateArtist(artist Artist) error {
	_, err := db.Model(artist).
		Insert()
	return err
}

func CreateAlbum(album Album) error {
	_, err := db.Model(album).
		Insert()
	return err
}

func CreateSong(song Song) error {
	_, err := db.Model(song).
		Insert()
	return err
}
