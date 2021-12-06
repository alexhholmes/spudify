#!/usr/bin/env bash
for file in tmp/songs/*
do
    cat $file | jq -r '.[] | .song_id' |
        while IFS= read -r song_id
        do
            cp royalty_free_music_sample.mp3 ${song_id}.mp3
            s3cmd put ${song_id}.mp3 s3://spudify/songs/${song_id}.mp3
            rm ${song_id}.mp3
        done
done
