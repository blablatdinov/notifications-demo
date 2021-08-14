<template>
  <div>
    <h1>Уведомления</h1>
    <ul>
      <li 
        v-for="notification in $store.state.notifications.notifications" 
        :key="notification.id"
      >
        {{notification.text}}
        <button @click="deleteNotification(notification.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import {mapActions, mapMutations} from 'vuex'

export default {
  data () {
    return {
      socket: null
    }
  },
  methods: {
    ...mapActions({
      getNotifications: 'notifications/getNotifications',
      deleteNotification: 'notifications/deleteNotification'
    }),
    ...mapMutations({
      setNotification: 'notifications/SET_NOTIFICATIONS'
    }),
    connectToSocket() {
      this.socket = new WebSocket(`ws://localhost:8000/ws/?token=${localStorage.getItem("token")}`)
      this.socket.onmessage = event => {
        this.$store.commit('notifications/SET_NOTIFICATIONS', {id: 0, text: event.data})
      }
    }
  },
  created() {
    this.getNotifications()
    this.connectToSocket()
  }
}
</script>

<style>

</style>