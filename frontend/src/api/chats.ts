import apiClient from './axios';
import {Chat} from '@/types';

// Интерфейс для ответа от API
export interface ChatResponse {
    chats: Chat[];
}

export const responseAllUserChats = async (token: string): Promise<Chat[]> => {
    const { data } = await apiClient.get<ChatResponse>('/chats', {
        headers: {'Authorization': 'Bearer ${token}'}
    });
    return data.chats;
}
