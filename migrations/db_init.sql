CREATE DATABASE spudify;
USE spudify;

CREATE TABLE User {
    user_id varchar(36),
    PRIMARY KEY (user_id),
    name varchar(24),
    age int,
    user_photo_link varchar(64),
    num_hours_listened int,
    num_skips int
}

CREATE TABLE Playlist {
    playlist_id varchar(36),
    PRIMARY KEY (playlist_id),
    name varchar(24),
    photo_link varchar(64),
    owner_id varchar(36),
    FOREIGN KEY (owner_id) REFERENCES User(user_id)
}

CREATE TABLE Artist {
    artist_id varchar(36),
    PRIMARY KEY (artist_id),
    name varchar(24),
    age int,
    bio varchar(300),
    artist_photo_link varchar(64)
}

CREATE TABLE Album {
    album_id varchar(36),
    PRIMARY KEY (album_id),
    title varchar(24),
    genre varchar(24),
    album_cover_link varchar(64),
    artist_id varchar(36),
    FOREIGN KEY (artist_id) REFERENCES Artist(artist_id)
}

CREATE TABLE Song {
    song_id varchar(36),
    PRIMARY KEY (song_id),
    name varchar(24),
    genre varchar(24),
    duration int,
    song_link varchar(64),
    album_id varchar (36),
    FOREIGN KEY (album_id) REFERENCES Album(album_id)
}

CREATE TABLE SongPlaylist {
    song_id varchar(36),
    FOREIGN KEY (song_id) REFERENCES Song(song_id),
    playlist_id varchar(36),
    FOREIGN KEY (playlist_id) REFERENCES Playlist(playlist_id)
}

CREATE TABLE PlaylistFollowers {
    playlist_id varchar(36),
    FOREIGN KEY (playlist_id) REFERENCES Playlist(playlist_id),
    user_id varchar(36),
    FOREIGN KEY (user_id) REFERENCES User(user_id)
}

CREATE TABLE UserFollowers {
    user_id varchar(36),
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    followee_id varchar(36),
    FOREIGN KEY (folowee_id) REFERENCES User(user_id)
}

CREATE TABLE ArtistFollowers {
    user_id varchar(36),
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    artist_id varchar(36),
    FOREIGN KEY (artist_id) REFERENCES Artist(artist_id)
}
