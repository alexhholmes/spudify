<template>
  <div class="content-page" id="home">
    <h1>Home</h1>

    <div id="artist-list">
      <div
          class="artist-block"
          v-for="(artist, index) in artists"
          :key="index"
          v-on:click="openPageForArtist(artist.name)"
      >
        <img src="../../assets/logo.png" class="artist-img" :alt="artist"/>
        <span>{{artist.name}}</span>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "Home",
  mounted() {
    fetch(`endpoint`, {headers: {}})
    .then(response => response.json())
    .then(data => {
      this.artists = data;
    })
  },
  props: {
    pageChange: Function
  },
  data () {
    return {
      artists: [
        {
          name: "x"
        },
        {
          name: "x"
        },
      ]
    }
  },

  methods: {

    openPageForArtist(artist) {
      this.pageChange(1, artist);
      console.log(artist)
    }
  }

}
</script>

<style scoped>

  h1 {
    margin-bottom: 2rem;
  }
  #home {
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
    grid-row-gap: 1rem;
    grid-column-gap: 1rem;
  }

  #artist-list {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    grid-template-rows: auto 1fr; /* NEW */
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
  }

</style>