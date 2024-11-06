import { Chat, ChatMessage } from '@/types'
import apiClient from './axios'
import { ChatMessagesResponse, ChatsResponse } from './responses'

export const getAllUserChats = async (): Promise<Chat[]> => {
  const { data } = await apiClient.get<ChatsResponse>('/v1/chats',
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return data.chats
}

export const getChatMessages = async (chatId: number): Promise<ChatMessage[]> => {
  const { data } = await apiClient.get<ChatMessagesResponse>(`/v1/chats/${chatId}/messages`,
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return data.messages
}

export const createChat = async (chatName: string): Promise<Chat> => {
  const { data } = await apiClient.post<Chat>('/v1/chats', { chatName },
    {
      withCredentials: true // Включаем отправку cookie
    }
  )
  return data
}
