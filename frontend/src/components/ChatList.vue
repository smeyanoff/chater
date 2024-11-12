<template>
  <aside class="chat-list">
    <ul>
      <!-- Вкладка "Личное" для чатов без группы -->
      <li>
        <h2 @click="toggleGroup('personal')" class="collapsible">
          <span :class="{ 'arrow-down': isPersonalOpen, 'arrow-right': !isPersonalOpen }"></span> Личное
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
              <div class="last-message-container">
                <span class="message-sender">{{ messageSender(chat.messages[chat.messages.length - 1]) }}:</span>
                <span>{{ chat.messages[chat.messages.length - 1].content }}</span>
              </div>
            </div>
          </li>
        </ul>
      </li>

      <!-- Вкладки для групп -->
      <li v-for="group in groups" :key="group.id">
        <h2 @click="toggleGroup(group.id)" class="collapsible">
          <span :class="{ 'arrow-down': isGroupOpen(group.id), 'arrow-right': !isGroupOpen(group.id) }"></span>
          {{ group.name }}
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
  </aside>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { Chat, ChatMessage, Group } from '@/types'

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
  emits: ['selectChat'],
  setup (props, { emit }) {
    const isPersonalOpen = ref(true)
    const openGroups = ref<Record<number, boolean>>({})

    const selectChat = (chat: Chat) => {
      emit('selectChat', chat)
    }

    const messageSender = (message: ChatMessage) => {
      return message.isCurrent ? 'Вы' : message.sender
    }

    // Чаты, не относящиеся к группам
    const personalChats = computed(() => {
      return props.chats.filter(chat => chat.groups.length === 0)
    })

    // Фильтр для чатов по ID группы
    const groupChats = (groupId: number) => {
      return props.chats.filter(chat => chat.groups.some(group => group.id === groupId))
    }

    // Переключение состояния группы по её ID
    const toggleGroup = (groupId: number | 'personal') => {
      if (groupId === 'personal') {
        isPersonalOpen.value = !isPersonalOpen.value
      } else {
        openGroups.value[groupId] = !openGroups.value[groupId]
      }
    }

    // Проверяем, открыта ли группа
    const isGroupOpen = (groupId: number) => {
      return !!openGroups.value[groupId]
    }

    return {
      selectChat,
      messageSender,
      personalChats,
      groupChats,
      toggleGroup,
      isGroupOpen,
      isPersonalOpen
    }
  }
})
</script>

<style scoped>
.chat-list {
  width: 18%;
  height: 100vh;
  border-right: 1px solid #ddd;
  overflow-y: auto;
  padding: 10px;
}

.collapsible {
  cursor: pointer;
  display: flex;
  align-items: center;
}

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

.chat-list h2 {
  font-size: 1.2em; /* Уменьшенный размер текста */
  margin: 5px 0;
  padding: 5px 10px;
  background-color: transparent; /* Убираем фон */
  border-radius: 0; /* Убираем скругление */
  text-align: left;
  color: #666; /* Меняем цвет на более приглушенный */
  font-weight: normal; /* Обычный вес шрифта */
}

.chat-list ul {
  padding-left: 0;
}

.chat-list li {
  list-style: none;
}
</style>
