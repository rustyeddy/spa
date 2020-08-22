import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import VueNativeSock from 'vue-native-websocket'

Vue.use(VueNativeSock, 'ws://localhost:1233/ws', { 
  store: store, 
  reconnection: true,
  reconnectionAttempts: 5,
  reconnectionDelay: 3000,
})

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
