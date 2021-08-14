import axios from 'axios'

export default {
  state: {
    messages: [],
  },
  mutations: {
    SET_MESSAGE(state, messageText) {
      state.messages = [messageText]
    }
  },
  actions: {
    async generateToken({ commit }, username, password) {
      let data = {username: username, password: password}
      await axios.post('/auth/sign-in', data)
        .then((response) => {
          localStorage.setItem('token', response.data.token)
        })
        .catch((error) => {
          if (error.response.data.message === 'sql: no rows in result set') {
            commit('SET_MESSAGE', 'Имя пользователя или пароль не верны')
          }
        })
    },
    async signUp({ commit }, username, password) {
      let data = {username: username, password: password}
      await axios.post('/auth/sign-up', data)
        .then((response) => {
        })
        .catch((error) => {
          if (error.response.data.message === 'sql: no rows in result set') {
            commit('SET_MESSAGE', 'Имя пользователя или пароль не верны')
          } else if (error.response.data.message === 'pq: duplicate key value violates unique constraint "users_username_key"') {
            commit('SET_MESSAGE', 'Пользователь с таким именем уже существует')
          }
        })
    }
  },
}
