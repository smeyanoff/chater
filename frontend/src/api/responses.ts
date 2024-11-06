import { Chat, ChatMessage, Group } from '@/types'

// Интерфейс для ответа от API
export interface ChatsResponse {
  chats: Chat[];
}

export interface GroupsResponse {
    groups: Group[];
}

export interface ChatMessagesResponse {
  messages: ChatMessage[];
}

export interface SuccessResponse {
    message: string;
}
