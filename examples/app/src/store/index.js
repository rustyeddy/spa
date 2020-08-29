import Vue from 'vue'
import Vuex from 'vuex'

import * as websocket from '@/store/modules/ws.js'
import * as time from '@/store/modules/time.js'
import * as quote from '@/store/modules/quote.js'

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        websocket,
        time,
        quote
    },
    state: {
    },
    mutations: {
    },
    actions: {
    },
})
