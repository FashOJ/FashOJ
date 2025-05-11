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
const confirmPassword = ref('')
const agree = ref(false)
const email = ref('')
const toast = useToast()

interface RegisterResponse {
    message: string
    token: string
    // 其他可能的响应字段
}

interface ErrorResponse {
    message: string
    // 其他可能的错误字段
}

const register = () => {
    if (password.value !== confirmPassword.value) {
        toast.error('两次输入的密码不匹配')
        return
    }
    if (!agree.value) {
        toast.error('请同意用户协议')
        return
    }

    axios
        .post<RegisterResponse>('http://localhost:3000/api/auth/register', {
            username: username.value,
            password: password.value,
            email: email.value,
        })
        .then((resp) => {
            const data = resp.data
            if (data.message === 'success') {
                user.username = username.value
                user.token = data.token
                user.isLogin = true
                toast.success('注册成功')
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
            <h2 class="register-title">创建新账户</h2>

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
                    <label for="email" class="form-label">电子邮箱</label>
                    <input
                        type="email"
                        class="form-control"
                        id="email"
                        placeholder="输入电子邮箱"
                        required
                        v-model="email"
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
                    <div class="form-text">密码长度至少8个字符</div>
                </div>

                <div class="mb-3">
                    <label for="confirmPassword" class="form-label">确认密码</label>
                    <input
                        type="password"
                        class="form-control"
                        id="confirmPassword"
                        placeholder="再次输入密码"
                        required
                        v-model="confirmPassword"
                    />
                </div>

                <div class="mb-3 form-check">
                    <input
                        type="checkbox"
                        class="form-check-input"
                        id="agreeTerms"
                        required
                        v-model="agree"
                    />
                    <label class="form-check-label" for="agreeTerms"
                        >我已阅读并同意<a href="#">服务条款</a></label
                    >
                </div>

                <button type="button" class="btn btn-primary btn-register w-100" @click="register">
                    注册
                </button>
            </form>

            <div class="login-link">
                已有账号? <router-link :to="{ name: 'loginView' }">登录</router-link>
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
