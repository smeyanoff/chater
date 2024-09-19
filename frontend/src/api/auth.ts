// src/api/auth.ts
import apiClient from './axios'

// Определяем типы данных для запросов и ответов
export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
}

export interface RegisterRequest {
  userName: string;
  email: string;
  password: string;
  confirmPassword: string;
}

export interface RegisterResponse {
  message: string;
}

// Функция для аутентификации
export const login = (data: LoginRequest) => {
  return apiClient.post<LoginResponse>('/auth/login', data)
}

// Функция для регистрации
export const register = (data: RegisterRequest) => {
  return apiClient.post<RegisterResponse>('/auth/register', data)
}
