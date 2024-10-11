<template>
  <section v-if="chat" class="chat-window">
    <header class="chat-header">{{ chat.name }}</header>
    <div class="messages-container" ref="messagesContainer"> <!-- ref на контейнер сообщений -->
      <!-- Проходим по сообщениям -->
      <div
        v-for="(message, index) in messages"
        :key="message.id"
        :class="{'message-outgoing': message.isCurrent, 'message-incoming': !message.isCurrent}"
      >
        <div class="message-meta">
          <span v-if="!message.isCurrent" class="message-sender">{{ message.sender }}</span>
          <span class="message-time">{{ formatTime(message.createdAt) }}</span>
        </div>
        <div class="message-content">
          <p>{{ message.content }}</p>
        </div>
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
import { defineComponent, ref, onMounted, watch, nextTick } from 'vue'
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
  emits: ['messageSent'],
  setup (props, { emit }) {
    const newMessage = ref('')
    const messagesContainer = ref<HTMLElement | null>(null) // ref на контейнер сообщений

    // Прокрутка к последнему сообщению
    const scrollToBottom = () => {
      nextTick(() => {
        if (messagesContainer.value) {
          messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight // Прокрутка контейнера до самого низа
        }
      })
    }

    // Форматирование времени для отображения
    const formatTime = (isoString: string) => {
      const date = new Date(isoString)
      const hours = String(date.getHours()).padStart(2, '0')
      const minutes = String(date.getMinutes()).padStart(2, '0')

      return `${hours}:${minutes}`
    }

    // Прокрутка при монтировании компонента
    onMounted(() => {
      scrollToBottom()
    })

    // Прокрутка при добавлении нового сообщения
    watch(
      () => props.messages.length, // Отслеживаем длину массива сообщений
      async () => {
        await nextTick() // Дожидаемся обновления DOM
        scrollToBottom() // Прокручиваем вниз после обновления
      }
    )

    // Отправка нового сообщения
    const sendMessage = async () => {
      const token = localStorage.getItem('authToken') // Получаем токен
      if (newMessage.value.trim() && props.chat && token) {
        try {
          // Отправляем сообщение на сервер
          const message = await createMessage(token, props.chat.id, newMessage.value)
          newMessage.value = '' // Очищаем поле после отправки
          emit('messageSent', message) // Отправляем новое сообщение в родительский компонент
        } catch (error) {
          console.error('Ошибка при отправке сообщения:', error)
        }
      }
    }

    return {
      newMessage,
      sendMessage,
      formatTime,
      messagesContainer // Возвращаем ссылку на контейнер сообщений
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
    padding: 1px;
    overflow-y: auto;
    scroll-behavior: smooth;
  }
  .message-outgoing {
    text-align: right;
    padding-right: 15px;
  }
  .message-incoming {
    text-align: left;
    padding-left: 15px;
  }

  .message-sender {
    font-weight: bold;
    color: #4CAF50;
    padding-right: 5px;
  }

  .message-time {
    font-size: 0.8em;
    color: #999;
  }

  .message-content {
    font-size: 1em;
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
