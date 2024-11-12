<template>
  <aside class="chat-list">
    <ul>
      <!-- Вкладка "Личное" для чатов без группы -->
      <li>
        <h2 class="collapsible">
          <div class="group-title" @click="toggleGroup('personal')">
            <span :class="{ 'arrow-down': isPersonalOpen, 'arrow-right': !isPersonalOpen }"></span>
            Личное
          </div>
          <button @click.stop="openChatCreationModal(null)" class="create-chat-button">+</button>
        </h2>
        <ul v-if="isPersonalOpen">
          <li
            v-for="chat in personalChats"
            :key="chat.id"
            :class="{ active: chat.id === selectedChatId }"
            @click="selectChat(chat)"
          >
            <div class="chat-item">
              <h3>{{ chat.name }}</h3>
              <div v-if="chat.messages" class="last-message-container">
                <span class="message-sender">{{ messageSender(chat.messages[0]) }}:</span>
                <span>{{ chat.messages[0].content }}</span>
              </div>
            </div>
          </li>
        </ul>
      </li>

      <!-- Вкладки для групп -->
      <li v-for="group in groups" :key="group.id">
        <h2 class="collapsible">
          <div class="group-title" @click="toggleGroup(group.id)">
            <span :class="{ 'arrow-down': isGroupOpen(group.id), 'arrow-right': !isGroupOpen(group.id) }"></span>
            {{ group.name }}
          </div>
          <button @click.stop="openChatCreationModal(group.id)" class="create-chat-button">+</button>
        </h2>
        <ul v-if="isGroupOpen(group.id)">
          <li
            v-for="chat in groupChats(group.id)"
            :key="chat.id"
            :class="{ active: chat.id === selectedChatId }"
            @click="selectChat(chat)"
          >
            <div class="chat-item">
              <h3>{{ chat.name }}</h3>
              <div class="last-message-container">
                <span class="message-sender">{{ messageSender(chat.messages[chat.messages.length - 1]) }}:</span>
                <span>{{ chat.messages[chat.messages.length - 1].content }}</span>
              </div>
            </div>
          </li>
        </ul>
      </li>
    </ul>

    <!-- Модальное окно для создания чата -->
    <div v-if="isModalOpen" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <h3>Создать новый чат</h3>
        <input v-model="newChatName" placeholder="Введите название чата" />
        <button @click="createNewChat">Создать</button>
        <button @click="closeModal">Отмена</button>
      </div>
    </div>
  </aside>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { Chat, ChatMessage, Group } from '@/types'
import { createChat } from '@/api/chats'

export default defineComponent({
  name: 'ChatList',
  props: {
    chats: {
      type: Array as () => Chat[],
      required: true
    },
    groups: {
      type: Array as () => Group[],
      required: true
    },
    selectedChatId: {
      type: Number,
      required: false
    }
  },
  emits: ['selectChat', 'updateChats'],
  setup (props, { emit }) {
    const isPersonalOpen = ref(true)
    const openGroups = ref<Record<number, boolean>>({})
    const isModalOpen = ref(false)
    const newChatName = ref('')
    const selectedGroupId = ref<number | null>(null)

    const selectChat = (chat: Chat) => {
      emit('selectChat', chat)
    }

    const messageSender = (message: ChatMessage | undefined) => {
      if (!message) return ''
      return message.isCurrent ? 'Вы' : message.sender
    }

    const personalChats = computed(() => {
      return props.chats.filter(chat => chat.groups.length === 0)
    })

    const groupChats = (groupId: number) => {
      return props.chats.filter(chat => chat.groups.some(group => group.id === groupId))
    }

    const toggleGroup = (groupId: number | 'personal') => {
      if (groupId === 'personal') {
        isPersonalOpen.value = !isPersonalOpen.value
      } else {
        openGroups.value[groupId] = !openGroups.value[groupId]
      }
    }

    const isGroupOpen = (groupId: number) => {
      return !!openGroups.value[groupId]
    }

    const openChatCreationModal = (groupId: number | null) => {
      selectedGroupId.value = groupId
      newChatName.value = ''
      isModalOpen.value = true
    }

    const closeModal = () => {
      isModalOpen.value = false
      newChatName.value = ''
    }

    const createNewChat = async () => {
      if (!newChatName.value.trim()) {
        return
      }

      try {
        const newChat = await createChat(newChatName.value, selectedGroupId.value || undefined)
        emit('updateChats', newChat)
        closeModal()
      } catch (error) {
        console.error('Ошибка при создании чата:', error)
      }
    }

    return {
      selectChat,
      messageSender,
      personalChats,
      groupChats,
      toggleGroup,
      isGroupOpen,
      isPersonalOpen,
      isModalOpen,
      newChatName,
      openChatCreationModal,
      closeModal,
      createNewChat
    }
  }
})
</script>

<style scoped>
.chat-list h2 {
  font-size: 1em;
  margin: 5px 0;
  padding: 5px 10px;
  background-color: transparent;
  border-radius: 0;
  color: #666;
  font-weight: normal;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.group-title {
  display: flex;
  align-items: center;
  max-width: 120px; /* Устанавливаем фиксированную ширину для названия группы */
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis; /* Текст обрезается, если не влезает */
}

.create-chat-button {
  font-size: 1.2em;
  color: #007bff;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  width: 20px; /* Уменьшаем ширину кнопки */
  height: 20px; /* Уменьшаем высоту кнопки */
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
}

.create-chat-button:hover {
  color: #0056b3;
}

/* Общий стиль для контейнера чатов */
.chat-list {
  width: 18%;
  height: 100vh;
  border-right: 1px solid #ddd;
  overflow-y: auto;
  padding: 10px;
}

/* Сворачивающиеся и разворачивающиеся группы */
.collapsible {
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: space-between; /* Раздвигаем название и кнопку "+" */
}

/* Иконки для сворачивания/разворачивания */
.arrow-right::before,
.arrow-down::before {
  content: '▸';
  font-size: 0.8em; /* Уменьшаем размер иконки */
  margin-right: 5px;
  color: #666; /* Приглушенный цвет для иконки */
}

.arrow-down::before {
  content: '▾';
}

/* Стиль для элементов чатов */
.chat-item {
  padding: 15px;
  padding-left: 10px;
  border-radius: 4px;
  transition: background-color 0.3s ease;
}

.chat-item h3 {
  margin-top: 5px;
}

.active .chat-item {
  background-color: #e0e0e0;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.chat-item:hover {
  background-color: #ddd;
  cursor: pointer;
}

.last-message-container {
  display: flex;
  align-items: center;
}

.message-sender {
  font-weight: bold;
  color: #4CAF50;
  padding-right: 5px;
}

.last-message-container span:last-child {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  flex-grow: 1;
}

/* Стили для модального окна и затемненного фона */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 30px;
  border-radius: 8px;
  width: 90%;
  max-width: 400px;
  text-align: center;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.modal-content input {
  width: 100%;
  padding: 10px;
  margin-top: 10px;
  margin-bottom: 20px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.modal-content button {
  padding: 10px 15px;
  margin: 5px;
  cursor: pointer;
}

.modal-content button:first-of-type {
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
}

.modal-content button:last-of-type {
  background-color: #f0f0f0;
  color: #333;
  border: none;
  border-radius: 4px;
}
</style>
