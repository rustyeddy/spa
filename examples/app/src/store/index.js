import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    socket: {
      isConnected: false,
      message: '',
      reconnectError: false,
    }
  },
  mutations: {
    SOCKET_ONOPEN(state, event)  {
      console.log("ONOPEN: ", event)
      state.socket.isConnected = true
    },
    SOCKET_ONCLOSE(state, event)  {
      console.log("ONCLOSE", event)
      state.socket.isConnected = false
    },
    SOCKET_ONERROR(state, event)  {
      console.error(state, event)
    },
    SOCKET_RECONNECT(state, count) {
      console.info(state, count)
    },
    SOCKET_RECONNECT_ERROR(state) {
      state.socket.reconnectError = true;
    }
  },
  actions: {
  },
  modules: {
  }
})
