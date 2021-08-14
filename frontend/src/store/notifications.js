import axios from 'axios'

export default {
  namespaced: true,
  state: {
    notifications: []
  },
  mutations: {
    SET_NOTIFICATIONS(state, notifications) {
      if (!notifications) {
        return
      }
      state.notifications = state.notifications.concat(notifications)
    },
    DELETE_NOTIFICATION(state, notificationId) {
      let newArr = state.notifications.filter(elem => {
        return elem.id != notificationId
      })
      state.notifications = newArr
    }
  },
  actions: {
    getNotifications({commit}) {
      axios.get('/api/v1/notifications/')
        .then((response) => {
          commit('SET_NOTIFICATIONS', response.data.result)
        })
    },
    sendNotification({commit}, data) {
      axios.post('/api/v1/notifications/', data)
        .then((response) => {
        })
    },
    deleteNotification({commit}, notificationId) {
      axios.delete(`/api/v1/notifications/${notificationId}`)
      commit('DELETE_NOTIFICATION', notificationId)
    }
  },
}