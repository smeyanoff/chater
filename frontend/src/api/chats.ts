import apiClient from './axios'
import { Chat, ChatMessage } from '@/types'

// Интерфейс для ответа от API
export interface ChatResponse {
  chats: Chat[];
}

export interface ChatMessagesResponse {
  messages: ChatMessage[];
}

export const getAllUserChats = async (): Promise<Chat[]> => {
  const { data } = await apiClient.get<ChatResponse>('api/v1/chats',
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return data.chats
}

export const getChatMessages = async (chatId: number): Promise<ChatMessage[]> => {
  const { data } = await apiClient.get<ChatMessagesResponse>(`api/v1/chats/${chatId}/messages`,
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return data.messages
}

export const createChat = async (chatName: string): Promise<Chat> => {
  const { data } = await apiClient.post<Chat>('api/v1/chats', { chatName },
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return data
}
