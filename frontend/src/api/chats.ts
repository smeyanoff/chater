import { Chat, ChatMessage } from '@/types'
import apiClient from './axios'
import { ChatMessagesResponse, ChatsResponse } from './responses'

export const getAllUserChats = async (): Promise<Chat[]> => {
  const { data } = await apiClient.get<ChatsResponse>('/v1/chats',
    {
      withCredentials: true
    }
  )

  return data.chats
}

export const getChatMessages = async (chatId: number): Promise<ChatMessage[]> => {
  const { data } = await apiClient.get<ChatMessagesResponse>(`/v1/chats/${chatId}/messages`,
    {
      withCredentials: true
    }
  )
  return data.messages
}

export const createChat = async (chatName: string, groupID?: number): Promise<Chat> => {
  const requestData: any = {
    name: chatName
  }

  // Если groupID указан, добавляем его в запрос
  if (groupID) {
    requestData.groupID = groupID
  }

  const { data } = await apiClient.post<Chat>('/v1/chats', requestData, {
    withCredentials: true
  })
  return data
}
