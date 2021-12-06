#!/usr/bin/env bash
# USAGE ./scrape_spotify.sh [file_path] [Spotify OAuth Token]
input="$1"
while IFS= read -r line
do
    # Get artist json
    curl -X "GET" "https://api.spotify.com/v1/artists/$line" -H "Accept: application/json" -H "Content-Type: application/json" -H "Authorization: Bearer $2" -o tmp/raw/raw_${line}.json
    cat tmp/raw/raw_${line}.json | jq '{id: .id, name: .name, image_url: .images[0].url}' > tmp/artists/${line}.json

    # Get artist albums json
    curl -X "GET" "https://api.spotify.com/v1/artists/$line/albums?include_groups=album&market=ES" -H "Accept: application/json" -H "Content-Type: application/json" -H "Authorization: Bearer $2" -o tmp/raw/raw_${line}_albums.json
    cat tmp/raw/raw_${line}_albums.json | jq '[.items[] | {album_id: .id, title: .name, image_url: .images[0].url}]' > tmp/albums/${line}_albums.json

    # Get album songs json
    cat tmp/albums/${line}_albums.json | jq -r '.[] | .album_id' |
        while IFS= read -r album_id
        do
            curl -X "GET" "https://api.spotify.com/v1/albums/${album_id}/tracks?market=ES" -H "Accept: application/json" -H "Content-Type: application/json" -H "Authorization: Bearer $2" -o tmp/raw/raw_${line}_${album_id}_songs.json
            cat tmp/raw/raw_${line}_${album_id}_songs.json | jq '[.items[] | {song_id: .id, name: .name, duration: .duration_ms}]' > tmp/songs/${line}_${album_id}_songs.json
        done
done < "$input"
