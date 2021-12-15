<template>
  <div class="content-page" id="home">
    <h1>{{showArtist ? "Artists" : "Playlists"}}</h1>

    <div v-if="showArtist" class="artist-list">
      <div
          class="artist-block"
          v-for="(artist, index) in artists"
          :key="index + 'artist'"
          v-on:click="openPageForArtist(artist)"
      >
        <img :src="`https://spudify.nyc3.digitaloceanspaces.com/artist_images/${artist.id}.jpg`" class="artist-img" :alt="artist"/>
        <span>{{artist.name}}</span>
      </div>
    </div>

    <!-- Populate the page with playlists from the playlists  -->
    <div v-else class="artist-list">
      <div
          class="artist-block"
          v-for="(playlist, index) in playlists"
          :key="index + 'playlist'"
          v-on:click="openPageForPlaylist(playlist)"
      >
        <img :src="`https://spudify.nyc3.digitaloceanspaces.com/playlist_images/${playlist.id}.jpg`" class="artist-img" :alt="playlist"/>
        <span>{{playlist.name}}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  mounted() {
    this.getData();
  },
  watch: {
    showArtist: function() { // watch it
      this.getData();
    }
  },
  props: {
    pageChange: Function,
    showArtist: Boolean,
  },
  data () {
    return {
      artists: [
        {
          name: "x",
          bio: "aasdfasdf",
          albums: []
        },
        {
          name: "x",
          bio: "asdfasdf",
          albums: []
        },
      ],
      playlists: [
        {
          id: "aasdfasdf",
          name: "x",
          description: "schooop"
        },
      ]
    }
  },

  methods: {

    getData()  {
      let endpoint;

      // Gets artists or playlists to show depending on the page the user selected
      if (this.showArtist){
        endpoint = 'http://localhost:8000/api/artists';
      } else {
        endpoint = 'http://localhost:8000/api/me/playlists';
      }

      fetch(endpoint, {headers:  {
        }})
          .then(response => response.json())
          .then(data => {
            console.log(data);
            console.log("fetch  data"  + endpoint)
            if (this.showArtist) this.artists = data;
            else this.playlists = data;
          })
    },

    openPageForArtist(artist) {
      this.pageChange(1, artist);
    },

    openPageForPlaylist(playlist){
      this.pageChange(3, playlist);
    }
  }

}
</script>

<style scoped>

  h1 {
    margin-bottom: 2rem;
  }
  #home {
    padding: 32px;
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

  .artist-list {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    grid-template-rows: auto 1fr;
    grid-column-gap: 1rem;
    grid-row-gap: 4rem;
    align-items: stretch;
  }

  .artist-block {
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
    box-shadow: 0 2px 2px 0 rgba(0,0,0,0.14), 0 3px 1px -2px rgba(0,0,0,0.12), 0 1px 5px 0 rgba(0,0,0,0.20);
  }

</style>