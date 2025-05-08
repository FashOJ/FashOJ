<script setup lang="ts">
import NavBar from '@/components/NavBar.vue'
import axios from 'axios'
import { onMounted, reactive } from 'vue'

interface Announcement {
    title: string | null
    abstract: string | null
}

const announcement: Announcement = reactive({ title: null, abstract: null })

const fetchLatestPost = () => {
    axios.get('http://localhost:3000/api/announcement/latest').then((res) => {
        if (res.data.message === 'success') {
            announcement.abstract = res.data.data.abstract
            announcement.title = res.data.data.title
        }
    })
}

onMounted(() => {
    fetchLatestPost()
})
</script>

<template>
    <NavBar></NavBar>

    <div class="announcement">
        <h1 class="title">{{ announcement.title }}</h1>
        <hr />
        <p class="text">{{ announcement.abstract }}</p>
    </div>
</template>

<style scoped>
hr {
    color: #a59d9d;
}

.announcement {
    margin: 0 auto;
    margin-top: 25px;
    width: 70vw;
    background-color: #fff;
    box-shadow: 0 0.375rem 1.375rem #afc2c980;
    padding: 15px;
    border-radius: 10px;
}

.title {
    margin-top: 15px;
    font-size: 25px;
}

.text {
    width: 100%;
    max-width: 100%;
    overflow-wrap: break-word;
    word-break: break-word;
    white-space: pre-wrap;
    hyphens: auto;
}
</style>
