<script setup lang="ts">
import NavBar from '@/components/NavBar.vue'
</script>

<template>
    <NavBar></NavBar>
    <div class="container">
        <div class="announcement">
            <div class="head">
                <h1 class="title">testTitle</h1>
                <div class="user">
                    <img
                        src="https://avatars.githubusercontent.com/u/46991452?v=4"
                        alt=""
                        class="avatar"
                    />
                    <p class="username">zine</p>
                </div>
            </div>
            <p class="text">testContent</p>
        </div>
        <hr />
        <div class="announcement">
            <div class="head">
                <h1 class="title">testTitle</h1>
                <div class="user">
                    <img
                        src="https://avatars.githubusercontent.com/u/46991452?v=4"
                        alt=""
                        class="avatar"
                    />
                    <p class="username">zine</p>
                </div>
            </div>

            <p class="text">testContent</p>
        </div>
    </div>
</template>

<style scoped>
hr {
    color: #a59d9d;
}

.container {
    background-color: #fff;
    box-shadow: 0 0.375rem 1.375rem #afc2c980;
    width: 70vw;
    border-radius: 10px;
    margin-top: 25px;
    padding-top: 15px;
}

.announcement {
    margin: 0 auto;
    width: 70vw;
    padding: 15px;
}

.head {
    display: flex;
    align-items: center;
}

.user {
    display: flex;
    align-self: center;
    margin-right: 15px;
    margin-left: auto;
}

.avatar {
    width: 20px;
    margin-right: 12px;
}
.username {
    margin: 0;
}

.title {
    margin:0;
    font-size: 25px;
}

.text {
    width: 100%;
    max-width: 100%;
    overflow-wrap: break-word;
    word-break: break-word;
    white-space: pre-wrap;
    hyphens: auto;
    margin-top: 15px;
}
</style>
