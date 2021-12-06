<template>
  <div class="content-page" id="artist">
    <div id="artist-heading">
      <img id="artist-img" src="" alt="" style="background-color: black"/>
      <div id="artist-info">
        <h4>Artist</h4>
        <h1>{{artist.name}}</h1>
        <h4>{{artist.bio}}</h4>
        <h4>{{artist.albums.length}}</h4>
      </div>
    </div>
    <h1>Artist</h1>

    <div id="song-list">
      <div
          class="song-container"
          v-for="(artist, index) in artists"
          :key="index"
      >
        <img src="../../assets/logo.png" class="artist-img" :alt="artist"/>
        <span>{{artist.name}}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "Artist",
  mounted() {
  },
  props: {
    artist: Object
  },
  data () {
    return {
      songs: []
    }
  }
}
</script>

<style scoped>
  #artist-heading {
    height: 500px;
    width: 100%;
    display: flex;
    align-items: flex-start;
    justify-content: flex-start;
  }

</style>