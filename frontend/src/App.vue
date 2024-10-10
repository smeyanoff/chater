<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'App',
  setup () {
    const isAuthenticated = ref(!!localStorage.getItem('authToken'))

    return {
      isAuthenticated
    }
  }
})
</script>
