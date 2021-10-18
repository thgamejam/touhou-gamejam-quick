import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import './assets/css/global.css'
import 'element-plus/dist/index.css'
import App from './App.vue'
import VueRouter from 'vue-router'

const app = createApp(App)

app.use(ElementPlus)
app.use(VueRouter)
app.mount('#app')