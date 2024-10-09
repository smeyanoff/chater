<template>
    <div :class="{'message-outgoing': message.isOutgoing, 'message-incoming': !message.isOutgoing}">
      <p>{{ message.text }}</p>
    </div>
  </template>

<script lang="ts">
import { defineComponent } from 'vue'
import { ChatMessage } from '@/types'

export default defineComponent({
  name: 'MessageItem',
  props: {
    message: {
      type: Object as () => ChatMessage,
      required: true
    }
  }
})
</script>

  <style scoped>
  /* Стили для отдельного сообщения */
  .message-outgoing {
    text-align: right;
  }

  .message-incoming {
    text-align: left;
  }
  </style>
