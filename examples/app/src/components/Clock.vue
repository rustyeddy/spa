<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
          <h1>{{ hour }}:{{ minute }}:{{ second }}</h1>
          <h2>{{ month }}/{{ day }}/{{ year }}</h2>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
 export default {
     name: 'Clock',
     computed: {
         year() {
             return this.$store.state.time.time.year
         },
         month() {
             return this.$store.state.time.time.month
         },
         day() {
             return this.$store.state.time.time.day
         },
         hour() {
             return this.$store.state.time.time.hour
         },
         minute() {
             return this.$store.state.time.time.minute
         },
         second() {
             return this.$store.state.time.time.second
         },
         timeStr() {
             return this.$store.state.time.time
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
