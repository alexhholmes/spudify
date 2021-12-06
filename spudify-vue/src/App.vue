<template>
  <v-app>
    <v-main>
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
  .content-page {
    height: 100%;
    width: 100%;
  }

</style>