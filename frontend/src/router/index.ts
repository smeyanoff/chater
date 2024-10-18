// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '@/views/AuthPage.vue'
import ChatView from '@/views/ChatView.vue'

const routes = [
  {
    path: '/', // Корневой маршрут
    redirect: '/auth' // Редирект на страницу логина
  },
  {
    path: '/auth',
    name: 'auth',
    component: AuthPage
  },
  {
    path: '/chats',
    name: 'chats', // Имя маршрута для списка чатов
    component: ChatView,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// router.beforeEach((to, from, next) => {
//   const isAuthenticated = !!Cookies.get('token') // Проверяем наличие токена

//   if (to.meta.requiresAuth && !isAuthenticated) {
//     next('/auth') // Перенаправляем на страницу логина, если не авторизован
//   } else {
//     next() // Если аутентификация не требуется или пользователь авторизован
//   }
// })

export default router
