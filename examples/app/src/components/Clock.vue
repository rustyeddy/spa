<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">

          {{ datestr }}

      </v-col>
    </v-row>
  </v-container>
</template>

<script>
 export default {
     name: 'Clock',
     data() {
         return {
             datetime: "10:10 10/10/2020"
         }
     },
     computed: {
         datestr() {
             var time = new Date();
             var timestr = time.getHours() + ":" + time.getMinutes() + ":" + time.getSeconds();
             return timestr
         }   
     }
 }
</script>
