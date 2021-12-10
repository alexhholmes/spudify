package models

type Artist struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
}

type Album struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Genre     string    `json:"genre"`
	ArtistID  string	`json:"artist_id"`
}

type Song struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Genre     string    `json:"genre"`
	Plays     int       `json:"plays"`
	Duration  int       `json:"duration"`
	ArtistID  string    `json:"artist_id"`
	AlbumID   string    `json:"album_id"`
}

type User struct {
	ID                string    `json:"id"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	NumHoursListened  int       `json:"num_hours_listened"`
	NumSkips          int       `json:"num_skips"`
	Session           string    `json:"session"`
}

type Playlist struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
}