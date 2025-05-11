<script setup lang="ts">
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import { useUserStore } from '@/stores/user'
import { ref } from 'vue'

const userStore = useUserStore()

const navItems = ref([
    { text: 'Home', isActive: true, name: 'homeView' },
    // { text: 'Problems', isActive: false,name:"problemView"},
    // { text: 'Contest', isActive: false,name:"contestView" },
    { text: 'Announcement', isActive: false, name: 'announcementView' },
])

</script>

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
                    >
                        <router-link
                            class="nav-link"
                            :active_name="navItem.name"
                            :class="{ active: $route.name === navItem.name }"
                            :aria-current="$route.name === navItem.name ? 'page' : undefined"
                            :to="{ name: navItem.name }"
                        >
                            {{ navItem.text }}
                        </router-link>
                    </li>
                </ul>
                <span class="navbar-text" v-if="userStore.isLogin">
                    <li class="nav-item dropdown">
                        <a
                            class="nav-link dropdown-toggle"
                            href="#"
                            role="button"
                            data-bs-toggle="dropdown"
                            aria-expanded="false"
                        >
                            {{ userStore.username }}
                        </a>
                        <ul class="dropdown-menu">
                            <li><a class="dropdown-item" href="#">Action</a></li>
                            <li><a class="dropdown-item" href="#">Another action</a></li>
                            <li><a class="dropdown-item" href="#">Something else here</a></li>
                        </ul>
                    </li>
                </span>
                <span class="navbar-text" v-else>
                    <router-link class="nav-item auth" :to="{ name: 'loginView' }"
                        >Login</router-link
                    >
                    /
                    <router-link class="nav-item auth" :to="{ name: 'registerView' }"
                        >Register</router-link
                    >
                </span>
            </div>
        </div>
    </nav>
</template>

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
.auth {
    margin: 3px;
    text-decoration: none;
    color: #555555;
}
</style>
