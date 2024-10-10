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
            <div v-for="message in chat.messages" :key="message.id">
                <p>{{ message.content }}</p>
            </div>
          </div>
        </li>
      </ul>
    </aside>
  </template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Chat } from '@/types'

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
  }
})
</script>

  <style scoped>
  .chat-list {
    width: 25%;
    height: 100vh;
    border-right: 1px solid #ddd;
    overflow-y: auto;
    background-color: #f7f7f7;
  }
  .chat-item {
    padding: 10px;
  }
  .active {
    background-color: #e6e6e6;
  }
  </style>
