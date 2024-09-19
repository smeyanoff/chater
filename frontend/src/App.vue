<template>
  <div id="app">
    <nav>
      <router-link to="/login">Login</router-link>
      <router-link to="/register">Register</router-link>
    </nav>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  setup () {
    const isAuthenticated = ref(!!localStorage.getItem('authToken')) // Проверяем токен в localStorage

    return {
      isAuthenticated
    }
  }
})
</script>
