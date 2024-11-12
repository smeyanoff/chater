// src/api/axios.ts
import axios from 'axios'
import router from '@/router'

const apiClient = axios.create({
  baseURL: 'http://localhost:54321', // Базовый URL твоего API
  headers: {
    'Content-Type': 'application/json'
  }
})

// Перехватчик ответов
apiClient.interceptors.response.use(
  response => response, // Пропускаем успешные ответы
  error => {
    if (error.response && error.response.status === 401) {
      // Если статус 401, значит, пользователь не авторизован
      console.error('Ошибка 401: Пользователь не авторизован')

      // Перенаправляем на страницу авторизации
      router.push('/auth')
    }
    return Promise.reject(new Error(error.error)) // Возвращаем ошибку для дальнейшей обработки
  }
)

export default apiClient
