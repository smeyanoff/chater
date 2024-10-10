<template>
    <div class="auth-page-container">
        <div class="auth-page">
            <div class="tabs">
                <button
                  @click="activeTab = 'login'"
                  :class="{ active: activeTab === 'login' }" class="tabs-button-left">
                    Login
                </button>
                <button
                  @click="activeTab = 'register'"
                  :class="{ active: activeTab === 'register' }" class="tabs-button-right">
                    Register
                </button>
            </div>

            <!-- В зависимости от состояния activeTab показываем соответствующую форму -->
            <div v-if="activeTab === 'login'">
                <LoginForm />
            </div>
            <div v-else>
                <RegisterForm />
            </div>
        </div>
    </div>
  </template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import LoginForm from '@/components/LoginForm.vue'
import RegisterForm from '@/components/RegisterForm.vue'

export default defineComponent({
  name: 'AuthPage',
  components: {
    LoginForm,
    RegisterForm
  },
  setup () {
    // Переменная для отслеживания активной вкладки
    const activeTab = ref('login')

    return {
      activeTab
    }
  }
})
</script>

<style>
/* Центрирование формы */
    .auth-page-container {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh; /* Растягиваем контейнер на всю высоту экрана */
      width: 100vw;  /* Растягиваем контейнер на всю ширину экрана */
      background-color: #e5e0e0;
      padding: 20px;
      box-sizing: border-box;
      font-family: 'Roboto', sans-serif;
    }

    .auth-page {
      max-width: 400px;
      width: 100%;
      padding: 20px;
      background-color: #f5f5f5;
      border-radius: 10px;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    }

    .tabs {
      display: flex;
      justify-content: center;
      margin-bottom: 10px;
    }

    button {
      width: 150px;
      padding: 10px 10px;
      font-size: 20px;
    }

    .tabs-button-left {
      border-radius: 8px 0 0 8px; /* Скругляем верхние углы */
    }

    .tabs-button-right {
      border-radius: 0 8px 8px 0; /* Скругляем верхние углы */
    }

    .auth-form {
      display: block;
      box-sizing: border-box;
      background-color: white;
      padding: 30px;
      border-radius: 10px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
      text-align: left;
    }

    /* Метки */
    .form-group label {
      display: block;
      margin-top: 4px;
      margin-bottom: 4px;
      color: #555;
    }

    /* Поля ввода */
    .form-group input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    transition: border-color 0.3s ease;
    box-sizing: border-box;
    }

    .form-group input:focus {
    border-color: #3b30e5;
    outline: none;
    }

    /* Кнопка */
    .form-button {
    margin-top: 20px;
    margin-left: 100px;
    width: 40%;
    border-radius: 8px;
    }

    /* Сообщение об ошибке */
    .error-message {
    color: red;
    margin-top: 10px;
    font-size: 14px;
    text-align: center;
    }

</style>
