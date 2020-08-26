<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
          <h1>{{ timeStr  }}</h1>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
 export default {
     name: 'Clock',
     computed: {
         timeStr() {
             return this.$store.state.time
         }
     }
 }
</script>

<style :scoped>
 body {
     background: black;
 }
 
 h1 {
     font-size: 96px;
     color: #2085a3;

 }
</style>
