<template>
  <div class="messenger-view">
    <!-- Компонент списка чатов, передаем группы, чаты и последние сообщения -->
    <ChatList
      :chats="chats"
      :groups="groups"
      :lastMessages="lastMessages"
      :selectedChatId="selectedChat?.id"
      @selectChat="selectChat"
      @updateChats="updateChats"
    />
    <!-- Окно чата для выбранного чата -->
    <ChatWindow
      v-if="selectedChat"
      :chat="selectedChat"
      :messages="chatMessages"
      @messageSent="addMessage"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { getAllUserChats, getChatMessages } from '@/api/chats'
import { getAllUserGroups } from '@/api/groups'
import ChatList from '@/components/ChatList.vue'
import ChatWindow from '@/components/ChatWindow.vue'
import { Chat, ChatMessage, Group } from '@/types'

export default defineComponent({
  name: 'ChatView',
  components: { ChatList, ChatWindow },
  setup () {
    const chats = ref<Chat[]>([])
    const groups = ref<Group[]>([])
    const selectedChat = ref<Chat | null>(null)
    const chatMessages = ref<ChatMessage[]>([]) // Инициализация пустым массивом
    const lastMessages = ref<Record<number, ChatMessage | null>>({}) // Объект для последних сообщений

    onMounted(async () => {
      try {
        chats.value = await getAllUserChats() || []
        groups.value = await getAllUserGroups() || [] // Загрузка групп с обработкой null

        // Инициализация lastMessages для каждого чата
        for (const chat of chats.value) {
          const messages = await getChatMessages(chat.id)
          lastMessages.value[chat.id] = messages.length ? messages[messages.length - 1] : null
        }
      } catch (error) {
        console.error('Ошибка при загрузке чатов, групп или сообщений:', error)
      }
    })

    // Обработчик выбора чата
    const selectChat = async (chat: Chat) => {
      selectedChat.value = chat
      try {
        chatMessages.value = await getChatMessages(chat.id) || []
      } catch (error) {
        console.error('Ошибка при загрузке сообщений:', error)
        chatMessages.value = [] // Устанавливаем пустой массив в случае ошибки
      }
    }

    const updateChats = (chat: Chat) => {
      chats.value.push(chat)
    }

    // Метод для добавления сообщения
    const addMessage = (message: ChatMessage) => {
      chatMessages.value.push(message)

      // Обновляем последнее сообщение для выбранного чата в объекте lastMessages
      if (selectedChat.value) {
        lastMessages.value[selectedChat.value.id] = message
      }
    }

    return {
      chats,
      groups,
      selectedChat,
      chatMessages,
      lastMessages,
      addMessage,
      selectChat,
      updateChats
    }
  }
})
</script>

<style scoped>
.messenger-view {
  display: flex;
  height: 100%;
  overflow: hidden;
}
</style>
