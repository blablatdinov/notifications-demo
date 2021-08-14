<template>
<div>
  <h1>Авторизация</h1>
  <form>
    <label for="username">Имя пользователя:</label>
    <input type="text" id="username" v-model="username">
    <label for="password">Пароль:</label>
    <input type="text" id="password" v-model="password">
    <ul>
      <li v-for="message in $store.state.auth.messages" :key="message">{{message}}</li>
    </ul>
    <button @click.prevent="login">auth</button>
  </form>
</div>
</template>

<script>
export default {
  data() {
    return {
      username: 'blablatdinov',
      password: '1',
    }
  },
  methods: {
    login() {
      this.$store.dispatch('generateToken', this.username, this.password)
    }
  }
}
</script>

<style>

</style>