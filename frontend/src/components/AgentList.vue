<template>
    <form @submit.prevent="registerAgent">
      <div>
        <label for="name">Имя агента:</label>
        <input v-model="agent.name" required />
      </div>
      <div>
        <label for="description">Описание:</label>
        <textarea v-model="agent.description"></textarea>
      </div>
      <div>
        <label for="endpoint">URL Endpoint:</label>
        <input type="url" v-model="agent.endpoint" required />
      </div>
      <button type="submit">Добавить агента</button>
    </form>
  </template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup () {
    const agent = ref({
      name: '',
      description: '',
      endpoint: ''
    })

    const registerAgent = async () => {
      try {
        const response = await fetch('/api/agents', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(agent.value)
        })
        if (response.ok) {
          alert('Агент успешно добавлен!')
        } else {
          const error = await response.json()
          alert(`Ошибка: ${error.message}`)
        }
      } catch (err) {
        console.error('Ошибка подключения:', err)
      }
    }

    return { agent, registerAgent }
  }
})
</script>
