#!/usr/bin/env bash
for file in tmp/artists/*
do
    id=$(cat $file | jq -r '.id')
    url=$(cat $file | jq -r '.image_url')
    wget "$url" -O tmp/artists/${id}.jpg
done
