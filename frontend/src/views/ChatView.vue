<template>
    <div class="messenger-view">
      <ChatList :chats="chats" :selectedChatId="selectedChat?.id" @selectChat="selectChat" />
      <ChatWindow :chat="selectedChat" :messages="chatMessages" />
    </div>
  </template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { getAllUserChats, getChatMessages } from '@/api/chats'
import ChatList from '@/components/ChatList.vue'
import ChatWindow from '@/components/ChatWindow.vue'
import { Chat, ChatMessage } from '@/types'
import router from '@/router'

export default defineComponent({
  name: 'ChatView',
  components: { ChatList, ChatWindow },
  setup () {
    const chats = ref<Chat[]>([])
    const selectedChat = ref<Chat | null>(null)
    const chatMessages = ref<ChatMessage[]>([])

    onMounted(async () => {
      const token = localStorage.getItem('authToken')
      if (token) {
        chats.value = await getAllUserChats(token)
      } else {
        router.push('auth')
      }
    })

    // Обработчик выбора чата
    const selectChat = async (chat: Chat) => {
      selectedChat.value = chat // Устанавливаем выбранный чат
      const token = localStorage.getItem('authToken')
      if (token) {
        // Загружаем сообщения для выбранного чата
        chatMessages.value = await getChatMessages(token, chat.id)
      } else {
        router.push('auth')
      }
    }

    return {
      chats,
      selectedChat,
      chatMessages,
      selectChat
    }
  }
})
</script>

  <style scoped>
  .messenger-view {
    display: flex;
    height: 100vh;
  }
  </style>
