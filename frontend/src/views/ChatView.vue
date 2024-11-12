<template>
  <div class="messenger-view">
    <!-- Компонент списка чатов, передаем группы и чаты -->
    <ChatList
      :chats="chats"
      :groups="groups"
      :selectedChatId="selectedChat?.id"
      @selectChat="selectChat"
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
    const chatMessages = ref<ChatMessage[]>([])

    onMounted(async () => {
      chats.value = await getAllUserChats()
      groups.value = await getAllUserGroups() // Загрузка групп
    })

    const selectChat = async (chat: Chat) => {
      selectedChat.value = chat
      chatMessages.value = await getChatMessages(chat.id)
    }

    const addMessage = (message: ChatMessage) => {
      chatMessages.value.push(message)
      const chatIndex = chats.value.findIndex(chat => chat.id === selectedChat.value?.id)
      if (chatIndex !== -1) {
        chats.value[chatIndex].messages = [message]
      }
    }

    return {
      chats,
      groups,
      selectedChat,
      chatMessages,
      addMessage,
      selectChat
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
