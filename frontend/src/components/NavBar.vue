<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
    <div class="container">
      <a class="navbar-brand" href="#">FashOJ</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarText"
        aria-controls="navbarText"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarText">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li
            class="nav-item"
            v-for="(navItem, index) in navItems"
            :key="index"
            @click="setActive(index)"
          >
            <a
              class="nav-link"
              :class="{ active: navItem.isActive }"
              :aria-current="navItem.isActive ? 'page' : undefined"
              href="#"
            >
              {{ navItem.text }}
            </a>
          </li>
        </ul>
        <span class="navbar-text">
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              zine yu
            </a>
            <ul class="dropdown-menu">
              <li><a class="dropdown-item" href="#">Action</a></li>
              <li><a class="dropdown-item" href="#">Another action</a></li>
              <li><a class="dropdown-item" href="#">Something else here</a></li>
            </ul>
          </li>
        </span>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import { ref } from 'vue'

const navItems = ref([
  { text: 'Home', isActive: true },
  { text: 'Problems', isActive: false },
  { text: 'Contest', isActive: false },
  {text:'Accouncement',isActive:false}
])

const setActive = (index: number) => {
  navItems.value.forEach((item) => (item.isActive = false))
  navItems.value[index].isActive = true
}
</script>

<style scoped>
.navbar {
  background-color: #fff !important;
  padding: 3px 0px;
}

.container {
  width: 80vw;
}

.dropdown {
  list-style: none;
}

.nav-item {
  margin: 0 25px;
  cursor: pointer;
}

.active {
  box-shadow: 0 3px #17a7cf;
}
</style>
