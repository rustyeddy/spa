import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        time: {
            year: 0,
            month: 0,
            day: 0,
            hour: 0,
            minute: 0,
            second: 0,
        },
        quote: {
            title: "",
            text: "Why can't we all just get along?",
            author: "Rodney King"
        },
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
            console.log("ONMESSAGE", event.data)
            state.socket.message = event.data;
            console.log("message: ", state.socket.message)
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
            state.time.year = t.year
            state.time.month = t.month
            state.time.day = t.day
            state.time.hour = t.hour
            state.time.minute = t.minute
            state.time.second = t.second
        },

        SET_QUOTE(state, q) {
            state.quote.title = q.title
            state.quote.author = q.author
            state.quote.text = q.text
        }
    },
    actions: {
        setTime({ commit }, msg) {
            console.log("setTime: ", msg)
            commit('SET_TIME', msg)
        },

        setQuote({ commit }, msg) {
            console.log("setQuote: ", msg)
            commit('SET_QUOTE', msg)
        }
    },
    modules: {
    }
})
