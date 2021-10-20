import {createRouter, createWebHistory} from 'vue-router'
import Home from '../App.vue'
import Game from "@/views/page/Game.vue";

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/game/:id',
    component: Game
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router