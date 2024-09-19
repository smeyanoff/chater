<template>
    <form @submit.prevent="submitRegister">
        <div>
            <label for="login">Login</label>
            <input v-model="registerData.userName" id="login" type="login" required />
        </div>
        <div>
            <label for="email">Email:</label>
            <input v-model="registerData.email" id="email" type="email" required />
        </div>
        <div>
            <label for="password">Password:</label>
            <input v-model="registerData.password" id="password" type="password" required />
        </div>
        <div>
            <label for="confirmPassword">Confirm Password:</label>
            <input v-model="registerData.confirmPassword" id="confirmPassword" type="password" required />
        </div>
        <button type="submit" :disabled="loading">Register</button>
        <p v-if="error">{{ error }}</p>
    </form>
  </template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { register, RegisterRequest } from '@/api/auth'

export default defineComponent({
  setup () {
    const registerData = reactive<RegisterRequest>({
      userName: '',
      email: '',
      password: '',
      confirmPassword: ''
    })
    const loading = ref(false)
    const error = ref<string | null>(null)

    const submitRegister = async () => {
      if (registerData.password !== registerData.confirmPassword) {
        error.value = 'Passwords do not match!'
        return
      }

      loading.value = true
      error.value = null

      try {
        const response = await register(registerData)
        console.log('Registration successful:', response.data)
      } catch (err) {
        error.value = 'Registration failed. Please try again.'
        console.error('Registration error:', err)
      } finally {
        loading.value = false
      }
    }

    return {
      registerData,
      submitRegister,
      loading,
      error
    }
  }
})
</script>
