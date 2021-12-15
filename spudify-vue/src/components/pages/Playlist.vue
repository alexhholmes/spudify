<template>
  <div id="playlist">
    <div id="playlist-heading">
      <img id="playlist-img" :src="`https://spudify.nyc3.digitaloceanspaces.com/playlist_images/${playlist.id}.jpg`" alt="" style="background-color: black"/>
      <div id="playlist-info">
        <span>playlist</span>
        <h1>{{playlist.name}}</h1>
        <span>Description: {{playlist.description}}</span>
        <span>{{songs.length}} songs</span>
      </div>
    </div>

    <div id="song-container">
      <div class="song-table-headers" style="margin-bottom: 1rem">
        <span style="margin-left:0">Index</span>
        <span style="margin-left: 31px" >Song</span>
        <span style="margin-left: 800px">Plays</span>
        <span style="margin-left: 430px">Duration</span>
      </div>
      <div
          class="song-row"
          v-for="(song, index) in songs"
          :key="index"
      >
        <div class="song-info">
          <span class="song-counter">{{index + 1}}</span>
          <img class="song-img" :src="`https://spudify.nyc3.digitaloceanspaces.com/album_images/${song.album_id}.jpg`" alt="song">
          <span>{{song.name}}</span>
          <span class="artist-name">{{song.artistName ? song.artistName : ""}}</span>
        </div>

        <span>{{song.plays}}</span>
        <span class="duration-span">{{millisecondsToTime(song.duration)}}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "playlist",
  props: {
    playlist: Object
  },
  mounted() {

    fetch(`http://localhost:8000/api/playlists/${this.playlist.id}`, {headers: {}})
        .then(response => response.json())
        .then(data => {

          let names = []
          for(let i = 0; i < data.length; i++) {
            fetch(`http://localhost:8000/api/artists/${data[i].artist_id}`, {headers: {}})
            .then(resp => resp.json())
            .then(artName => {
              names.push(artName.name);
            })
            .then(() => {
              let tempArr = [...data];
              for (let i = 0; i < data.length; i++) {
                let temp = tempArr[i];
                temp = {...temp, ...{artistName: names[i]}};
                tempArr[i]= temp;
              }
              this.songs = tempArr;
            })
          }
        });
  },
  data () {
    return {
      songs: [
        {
          id:`someid`,
          name: `somename`,
          genre: `somegenre`,
          plays:`someplay`,
          duration:`someduration`,
          artist_id:`some artist id`,
          album_id: `some playlist`
        },
        {
          id:`someid`,
          name: `somename`,
          genre: `somegenre`,
          plays:`someplay`,
          duration:`someduration`,
          artist_id:`some artist id`,
          album_id: `some album`
        },
      ]
    }
  },

  methods: {

    millisecondsToTime(milli) {
      let milliseconds = milli % 1000;
      let seconds = Math.floor((milli / 1000) % 60);
      let minutes = Math.floor((milli / (60 * 1000)) % 60);

      return minutes + ":" + seconds + "." + milliseconds;
    }
  }

}
</script>

<style scoped>

#playlist {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  grid-row-gap: 1rem;
  grid-column-gap: 1rem;
  overflow-y: hidden;
}


#playlist-heading {
  padding: 32px;
  width: 100%;
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  background-color: #686868;
  overflow-y: hidden;
}


#playlist-img {
  height: 200px;
  width: 200px;
  margin-right: 2rem;
}

#playlist-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: space-between;
  height: 250px;
  width: 100%;
}

#song-container {
  padding: 32px;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  background: linear-gradient(180deg, #3F3F3F 0%, rgba(140, 140, 140, 0) 99.99%, rgba(255, 255, 255, 0) 100%);
  overflow: auto;
}

.song-row {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 2rem;
}

.song-info {
  display: grid;
  grid-template-areas: "s a b"
                         "s a c";
}

.song-counter {
  grid-area: s;
  align-self: center;
  justify-self: center;
  margin-right: 4rem;
}

.song-img {
  height: 60px;
  width: 60px;
  margin-right: 1rem;
  grid-area: a;
}

.duration-span {
  min-width: 70px;
}

.song-title {
  grid-area: b;
}

.song-artist {
  grid-area: c;
}

.artist-name {
  min-width: 300px;
}

</style>
