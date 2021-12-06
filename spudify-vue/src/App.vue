<template>
  <v-app>
    <v-main style="width: 100vw; height: 100vh; max-width: 100%" >
     <Dashboard/>
    </v-main>
  </v-app>
</template>

<script>
import Dashboard from "@/components/Dashboard";

export default {
  name: 'App',
  components: {
    Dashboard,
  },

  data: () => ({
    //
  }),
};
</script>

<style>
  * {
    box-sizing: border-box;

  }

  *::-webkit-scrollbar {
    width: 12px;               /* width of the entire scrollbar */
  }

  *::-webkit-scrollbar-track {
    background: rgba(0,0,0, 0.7);        /* color of the tracking area */

  }

  *::-webkit-scrollbar-thumb {
    background-color: black;    /* color of the scroll thumb */
    border-radius: 20px;       /* roundness of the scroll thumb */
    border: 3px solid rgba(0,0,0, 0.1);  /* creates padding around scroll thumb */
  }
  .content-page {
    height: 100%;
    width: 100%;
  }

</style>