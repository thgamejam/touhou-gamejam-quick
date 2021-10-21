import {createRouter, createWebHistory} from 'vue-router'
// import Home from '../App.vue'
import Game from "@/views/page/Game.vue";
import GameJam from "@/views/page/GameJam.vue";
import Test from '@/views/component/GameExhibit.vue'
import Test2 from '@/views/component/AuthorInfo.vue'
import Test3 from '@/views/component/GameDownload.vue'

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
  },
  {
    path: '/test',
    component: Test
  },
  {
    path: '/test2',
    component: Test2
  },
  {
    path: '/test3',
    component: Test3
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router