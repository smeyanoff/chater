// src/api/auth.ts
import apiClient from './axios'
import { SuccessResponse } from './responses'

// Определяем типы данных для запросов и ответов
export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  userName: string;
  email: string;
  password: string;
  confirmPassword: string;
}

// Функция для аутентификации
export const login = async (data: LoginRequest): Promise<SuccessResponse> => {
  const response = await apiClient.post<SuccessResponse>('/v1/auth/login', data,
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return response.data
}

// Функция для регистрации
export const register = async (data: RegisterRequest): Promise<SuccessResponse> => {
  const response = await apiClient.post<SuccessResponse>('/v1/auth/register', data)
  return response.data
}
