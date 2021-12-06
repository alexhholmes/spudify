<template>
  <div id="album">
    <div id="album-heading">
      <img id="album-img" src="../../assets/logo.png" alt="" style="background-color: black"/>
      <div id="album-info">
        <span>Album</span>
        <h1>{{album.title}}</h1>
        <span>{{album.genre}}</span>
        <span>{{songs.length}} songs</span>
      </div>
    </div>

    <div id="song-container">
      <div
          class="song-row"
          v-for="(song, index) in songs"
          :key="index"
      >
        <div class="song-info">
          <span class="song-counter">{{index + 1}}</span>
          <img class="song-img" src="../../assets/logo.png" alt="song">
          <span>{{song.name}}</span>
          <span>{{song.artist_id}}</span>
        </div>

        <span>{{song.plays}}</span>
        <span>{{song.duration}}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "Album",
  props: {
    album: Object
  },
  mounted() {

    fetch(`/albums/${this.album.id}/songs`, {headers: {}})
        .then(response => response.json())
        .then(data => {
          this.songs = data;
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
          album_id: `some album`
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
}
</script>

<style scoped>

  #album {
    height: 70%;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    grid-row-gap: 1rem;
    grid-column-gap: 1rem;
    overflow-y: auto;
  }


  #album-heading {
    padding: 32px;
    width: 100%;
    display: flex;
    align-items: flex-start;
    justify-content: flex-start;
    background-color: #686868;
  }


  #album-img {
    height: 250px;
    width: 250px;
    margin-right: 2rem;
  }

  #album-info {
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

  .song-title {
    grid-area: b;
  }

  .song-artist {
    grid-area: c;
  }

</style>