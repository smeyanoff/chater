<template>
    <footer class="chat-input">
      <input v-model="message" type="text" placeholder="Напишите сообщение..." @keyup.enter="submitMessage" />
      <button @click="submitMessage">Отправить</button>
    </footer>
  </template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'MessageInput',
  emits: ['sendMessage'],
  setup (_, { emit }) {
    const message = ref('')

    const submitMessage = () => {
      if (message.value.trim()) {
        emit('sendMessage', message.value)
        message.value = '' // Очистить поле после отправки
      }
    }

    return {
      message,
      submitMessage
    }
  }
})
</script>

  <style scoped>
  /* Стили для ввода сообщения */
  </style>
