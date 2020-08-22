<template>
    <v-app>
        <v-main>
            <Clock/>
            <Quote/>                    
        </v-main>
    </v-app>
</template>

<script>
 import Clock from './components/Clock';
 import Quote from './components/Quote';

 export default {
     name: 'App',
     components: {
         Clock,
         Quote,
     },
 };
</script>
