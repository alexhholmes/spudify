-- Add migration script here

-- ID's use nanoid (length 22 symbols)
CREATE TABLE artists (
    id varchar(22),
    PRIMARY KEY (id),
    name varchar(24) NOT NULL,
    age int,
    bio varchar(300) NOT NULL
);

CREATE TABLE albums (
    id varchar(22),
    PRIMARY KEY (id),
    title varchar(24) NOT NULL,
    genre varchar(24) NOT NULL,
    artist_id varchar(22) NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id)
);

CREATE INDEX ON albums (artist_id);

CREATE TABLE songs (
    id varchar(22),
    PRIMARY KEY (id),
    name varchar(24) NOT NULL,
    genre varchar(24),
    plays int NOT NULL,
    duration int NOT NULL,
    artist_id varchar(22) NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id),
    album_id varchar(22) NOT NULL,
    FOREIGN KEY (album_id) REFERENCES albums(id)
);

CREATE INDEX ON songs(artist_id);
CREATE INDEX ON songs(album_id);

CREATE TABLE users (
    id varchar(22),
    PRIMARY KEY (id),
    username varchar(24) NOT NULL,
    password varchar(24) NOT NULL,
    num_hours_listened int NOT NULL,
    num_skips int NOT NULL
);

CREATE INDEX ON users(username);

CREATE TABLE playlists (
    id varchar(22),
    PRIMARY KEY (id),
    name varchar(24) NOT NULL,
    description varchar(300),
    owner_id varchar(22),
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE INDEX ON playlists(owner_id);

CREATE TABLE songs_playlists (
    song_id varchar(22) NOT NULL,
    FOREIGN KEY (song_id) REFERENCES songs(id),
    playlist_id varchar(22) NOT NULL,
    FOREIGN KEY (playlist_id) REFERENCES playlists(id)
);

CREATE INDEX ON songs_playlists(song_id);
CREATE INDEX ON songs_playlists(playlist_id);

CREATE TABLE artist_followers (
    user_id varchar(22) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    artist_id varchar(22) NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id)
);

CREATE INDEX ON artist_followers(user_id);
CREATE INDEX ON artist_followers(artist_id);

CREATE TABLE liked_songs (
    user_id varchar(22) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    song_id varchar(22) NOT NULL,
    FOREIGN KEY (song_id) REFERENCES songs(id)
);

CREATE INDEX ON liked_songs(user_id);
CREATE INDEX ON liked_songs(song_id);

-- Initial DB data
INSERT INTO artists
VALUES ('p7ddcTyzvAtXkxkq6NyQT', 'Bob Joe', 21, 'I like music');

INSERT INTO artists
VALUES ('pf6CtmEXNMyFxMSzrPDma', 'Billy Joe', 23, 'I dislike music');

INSERT INTO artists
VALUES ('3npJyksSq8DzAXyJR9Yq3', 'Joe Lewis', 27, 'I like most music');

INSERT INTO artists
VALUES ('uvkrapMyTdzj2wk1SRU5L', 'Sally Bob', 22, 'I like some music');
