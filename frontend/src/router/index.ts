// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '@/views/AuthPage.vue'

const routes = [
  {
    path: '/auth',
    name: 'Auth',
    component: AuthPage
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('authToken') // Проверяем наличие токена

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/auth') // Перенаправляем на страницу логина, если не авторизован
  } else {
    next() // Если аутентификация не требуется или пользователь авторизован
  }
})

export default router
