<template>
    <aside class="sidebar">
      <nav>
        <ul>
          <li>
            <router-link to="/chats" class="nav-link" active-class="active-link">
                <span class="material-icons">chat</span>
                <span>Чаты</span>
            </router-link>
          </li>
          <li>
            <router-link to="/settings" class="nav-link" active-class="active-link">
              Настройки
            </router-link>
          </li>
        </ul>
      </nav>
    </aside>
  </template>

<script lang="ts">
export default {
  name: 'Sidebar'
}
</script>

  <style scoped>
  .sidebar {
    width: 7%;
    height: 100vh;
    background-color: #f8f8f8;
    border-right: 1px solid #ddd;
  }

  nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }

  .nav-link {
    text-decoration: none;
    text-align: center;
    color: #333;
    padding: 20px 10px;
    display: block;
    font-size: small;
  }

  .nav-link:hover {
    background-color: #ddd;
  }

  .active-link {
    background-color: #007bff;
    color: white;
  }
  </style>
