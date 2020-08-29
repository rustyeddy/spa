<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
          <p>{{ text }}</p>
          <h4>{{ title }}</h4>
          <h3>{{ author }}</h3>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
 export default {
     name: 'Quote',
     computed: {
         text() {
             return this.$store.state.quote.quote.text
         },
         title() {
             return this.$store.state.quote.quote.title
         },
         author() {
             return this.$store.state.quote.quote.author
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
