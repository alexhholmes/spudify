#!/usr/bin/env bash
for file in tmp/albums/*
do
    artist_id=$(basename "$file" _albums.json)
    cat $file | jq -r '.[] | "\(.album_id) \(.image_url)"' |
        while IFS=" " read -r album_id url
        do
            wget "$url" -O tmp/albums/${album_id}.jpg
        done
done
