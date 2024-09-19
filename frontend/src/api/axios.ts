// src/api/axios.ts
import axios from 'axios'

const apiClient = axios.create({
  baseURL: 'http://localhost:54321', // Базовый URL твоего API
  headers: {
    'Content-Type': 'application/json'
  }
})

export default apiClient
