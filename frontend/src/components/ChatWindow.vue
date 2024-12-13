<template>
  <section v-if="chat" class="chat-window">
    <!-- Заголовок чата -->
    <header class="chat-header" @click="openChatInfoModal">
      {{ chat.name }}
    </header>

    <!-- Контейнер для сообщений -->
    <div class="messages-container" ref="messagesContainer">
      <div
        v-for="(message, index) in messages || []"
        :key="message.id"
        :class="['message', { 'message-outgoing': message.isCurrent, 'message-incoming': !message.isCurrent }]"
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

    <!-- Поле ввода сообщения -->
    <footer class="chat-input">
      <textarea
        v-model="newMessage"
        placeholder="Написать сообщение..."
        @keydown="handleKeydown"
        @input="autoResize"
        ref="messageInput"
        class="input-field"
        rows="1"
      />
    </footer>

    <!-- Модальное окно -->
    <ChatInfoModal
      :isOpen="isChatInfoModalOpen"
      :chat="chat"
      @close="closeChatInfoModal"
    />
  </section>
</template>

<script lang="ts">
import { defineComponent, ref, PropType, watch, onBeforeUnmount, nextTick } from 'vue'
import { Chat, ChatMessage } from '@/types'
import { webSocketClient } from '@/api/websocket'
import ChatInfoModal from './ChatInfo/ChatInfoModal.vue'

export default defineComponent({
  name: 'ChatWindow',
  components: { ChatInfoModal },
  props: {
    chat: {
      type: Object as PropType<Chat | null>,
      required: false,
      default: null
    },
    messages: {
      type: Array as PropType<ChatMessage[] | null>,
      default: () => []
    }
  },
  emits: ['messageSent'],
  setup (props, { emit }) {
    const newMessage = ref<string>('')
    const isChatInfoModalOpen = ref(false) // Управление видимостью модального окна
    const messageInput = ref<HTMLTextAreaElement | null>(null)
    const messagesContainer = ref<HTMLElement | null>(null)

    // Открыть модальное окно
    const openChatInfoModal = () => {
      isChatInfoModalOpen.value = true
    }

    // Закрыть модальное окно
    const closeChatInfoModal = () => {
      isChatInfoModalOpen.value = false
    }

    const sendMessage = () => {
      try {
        if (!newMessage.value.trim()) {
          throw new Error('Message is empty')
        }
        if (!webSocketClient.isConnected()) {
          throw new Error('WebSocket not connected')
        }
        webSocketClient.send({ content: newMessage.value.trim() })
        newMessage.value = ''
      } catch (error) {
        if (error instanceof Error) {
          console.error('Send message error:', error.message)
        }
      }
    }

    const formatTime = (isoString: string) => {
      const date = new Date(isoString)
      const hours = String(date.getHours()).padStart(2, '0')
      const minutes = String(date.getMinutes()).padStart(2, '0')
      return `${hours}:${minutes}`
    }

    const handleKeydown = (event: KeyboardEvent) => {
      if (event.key === 'Enter' && !event.shiftKey) {
        sendMessage()
        event.preventDefault()
      }
    }

    const autoResize = () => {
      if (messageInput.value) {
        messageInput.value.style.height = 'auto'
        messageInput.value.style.height = `${messageInput.value.scrollHeight}px`
      }
    }

    onBeforeUnmount(
      async () => {
        if (webSocketClient.isConnected()) {
          webSocketClient.close()
        }
      }
    )

    // Автоматическая прокрутка к последнему сообщению
    const scrollToBottom = () => {
      nextTick(() => {
        if (messagesContainer.value) {
          messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
        }
      })
    }

    // --- Watchers ---
    watch(
      () => props.messages,
      () => {
        scrollToBottom()
      },
      { deep: true }
    )

    // Следим за сменой чатов и открываем новое WebSocket соединение
    watch(
      () => props.chat,
      async (newChat) => {
        if (!newChat) {
          console.log('No chat selected.')
          return
        }
        if (webSocketClient.isConnected()) {
          console.log('Closing previous WebSocket connection...')
          webSocketClient.close()
        }
        console.log('Switching to new chat:', newChat)
        // Открываем новое соединение
        try {
          console.log('Opening new WebSocket connection...')
          await webSocketClient.connect(`ws://localhost:54321/v1/chats/${newChat.id}/messages/ws`)

          if (webSocketClient.isConnected()) {
            // Подписываемся на получение сообщений
            webSocketClient.onMessage((message: unknown) => {
            // Проверяем, является ли message объектом и имеет ли нужные поля
              if (isChatMessage(message)) {
                const chatMessage = message as ChatMessage
                console.log('New message received:', chatMessage.content)
                emit('messageSent', chatMessage)
              } else {
                console.error('Invalid message format:', message)
              }
            })
          }
        } catch (error) {
          console.error('WebSocket connection error:', error)
        }
      },
      { immediate: true } // Выполнить при первом рендере
    )

    // Проверка, является ли объект ChatMessage
    function isChatMessage (message: unknown): message is ChatMessage {
      return (
        typeof message === 'object' &&
        message !== null &&
        'id' in message &&
        'content' in message &&
        'createdAt' in message &&
        'sender' in message &&
        'isCurrent' in message
      )
    }

    return {
      newMessage,
      isChatInfoModalOpen,
      messagesContainer,
      openChatInfoModal,
      closeChatInfoModal,
      sendMessage,
      formatTime,
      handleKeydown,
      autoResize,
      messageInput
    }
  }
})
</script>

<style scoped>
.chat-header:hover {
  color: #007aff;
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
  max-width: 400px;
  width: 90%;
}

.modal-content h3 {
  margin-bottom: 10px;
}

.modal-content p {
  margin: 5px 0;
}

.modal-content ul {
  margin: 0;
  padding: 0;
  list-style: none;
}

.modal-content ul li {
  margin: 5px 0;
}
.chat-window {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.chat-header {
  padding: 10px;
  border-bottom: 1px solid #ddd;
  font-size: large;
  font-weight: bold;
  cursor: pointer;
}

.messages-container {
  flex: 1;
  padding: 1px;
  overflow-y: auto;
  scroll-behavior: smooth;
  display: flex;
  flex-direction: column; /* Сообщения располагаются по вертикали */
}

/* Основной стиль для сообщения */
.message {
  display: flex;
  flex-direction: column;
  max-width: 70%;
  min-width: 10%;
  margin-bottom: 10px;
}

/* Входящие сообщения */
.message-incoming {
  align-self: flex-start; /* Выровнять контейнер слева */
  text-align: left;
  padding-left: 15px;
}

/* Исходящие сообщения */
.message-outgoing {
  align-self: flex-end; /* Выровнять контейнер справа */
  text-align: right;
  padding-right: 15px;
}

.message-content {
  display: inline-block;
  background-color: #e5e5ea;
  border-radius: 15px;
  padding: 5px 8px;
  word-break: break-word;
  white-space: pre-wrap;
  text-align: left; /* Текст всегда выравнен влево */
}

/* Входящие сообщения - изменяем цвет */
.message-incoming .message-content {
  background-color: #f0f0f0;
}

/* Исходящие сообщения - изменяем цвет */
.message-outgoing .message-content {
  background-color: #007aff;
  color: white;
}

/* Контейнер для имени и времени */
.message-meta {
  display: flex;
  margin-bottom: 5px;
  margin-left: 15px;
}

/* Имя отправителя */
.message-sender {
  font-weight: bold;
  color: #4CAF50;
  padding-right: 5px;
}

/* Время отправки */
.message-time {
  align-self:center;
  font-size: 0.8em;
  color: #999;
}

.chat-input {
  padding: 10px;
  border-top: 1px solid #ddd;
  display: flex;
}

.input-field {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
  resize: none; /* Отключаем изменение размера textarea */
  min-height: 24px; /* Устанавливаем минимальную высоту для одной строки */
  max-height: 240px;
  overflow: auto;
}

/* Добавляем автоматическое изменение высоты */
.input-field:focus {
  outline: none;
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
