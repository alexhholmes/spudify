<template>
  <div id="main-container">

    <v-navigation-drawer permanent id="side-nav">
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6">
            Spudify
          </v-list-item-title>
          <v-list-item-subtitle>
            Spotify but worse...
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
          dense
          nav
      >
        <v-list-item
            v-for="item in items"
            :key="item.title"
            link
            v-on:click="goHome(item.title)"
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <div id="content-container">
      <Home v-if="page === 0" :page-change="pageChange" :show-artist="showArtists"/>
      <Artist v-else-if="page === 1" :page-change="pageChange" :artist="selectedArtist"/>
      <Album v-else-if="page === 2" :album="selectedAlbum"/>
      <Playlist v-else-if="page === 3" :playlist="selectedPlaylist"/>
    </div>

    <div id="footer-container">

    </div>
    <v-dialog
        v-model="showLoginDialog"
        persistent
        max-width="290"
    >
      <div id="login-dialog-container">

        <v-text-field label="username" v-model="username"/>
        <v-text-field label="password" v-model="password"/>
        <v-btn color="#1DB954" v-on:click="attemptLogin">login</v-btn>
      </div>
    </v-dialog>
  </div>
</template>

<script>
import Home from "@/components/pages/Home";
import Artist from "@/components/pages/Artist";
import {checkCookie} from "@/util/CookieHelper";
import Album from "@/components/pages/Album";
import Playlist from "@/components/pages/Playlist";

export default {
  name: "Dashboard",
  components: {Playlist, Album, Artist, Home},

  mounted() {
    // TODO: Remove ! for this to work properly
    if (checkCookie("session")) {
      this.showLoginDialog = false;
      // TODO: Check if cookie is actually valid
    } else {
      // TODO make true
      this.showLoginDialog = true;
    }
  },
  data () {
    return {
      showArtists: true,
      page: 0,
      username: "",
      password: "",
      showLoginDialog: false,
      selectedArtist: false,
      selectedAlbum: null,
      selectedPlaylist: null,
      items: [
        { title: 'Artists', icon: 'mdi-account-music' },
        { title: 'Playlists', icon: 'mdi-playlist-music' },
      ],
      right: null,
    }
  },

  methods: {
    goHome(title) {
      this.selectedAlbum = null;
      this.selectedArtist = null;
      this.selectedPlaylist = null;
      this.showArtists = title === "Artists";
      this.pageChange(0, {})
    },

    attemptLogin () {

      // Request payload
      let loginData = {
        username: this.username,
        password: this.password
      }
      console.log(loginData);

      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: "username=" + loginData.username + "&password=" + loginData.password
      };
      fetch(`http://localhost:8000/api/login`, requestOptions)
      .then(response => {
        console.log(response.status)
        if (response.ok) {
          this.showLoginDialog = false;
        } else {
          // TODO: If unsuccessful
          this.username = "";
          this.password = "";
        }
      })
    },

    pageChange(pageID, data) {
      this.page = pageID;
      if (pageID === 1) {
        this.selectedArtist = data;
      }
      else if (pageID === 2) {
        this.selectedAlbum = data;
      } else if (pageID === 3) {
        this.selectedPlaylist = data;
      }
    }
  }

}
</script>

<style scoped>
  #main-container {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: 10% 90%;
    grid-template-rows: 90% 10%;
    grid-template-areas: "side content"
                         "footer footer";
  }

  #content-container {
    grid-area: content;
    width: 100%;
    background-color: #686868;
  }

  #footer-container {
    grid-area: footer;
    background-color: #434343;
  }

  #login-dialog-container {
    padding: 16px;
    background-color: #2f2f2f;
  }

</style>