<template>
    <v-app>
        <v-main>
            <Clock/>
        </v-main>
    </v-app>
</template>

<script>
 import Clock from './components/Clock';

 export default {
     name: 'App',

     components: {
         Clock,
     },
 };
</script>
