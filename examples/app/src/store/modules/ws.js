export const state = {
    socket: {
        isConnected: false,
        message: '',
        reconnectError: false,
    }
}

export const mutations = {
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


}
