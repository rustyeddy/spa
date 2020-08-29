export const state = {
    quote: {
        title: "",
        text: "Why can't we all just get along?",
        author: "Rodney King"
    },
}

export const actions = {
    setQuote({ commit }, msg) {
        console.log("setQuote: ", msg)
        commit('SET_QUOTE', msg)
    }
}

export const mutations = {
    SET_QUOTE(state, q) {
        state.quote.title = q.title
        state.quote.author = q.author
        state.quote.text = q.text
    }
}
