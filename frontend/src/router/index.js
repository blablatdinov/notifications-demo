import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Main',
    component: () => import(/* webpackChunkName: "main" */ '@/views/Main.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import(/* webpackChunkName: "login" */ '@/views/Login.vue')
  },
  {
    path: '/reg',
    name: 'Registration',
    component: () => import(/* webpackChunkName: "registration" */ '@/views/Registration.vue')
  },
  {
    path: '/send',
    name: 'SendNotification',
    component: () => import(/* webpackChunkName: "send" */ '@/views/SendNotification.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
