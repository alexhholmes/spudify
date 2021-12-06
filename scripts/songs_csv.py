import json
import csv
import sys
import os

with open('songs.csv', 'w', newline='\n') as csvfile:
    spamwriter = csv.writer(csvfile, delimiter=',')
    for filename in sys.argv[1:]:
        f = open(filename)
        data = json.load(f)
        base = os.path.basename(filename)
        base = os.path.splitext(base)[0]
        artist_id = base[0:22]
        album_id = base[23:45]


        for song in data:
            spamwriter.writerow([str(song['song_id']), str(song['name']), "rock", 0, song['duration'], artist_id, album_id])

        f.close()
