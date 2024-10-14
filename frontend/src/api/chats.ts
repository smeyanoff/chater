import apiClient from './axios'
import { Chat, ChatMessage } from '@/types'

// Интерфейс для ответа от API
export interface ChatResponse {
  chats: Chat[];
}

export interface ChatMessagesResponse {
  messages: ChatMessage[];
}

export const getAllUserChats = async (token: string): Promise<Chat[]> => {
  const { data } = await apiClient.get<ChatResponse>('api/v1/chats', {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data.chats
}

export const getChatMessages = async (token: string, chatId: number): Promise<ChatMessage[]> => {
  const { data } = await apiClient.get<ChatMessagesResponse>(`api/v1/chats/${chatId}/messages`, {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data.messages
}

export const createChat = async (token: string, chatName: string): Promise<Chat> => {
  const { data } = await apiClient.post<Chat>('api/v1/chats', { chatName }, {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data
}
