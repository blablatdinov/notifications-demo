import { createStore } from 'vuex'
import auth from './auth'
import notifications from './notifications'
import users from './users'

export default createStore({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    auth,
    notifications,
    users
  }
})
