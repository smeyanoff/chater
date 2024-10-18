<template>
  <section v-if="chat" class="chat-window">
    <header class="chat-header">{{ chat.name }}</header>
    <div class="messages-container" ref="messagesContainer">
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
import { defineComponent, ref, onBeforeUnmount, watch, nextTick } from 'vue'
import { Chat, ChatMessage } from '@/types'
import { webSocketClient } from '@/api/websocket'

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
    },
    socketClient: {
      type: Object,
      required: true
    }
  },
  emits: ['messageSent'],
  setup (props, { emit }) {
    const newMessage = ref('')
    const messagesContainer = ref<HTMLElement | null>(null)

    // Закрываем WebSocket перед удалением компонента
    onBeforeUnmount(() => {
      if (webSocketClient.isConnected()) {
        console.log('Closing WebSocket connection...')
        webSocketClient.close()
      }
    })

    // Отправка сообщения через WebSocket
    const sendMessage = () => {
      if (newMessage.value.trim() && webSocketClient.isConnected()) {
        console.log('Sending new message')
        webSocketClient.send({
          content: newMessage.value.trim()
        })
        newMessage.value = '' // Очищаем поле ввода
      } else {
        console.error('WebSocket не подключен или сообщение пустое')
      }
    }

    // Автоматическая прокрутка к последнему сообщению
    const scrollToBottom = () => {
      nextTick(() => {
        if (messagesContainer.value) {
          messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
        }
      })
    }

    // Формат времени для сообщений
    const formatTime = (isoString: string) => {
      const date = new Date(isoString)
      const hours = String(date.getHours()).padStart(2, '0')
      const minutes = String(date.getMinutes()).padStart(2, '0')
      return `${hours}:${minutes}`
    }

    // Следим за изменениями в сообщениях и скроллим вниз
    watch(
      () => props.messages.length,
      async () => {
        await nextTick()
        scrollToBottom()
      }
    )

    // Следим за сменой чатов и открываем новое WebSocket соединение
    watch(
      () => props.chat,
      async (newChat, oldChat) => {
        if (newChat) {
          console.log('Switching to new chat:', newChat.name)

          // Закрываем старое соединение, если оно существует
          if (webSocketClient.isConnected()) {
            console.log('Closing previous WebSocket connection...')
            webSocketClient.close()
          }

          // Открываем новое соединение
          try {
            console.log('Opening new WebSocket connection...')
            await webSocketClient.connect(`ws://localhost:54321/api/v1/chats/${newChat.id}/messages/ws`)

            if (webSocketClient.isConnected()) {
              // Подписываемся на получение сообщений
              webSocketClient.onMessage((message: ChatMessage) => {
                console.log('New message received:', message)
                emit('messageSent', message) // Отправляем сообщение родителю
              })
            }
          } catch (error) {
            console.error('WebSocket connection error:', error)
          }
        }
      },
      { immediate: true } // Выполнить при первом рендере
    )

    return {
      newMessage,
      sendMessage,
      formatTime,
      messagesContainer
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
  .input-field {    width: 100%;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ddd;
    border-radius: 4px;
    box-sizing: border-box;
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
