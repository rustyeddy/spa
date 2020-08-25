import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        time: "",
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
        SOCKET_ONMESSAGE(state, event) {
            console.log("ONMESSAGE", event)
            state.socket.message = event;
        },
        SOCKET_ONERROR(state, event)  {
            console.error(state, event)
        },
        SOCKET_RECONNECT(state, count) {
            console.info(state, count)
        },
        SOCKET_RECONNECT_ERROR(state) {
            state.socket.reconnectError = true;
        },
        SET_TIME(state, t) {
            state.time = t
        }
    },
    actions: {
        handleTimeChange: function({ dispatch }, msg) {
            dispatch('settime', msg)
        }
    },
    modules: {
    }
})
