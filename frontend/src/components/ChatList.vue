<template>
  <aside class="chat-list">
    <ul>
      <li
        v-for="chat in chats"
        :key="chat.id"
        :class="{ active: chat.id === selectedChatId }"
        @click="selectChat(chat)"
      >
        <div class="chat-item">
          <h3>{{ chat.name }}</h3>
          <div class="last-message-container">
            <span class="message-sender">{{ messageSender(chat.messages[chat.messages.length - 1]) }}:</span>
            <span>{{ chat.messages[chat.messages.length - 1].content }}</span> <!-- Это сообщение, которое будет обрезано -->
          </div>
        </div>
      </li>
    </ul>
  </aside>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Chat, ChatMessage } from '@/types'

export default defineComponent({
  name: 'ChatList',
  props: {
    chats: {
      type: Array as () => Chat[],
      required: true
    },
    selectedChatId: {
      type: Number,
      required: false
    }
  },
  emits: ['selectChat'],
  methods: {
    selectChat (chat: Chat) {
      this.$emit('selectChat', chat)
    }
  },
  setup () {
    const messageSender = (message: ChatMessage) => {
      if (message.isCurrent) {
        return 'Вы'
      } else {
        return message.sender
      }
    }
    return {
      messageSender
    }
  }
})
</script>

<style scoped>
.chat-list {
  width: 25%;
  height: 100vh;
  border-right: 1px solid #ddd;
  overflow-y: auto;
}
.chat-item {
padding: 10px;
border-radius: 4px;
transition: background-color 0.3s ease;  /* Плавный переход для цвета фона */
margin: 10px;
}
  /* Стили для активного чата */
.active .chat-item {
  background-color: #e0e0e0;  /* Цвет фона для выбранного чата */
  border-radius: 4px;  /* Закругляем углы */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);  /* Добавляем тень */
  margin: 5px;  /* Делаем небольшой отступ */
}

/* Наведение для всех чатов */
.chat-item:hover {
  background-color: #ddd;  /* Изменение цвета при наведении */
  cursor: pointer;  /* Показываем курсор как указатель */
}

.last-message-container {
  display: flex;
  align-items: center;
}

.message-sender {
  font-weight: bold;
  color: #4CAF50;
  padding-right: 5px;
}

.last-message-container span:last-child {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  flex-grow: 1; /* Занимает оставшееся пространство */
}
</style>
