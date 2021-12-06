<template>
  <div id="main-container">

    <v-navigation-drawer permanent id="side-nav">
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6">
            Application
          </v-list-item-title>
          <v-list-item-subtitle>
            subtext
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
      <Home v-if="page === 0" :page-change="pageChange"/>
      <Artist v-else-if="page === 1"/>
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

export default {
  name: "Dashboard",
  components: {Artist, Home},
  mounted() {
    if (checkCookie("auth")) {
      console.log("hi")
    } else {
      this.showLoginDialog = true;
    }
  },
  data () {
    return {
      page: 0,
      username: "",
      password: "",
      showLoginDialog: false,
      selectedArtist: null,
      items: [
        { title: 'Dashboard', icon: 'mdi-view-dashboard' },
        { title: 'Photos', icon: 'mdi-image' },
        { title: 'About', icon: 'mdi-help-box' },
      ],
      right: null,
    }
  },

  methods: {
    pageChange(pageID, data) {
      this.page = pageID;
      if (pageID === 1) {
        this.selectedArtist = data;
      }
      console.log(data);
    }
  }

}
</script>

<style scoped>
  #main-container {
    height: 100%;
    width: 100%;
    display: grid;
    grid-template-columns: auto 100%;
    grid-template-rows: 90% 10%;
    grid-template-areas: "side content"
                         "footer footer";
  }

  #content-container {
    padding: 32px;
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
  }

</style>