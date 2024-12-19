<template>
    <div v-if="isOpen" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <!-- Крестик для закрытия -->
        <span class="close-icon" @click="closeModal">✖</span>

        <!-- Имя группы и кнопка редактирования -->
        <div class="group-header">
          <p class="chat-name">{{ chat.name }}</p>
          <button class="edit-name-btn" @click="editGroupName" aria-label="Редактировать имя">
            <span class="material-icons">edit</span>
          </button>
        </div>

        <!-- Участники -->
        <div class="section">
          <button class="section-header" @click="openAddParticipantModal">
            <span>Участники</span>
            <span class="participants-count">{{ participantsCount }}</span>
            <span class="material-icons add-icon">person_add</span>
          </button>
          <ul class="list">
            <li
              v-for="member in chat.members"
              :key="member.id"
              class="list-item"
            >
              {{ member.username }}
            </li>
          </ul>
        </div>

        <!-- Группы -->
        <div class="section">
          <button class="section-header" @click="openAddGroupModal">
            <span>Группы</span>
            <span class="groups-count">{{ groupsCount }}</span>
            <span class="material-icons add-icon">group_add</span>
          </button>
          <ul class="list">
            <li
              v-for="group in chat.groups || []"
              :key="group.id"
              class="list-item"
            >
              {{ group.name }}
            </li>
          </ul>
        </div>
      </div>
    </div>

    <AddParticipantModal
      :isOpen="isAddParticipantOpen"
      @close="closeAddParticipantModal"
    />
    <AddGroupModal
      :isOpen="isAddGroupOpen"
      @close="closeAddGroupModal"
/>
  </template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { Chat } from '@/types'
import AddParticipantModal from './AddParticipantModal.vue'
import AddGroupModal from './AddGroupModal.vue'

export default defineComponent({
  name: 'ChatInfoModal',
  props: {
    isOpen: {
      type: Boolean,
      required: true
    },
    chat: {
      type: Object as () => Chat,
      required: true
    }
  },
  emits: ['close', 'addMember', 'addGroup', 'editChatName'],
  components: { AddParticipantModal, AddGroupModal },
  setup (props, { emit }) {
    const isAddParticipantOpen = ref(false)
    const isAddGroupOpen = ref(false)

    const openAddParticipantModal = () => {
      isAddParticipantOpen.value = true
    }

    const openAddGroupModal = () => {
      isAddGroupOpen.value = true
    }

    const closeAddParticipantModal = () => {
      isAddParticipantOpen.value = false
    }

    const closeAddGroupModal = () => {
      isAddGroupOpen.value = false
    }

    const closeModal = () => {
      emit('close')
    }

    const editGroupName = () => {
      emit('editChatName')
    }

    const participantsCount = computed(() => {
      const count = props.chat.members.length
      if (count === 1) return '1 участник'
      if (count >= 2 && count <= 4) return `${count} участника`
      return `${count} участников`
    })

    const groupsCount = computed(() => {
      const count = props.chat.groups?.length || 0
      if (count === 1) return '1 группа'
      if (count >= 2 && count <= 4) return `${count} группы`
      return `${count} групп`
    })

    return {
      closeModal,
      isAddParticipantOpen,
      isAddGroupOpen,
      openAddParticipantModal,
      openAddGroupModal,
      closeAddParticipantModal,
      closeAddGroupModal,
      editGroupName,
      participantsCount,
      groupsCount
    }
  }
})
</script>

  <style scoped>
  /* Подключение шрифта Google Material Icons */
  @import url('https://fonts.googleapis.com/icon?family=Material+Icons');

  /* Модальное окно */
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
    max-width: 500px;
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

  .close-icon:hover {
    color: #333;
  }

  /* Заголовок группы */
  .group-header {
    display: flex;
    justify-content: left;
    align-items: center;
  }

  .chat-name {
    font-size: 1.5em;
    font-weight: bold;
    color: #333;
    margin-right: 10px;
    margin-left: 10px;
  }

  .edit-name-btn {
    background: none;
    border: none;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
  }

  .edit-name-btn:hover {
    background: none;
  }

  .edit-name-btn .material-icons {
    font-size: 1em;
    color: #999;
    transition: color 0.3s ease;
  }

  .edit-name-btn:hover .material-icons {
    color: #007aff; /* Цвет иконки при наведении */
  }

  /* Заголовок секции */
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: white;
    padding: 8px 10px;
    border-radius: 6px;
    border: none;
    font-size: 1em;
    font-weight: bold;
    width: 100%;
    text-align: left;
    color: #333;
    cursor: pointer;
    margin-top: 10px;
    transition: background-color 0.3s ease, color 0.3s ease;
  }

  .section-header:hover {
    background: #f8f8f8;
    color: #007aff;
  }

  /* Иконка добавления */
  .add-icon {
    font-size: 1.2em;
    color: #999;
    transition: color 0.3s ease;
  }

  .section-header:hover .add-icon {
    color: #007aff;
  }

  /* Списки */
  .list {
    list-style: none;
    margin-top: 5px;
    padding: 10px;
    background-color: #f8f8f8;
    border-radius: 8px;
  }

  .list-item {
    padding: 10px;
    border-radius: 4px;
    transition: background-color 0.3s ease;
    cursor: pointer;
  }

  .list-item:hover {
    background-color: #eaeaea;
  }
  </style>
