import {createRouter, createWebHistory} from 'vue-router'
// import Home from '../App.vue'
import Game from "@/views/page/Game.vue";
import GameJam from "@/views/page/GameJam.vue";

const routes = [
  {
    path: '/',
    component: GameJam
  },
  {
    path: '/game/:id',
    component: Game
  },
  {
    path: '/gamejam',
    component: GameJam
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router