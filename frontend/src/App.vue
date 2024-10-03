<template>
  <div id="app">
    <!-- Здесь будет отображаться контент в зависимости от маршрута -->
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'App',
  setup () {
    // Определяем, нужно ли скрывать шапку (например, на страницах аутентификации)
    const isAuthenticated = ref(!!localStorage.getItem('authToken'))

    return {
      isAuthenticated // Используем в шаблоне для условного рендеринга
    }
  }
})
</script>

<style>
/* Пример простой стилизации */
nav {
  display: flex;
  gap: 10px;
  padding: 10px;
  background-color: #f5f5f5;
}

nav a {
  padding: 5px 10px;
  color: #6200ea;
  text-decoration: none;
  border-radius: 5px;
}

nav a.router-link-active {
  background-color: #6200ea;
  color: white;
}
</style>
