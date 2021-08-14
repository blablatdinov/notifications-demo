<template>
<div>
  {{ selectedUsers }}
  <form>
    <select v-model="selectedUsers" multiple>
      <option v-for="user in $store.state.users.users" :key="user.id">
        {{ user.username }}
      </option>
    </select>
    <input type="text" v-model="text">
    {{text}}
    <button @click.prevent="sendNotification">Отправить</button>
  </form>

</div>
</template>

<script>
export default {
  data() {
    return {
      selectedUsers: [],
      text: '',
    }
  },
  methods: {
    getUsersList() {
      this.$store.dispatch('users/getUsersList')
    },
    sendNotification() {
      this.$store.dispatch('notifications/sendNotification', {username: this.selectedUsers, message: this.text})
    }
  },
  created() {
    this.getUsersList()
  }
}
</script>

<style>

</style>