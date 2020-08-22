<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
          <p>{{ quote }}</p>
          <h3>{{ author }}</h3>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
 export default {
     name: 'Quote',
     data() {
         return {
             quote: "Why can't we all just get along?",
             author: "Rodney King"
         }
     },
     computed: {
         datestr() {
             var time = new Date();
             var timestr = time.getHours() + ":" +
                           time.getMinutes() + ":" +
                           time.getSeconds();
             return timestr
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
