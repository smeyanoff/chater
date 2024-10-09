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
  const { data } = await apiClient.get<ChatResponse>('chats', {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data.chats
}

export const getChatMessages = async (token: string, chatId: number): Promise<ChatMessage[]> => {
  const { data } = await apiClient.get<ChatMessagesResponse>(`chats/${chatId}/messages`, {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data.messages
}

export const createChat = async (token: string, chatName: string): Promise<Chat> => {
  const { data } = await apiClient.post<Chat>('chats', { chatName }, {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data
}

export const createMessage = async (token: string, chatId: number, content: string): Promise<ChatMessage> => {
  const { data } = await apiClient.post<ChatMessage>(`chats/${chatId}/messages`, { content }, {
    headers: { Authorization: `Bearer ${token}` }
  })
  return data
}
