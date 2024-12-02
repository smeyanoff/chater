<template>
  <div id="app">
    <Sidebar v-if="!isAuthPage" />
    <main :class="{ 'main-content-auth': isAuthPage, 'main-content': !isAuthPage }">
      <router-view />
    </main>
  </div>
</template>

<script lang="ts">
import Sidebar from '@/components/Sidebar.vue'
import { defineComponent, computed } from 'vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'App',
  components: {
    Sidebar
  },
  setup () {
    const route = useRoute()

    // Проверка: если маршрут '/auth', значит, мы на странице аутентификации
    const isAuthPage = computed(() => route.path === '/auth')

    return {
      isAuthPage
    }
  }
})
</script>

<style scoped>
#app {
  display: flex;
  height: 100vh;
}

.main-content {
  flex: 1;
  overflow-y: auto;
}

.main-content-auth {
  /* Стили для страницы аутентификации */
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
</style>
