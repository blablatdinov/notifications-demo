import { createApp } from 'vue'
import axios from 'axios'
import App from './App.vue'
import router from './router'
import store from './store'

createApp(App).use(store).use(router).mount('#app')

axios.defaults.baseURL = 'http://localhost:8000';
axios.interceptors.request.use(
  function (config) {
    const { headers } = config
    const token = localStorage.getItem('token')
    headers['Authorization'] = `Bearer ${token}`
    return config
  },
  function (error) {
    return Promise.reject(error)
  }
)
// axios.interceptors.response.use(
//   function (response) {
//     return response
//   },
//   function (error) {
//     if (error.status === 401) {
//       router.push('/login')
//     }
//   }
// )
