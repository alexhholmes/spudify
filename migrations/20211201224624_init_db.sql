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

CREATE TABLE songs (
    id varchar(22),
    PRIMARY KEY (id),
    name varchar(24) NOT NULL,
    genre varchar(24),
    duration int NOT NULL,
    artist_id varchar(22) NOT NULL,
    FOREIGN KEY (artist_id) REFERENCES artists(id),
    album_id varchar(22) NOT NULL,
    FOREIGN KEY (album_id) REFERENCES albums(id)
);

-- Initial DB data
INSERT INTO artists
VALUES ('p7ddcTyzvAtXkxkq6NyQT', 'Bob Joe', 21, 'I like music');

INSERT INTO artists
VALUES ('pf6CtmEXNMyFxMSzrPDma', 'Billy Joe', 23, 'I dislike music');

INSERT INTO artists
VALUES ('3npJyksSq8DzAXyJR9Yq3', 'Joe Lewis', 27, 'I like most music');

INSERT INTO artists
VALUES ('uvkrapMyTdzj2wk1SRU5L', 'Sally Bob', 22, 'I like some music');

