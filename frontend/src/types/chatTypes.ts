import { Group } from './groupTypes'
import { User } from './userTypes'

// Интерфейс для сообщений
export interface ChatMessage {
  id: number;
  content: string;
  createdAt: string;
  sender: string;
  isCurrent: boolean;
}

// Интерфейс для чата
export interface Chat {
  id: number;
  name: string;
  members: User[];
  groups?: Group[];
}
