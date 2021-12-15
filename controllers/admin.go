package controllers

import (
	. "example/spudify/models"
)

func CreateArtist(artist Artist) (string, error) {
	_, err := db.Model(&artist).
		Insert()
	return artist.ID, err
}

func CreateAlbum(album Album) (string, string, error) {
	_, err := db.Model(&album).
		Insert()
	return album.ID, album.ArtistID, err
}

func CreateSong(song Song) (string, error) {
	_, err := db.Model(&song).
		Insert()
	return song.ID, err
}
