import {createRouter, createWebHistory} from 'vue-router'
import Home from '../App.vue'
import Game from "@/views/page/Game.vue";
import Test from '@/views/component/GameExhibit.vue'
import Test2 from '@/views/component/AuthorInfo.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/game/:id',
    component: Game
  },
  {
    path: '/test',
    component: Test
  },
  {
    path: '/test2',
    component: Test2
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router