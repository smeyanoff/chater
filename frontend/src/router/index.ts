// src/router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import submitLogin from '@/components/Login.vue'
import submitRegister from '@/components/Registration.vue'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: submitLogin
  },
  {
    path: '/register',
    name: 'Register',
    component: submitRegister
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('authToken') // Проверяем наличие токена

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login') // Перенаправляем на страницу логина, если не авторизован
  } else {
    next() // Если аутентификация не требуется или пользователь авторизован
  }
})

export default router
