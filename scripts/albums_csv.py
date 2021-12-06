import json
import csv
import sys
import os

with open('albums.csv', 'w', newline='\n') as csvfile:
    spamwriter = csv.writer(csvfile, delimiter=',')
    for filename in sys.argv[1:]:
        f = open(filename)
        data = json.load(f)
        artist_id = os.path.basename(filename)
        artist_id = os.path.splitext(artist_id)[0]
        artist_id = artist_id[:-7]

        for album in data:
            spamwriter.writerow([str(album['album_id']), str(album['title']), "rock", artist_id])

        f.close()
