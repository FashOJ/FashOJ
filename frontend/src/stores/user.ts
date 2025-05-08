import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', () => {
    const username = ref('')
    const token = ref('')
    const isLogin = ref(false)

    return { persist: true, username, token, isLogin }
})
