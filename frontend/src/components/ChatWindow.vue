<template>
    <section v-if="chat" class="chat-window">
      <header class="chat-header">{{ chat.name }}</header>
      <div class="messages-container">
        <!-- Проходим по сообщениям -->
        <div
          v-for="message in messages"
          :key="message.id"
          :class="{'message-outgoing': message.isCurrent, 'message-incoming': !message.isCurrent}"
        >
          <p>{{ message.content }}</p>
        </div>
      </div>
      <footer class="chat-input">
        <input
            v-model="newMessage"
            type="text"
            placeholder="Написать сообщение..."
            @keyup.enter="sendMessage"
            class="input-field"

        />
      </footer>
    </section>
    <section v-else class="chat-placeholder">
      <p>Выберите чат, чтобы начать переписку</p>
    </section>
  </template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { Chat, ChatMessage } from '@/types'
import { createMessage } from '@/api/chats'

export default defineComponent({
  name: 'ChatWindow',
  props: {
    chat: {
      type: Object as () => Chat | null,
      required: false
    },
    messages: {
      type: Array as () => ChatMessage[],
      required: true
    }
  },
  setup (props) {
    const newMessage = ref('')
    // Функция для отправки сообщения
    const sendMessage = async () => {
      const token = localStorage.getItem('authToken') // Получаем токен
      if (newMessage.value.trim() && props.chat && token) {
        try {
          // Отправляем сообщение на сервер
          const message = await createMessage(token, props.chat.id, newMessage.value)
          // Очищаем поле после отправки
          newMessage.value = ''
          // Добавляем новое сообщение в список сообщений
          props.messages.push(message)
          props.chat.messages = [message]
        } catch (error) {
          console.error('Ошибка при отправке сообщения:', error)
        }
      }
    }

    return {
      newMessage,
      sendMessage
    }
  }
})
</script>

  <style scoped>
  .chat-window {
    flex: 1;
    display: flex;
    flex-direction: column;
  }
  .chat-header {
    padding: 10px;
    border-bottom: 1px solid #ddd;
  }
  .messages-container {
    flex: 1;
    padding: 10px;
    overflow-y: auto;
  }
  .message-outgoing {
    text-align: right;
  }
  .message-incoming {
    text-align: left;
  }
  .chat-input {
    padding: 10px;
    border-top: 1px solid #ddd;
    display: flex;
  }
  .input-field {
    width: 100%;  /* Поле ввода занимает всю ширину контейнера */
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ddd;
    border-radius: 4px;
    box-sizing: border-box;  /* Для корректной работы с размерами */
    }
  .chat-placeholder {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 16px;
    color: #888;
  }
  </style>
