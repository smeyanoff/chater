// Интерфейс для участника чата
export interface ChatMember {
  id: number;
  username: string;
}

// Интерфейс для сообщений
export interface ChatMessage {
  id: number;
  content: string;
  createdAt: string;
  sender: string;
  senderId: number;
}

// Интерфейс для чата
export interface Chat {
  id: number;
  name: string;
  createdAt: string;
  updatedAt: string;
  members: ChatMember[];
  messages: ChatMessage[];
}
