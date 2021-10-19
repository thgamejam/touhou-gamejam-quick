import { createRouter, createWebHistory } from 'vue-router'
import Home from '../App.vue'

const routes = [{
  path: '/',
  name: 'Home',
  component: Home
},
// {
//   path: '/maintain',
//   name: 'About',
//   component: () =>
//     import( /* webpackChunkName: "about" */ '../views/Maintain.vue')
// },
// {
//   path: '/root_home',
//   name: 'RootHome',
//   component: RootHome,
//   children: [{
//     path: "/root_home/add_user",
//     name: "add-user",
//     component: () =>
//       import("../components/root/page/AddUserPage.vue")
//   },
//   ]
// }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router