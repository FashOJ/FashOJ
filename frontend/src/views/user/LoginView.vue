<script setup lang="ts">
import NavBar from '@/components/NavBar.vue'
import axios from 'axios'

import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'
import { useToast } from 'vue-toastification'
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import router from '@/router'

const user = useUserStore()
const username = ref('')
const password = ref('')
const toast = useToast()

interface LoginResponse {
    message: string
    token: string
    // 其他可能的响应字段
}

interface ErrorResponse {
    message: string
    // 其他可能的错误字段
}

const login = () => {
    axios
        .post<LoginResponse>('http://localhost:3000/api/auth/login', {
            username: username.value,
            password: password.value,
        })
        .then((resp) => {
            const data = resp.data
            if (data.message === 'success') {
                user.username = username.value
                user.token = data.token
                toast.success('登录成功')
                router.push({ name: 'homeView' })
            } else {
                toast.error(data.message)
            }
        })
        .catch(error => {
            if (axios.isAxiosError<ErrorResponse>(error)) {
                const errorMessage = error.response?.data?.message || error.message
                toast.error(errorMessage, { timeout: 2000 })
            } else {
                toast.error('发生未知错误', { timeout: 2000 })
            }
        });
}
</script>

<template>
    <NavBar></NavBar>

    <div class="container">
        <div class="register-container">
            <h2 class="register-title">登录</h2>

            <form>
                <div class="mb-3">
                    <label for="username" class="form-label">用户名</label>
                    <input
                        type="text"
                        class="form-control"
                        id="username"
                        placeholder="输入用户名"
                        required
                        v-model="username"
                    />
                </div>

                <div class="mb-3">
                    <label for="password" class="form-label">密码</label>
                    <input
                        type="password"
                        class="form-control"
                        id="password"
                        placeholder="输入密码"
                        required
                        v-model="password"
                    />
                </div>


                <button type="button" class="btn btn-primary btn-register w-100" @click="login">
                    登录
                </button>
            </form>

            <div class="login-link">
                没有账号? <router-link :to="{ name: 'registerView' }">注册</router-link>
            </div>
        </div>
    </div>
</template>

<style scoped>
.register-container {
    width: 30vw;
    max-width: 500px;
    margin: 10vh auto;
    padding: 30px;
    background-color: white;
    border-radius: 10px;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
}
.register-title {
    text-align: center;
    margin-bottom: 30px;
    color: #333;
}
.form-control:focus {
    border-color: #86b7fe;
    box-shadow: 0 0 0 0.25rem rgba(13, 110, 253, 0.25);
}
.btn-register {
    background-color: #0d6efd;
    border: none;
    padding: 10px;
    font-weight: 600;
}
.btn-register:hover {
    background-color: #0b5ed7;
}
.login-link {
    text-align: center;
    margin-top: 20px;
}
</style>
