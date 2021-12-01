-- Add migration script here
CREATE TABLE artists (
    id varchar(36),
    PRIMARY KEY (id),
    name varchar(24),
    age int,
    bio varchar(300)
);

CREATE TABLE albums (
    id varchar(36),
    PRIMARY KEY (id),
    title varchar(24),
    genre varchar(24),
    artist_id varchar(36),
    FOREIGN KEY (artist_id) REFERENCES artists(id)
);

CREATE TABLE songs (
    id varchar(36),
    PRIMARY KEY (id),
    name varchar(24),
    genre varchar(24),
    duration int,
    artist_id varchar(36),
    FOREIGN KEY (artist_id) REFERENCES artists(id),
    album_id varchar(36),
    FOREIGN KEY (album_id) REFERENCES albums(id)
);
