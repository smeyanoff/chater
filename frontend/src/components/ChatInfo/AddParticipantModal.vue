<template>
    <div v-if="isOpen" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <span class="close-icon" @click="closeModal">✖</span>
        <h2>Добавить участника</h2>
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Введите имя пользователя"
          class="search-input"
        />
      </div>
    </div>
  </template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { User } from '@/types'

export default defineComponent({
  name: 'AddParticipantModal',
  props: {
    isOpen: {
      type: Boolean,
      required: true
    }
  },
  emits: ['close', 'addUser'],
  setup (props, { emit }) {
    const searchQuery = ref('')

    const closeModal = () => {
      emit('close')
    }

    const addUser = (user: User) => {
      emit('addUser', user)
    }

    return {
      searchQuery,
      closeModal,
      addUser
    }
  }
})
</script>

  <style scoped>
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
    z-index: 0;
  }

  .modal-content {
    background: white;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
    max-width: 250px;
    width: 90%;
    position: relative;
  }

  /* Крестик для закрытия */
  .close-icon {
    position: absolute;
    top: 10px;
    right: 10px;
    font-size: 1.5em;
    color: #999;
    cursor: pointer;
    transition: color 0.3s ease;
  }

  .search-input {
    width: 90%;
    padding: 5px;
    margin-left: 5px;
  }

  </style>
