import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RegisterView from '@/views/user/RegisterView.vue'
import LoginView from '@/views/user/LoginView.vue'
import '@/assets/css/global.css'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
        },
        {
            path: '/register',
            name: 'registerView',
            component: RegisterView,
        },
        {
            path: '/login',
            name: 'loginView',
            component:LoginView,
        }
    ],
})

export default router
