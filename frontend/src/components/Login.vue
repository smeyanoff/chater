<template>
    <form @submit.prevent="submitLogin">
      <div>
        <label for="login">Login:</label>
        <input v-model="loginData.username" id="login" type="login" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input v-model="loginData.password" id="password" type="password" required />
      </div>
      <button type="submit" :disabled="loading">Login</button>
      <p v-if="error">{{ error }}</p>
    </form>
  </template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { login, LoginRequest } from '@/api/auth'

export default defineComponent({
  setup () {
    const loginData = reactive<LoginRequest>({
      username: '',
      password: ''
    })
    const loading = ref(false)
    const error = ref<string | null>(null)

    const submitLogin = async () => {
      loading.value = true
      error.value = null

      try {
        const response = await login(loginData)
        const token = response.data.token
        localStorage.setItem('authToken', token)

        console.log('Login successful:', response.data)
      } catch (err) {
        error.value = 'Login failed. Please try again.'
        console.error('Login error:', err)
      } finally {
        loading.value = false
      }
    }

    return {
      loginData,
      submitLogin,
      loading,
      error
    }
  }
})
</script>
