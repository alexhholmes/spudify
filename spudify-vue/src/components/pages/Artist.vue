<template>
  <div class="content-page" id="artist">
    <div id="artist-heading">
      <img id="artist-img" :src="`https://spudify.nyc3.digitaloceanspaces.com/artist_images/${artist.id}.jpg`" alt="" style="background-color: black"/>
      <div id="artist-info">
        <span>Artist</span>
        <h1>{{artist.name}}</h1>
        <span>{{artist.bio}}</span>
      </div>
    </div>
    <div id="album-list">
      <div
          class="album-block"
          v-for="(album, index) in albums"
          :key="index"
          v-on:click="openPageForAlbum(album)"
      >
        <img :src="`https://spudify.nyc3.digitaloceanspaces.com/album_images/${album.id}.jpg`" class="artist-img" :alt="artist"/>
        <span>{{album.title}}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "Artist",
  mounted() {

    fetch(`http://localhost:8000/api/artists/${this.artist.id}/albums`, {headers: {}})
        .then(response => response.json())
        .then(data => {
          this.albums = data;
        });
  },
  props: {
    artist: Object,
    pageChange: Function
  },
  data () {
    return {
      albums: [
        {
          id:`39djx`,
          title: `Some Title`,
          genre: `Some Genre`,
          artist_id:`someId`,
        }
      ]
    }
  },
  methods: {
    openPageForAlbum(album) {
      this.pageChange(2, album);
    }
  }
}
</script>

<style scoped>
  #artist {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    grid-row-gap: 1rem;
    grid-column-gap: 1rem;
    overflow-y: auto;
  }

  #artist-heading {
    padding: 32px;
    width: 100%;
    display: flex;
    align-items: flex-start;
    justify-content: flex-start;
    background-color: #686868;
    margin-bottom: 2rem;
  }

  #artist-img {
    height: 250px;
    width: 250px;
    margin-right: 2rem;
  }

  #artist-info {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: space-between;
    height: 250px;
    width: 100%;
  }

  #album-list {
    padding: 32px;
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    grid-template-rows: auto 1fr; /* NEW */
    grid-column-gap: 1rem;
    grid-row-gap: 4rem;
    align-items: stretch;
    background: linear-gradient(180deg, #3F3F3F 0%, rgba(140, 140, 140, 0) 99.99%, rgba(255, 255, 255, 0) 100%);
  }
  .album-block {
    height: 200px;
    width: 150px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;
    cursor: pointer;
  }

  .artist-img {
    height: 150px;
    width: 150px;
    border-radius: 50%;
    object-fit: scale-down;
  }
</style>