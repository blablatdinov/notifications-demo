import axios from "axios"

export default {
  namespaced: true,
  state: () => {
    users: []
  },
  mutations: {
    SET_USERS(state, users) {
      state.users = users
    }
  },
  actions: {
    getUsersList({commit}) {
      axios.get('/api/v1/users/')
        .then((response) => commit('SET_USERS', response.data.result))
    }
  },
}