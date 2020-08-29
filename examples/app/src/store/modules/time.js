export const state = {
    time: {
        year: 0,
        month: 0,
        day: 0,
        hour: 0,
        minute: 0,
        second: 0,
    },
}

export const mutations = {
    SET_TIME(state, t) {
        state.time.year = t.year
        state.time.month = t.month
        state.time.day = t.day
        state.time.hour = t.hour
        state.time.minute = t.minute
        state.time.second = t.second
    },
}

export const actions = {
    setTime({ commit }, msg) {
        console.log("setTime: ", msg)
        commit('SET_TIME', msg)
    },
}
