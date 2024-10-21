<template>
  <section v-if="chat" class="chat-window">
    <header class="chat-header">{{ chat.name }}</header>
    <div class="messages-container" ref="messagesContainer">
      <div
        v-for="(message, index) in messages"
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
  </section>
</template>

<script lang="ts">
import { defineComponent, ref, onBeforeUnmount, watch, nextTick, onMounted } from 'vue'
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
    }
  },
  emits: ['messageSent'],
  setup (props, { emit }) {
    const newMessage = ref<string>('')
    const messagesContainer = ref<HTMLElement | null>(null)
    const messageInput = ref<HTMLTextAreaElement | null>(null)

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

    const handleKeydown = (event: KeyboardEvent) => {
      if (event.key === 'Enter') {
        if (event.altKey) {
          // Если нажато Alt + Enter, вставляем перенос строки
          const cursorPosition = (event.target as HTMLTextAreaElement).selectionStart
          newMessage.value =
            newMessage.value.slice(0, cursorPosition) + '\n' + newMessage.value.slice(cursorPosition)
          autoResize()
          event.preventDefault() // Останавливаем стандартное поведение Enter
        } else {
          // Если просто Enter, то отправляем сообщение
          sendMessage()
          autoResize()
          event.preventDefault() // Останавливаем стандартное поведение Enter
        }
      }
    }

    // Функция для автоматического изменения высоты textarea
    const autoResize = () => {
      nextTick(() => {
        if (messageInput.value) {
          messageInput.value.style.height = 'auto' // Сброс высоты для вычисления
          messageInput.value.style.height = messageInput.value.scrollHeight + 'px'
        }
      })
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

    onMounted(() => {
      if (messageInput.value) {
        autoResize() // Сначала устанавливаем минимальную высоту
      }
    })

    // Следим за сменой чатов и открываем новое WebSocket соединение
    watch(
      () => props.chat,
      async (newChat) => {
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
      formatTime,
      autoResize,
      handleKeydown,
      messagesContainer,
      messageInput
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
  font-size: large;
  font-weight: bold;
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
