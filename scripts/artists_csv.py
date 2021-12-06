import json
import csv
import sys

bio = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent aliquam consectetur interdum. Mauris nisi justo, accumsan vitae mauris ac, sollicitudin tincidunt lacus. Cras suscipit lorem ex, ac feugiat ipsum rutrum placerat. Maecenas bibendum ipsum non augue consectetur sollicitudin porttitor."

with open('artists.csv', 'w', newline='\n') as csvfile:
    spamwriter = csv.writer(csvfile, delimiter=',')
    for filename in sys.argv[1:]:
        f = open(filename)
        data = json.load(f)

        spamwriter.writerow([str(data['id']), str(data['name']), bio])

        f.close()
