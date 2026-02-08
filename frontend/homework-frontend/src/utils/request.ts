import axios from 'axios'
import { ElMessage } from 'element-plus'

// 1. 创建 axios 实例
const service = axios.create({
    // 这里的 '/api' 对应我们在 vite.config.ts 里配置的代理
    baseURL: '/api',
    timeout: 5000 // 请求超时时间
})

// 2. 请求拦截器 (类似 Gin 的 Middleware)
service.interceptors.request.use(
    (config) => {
        // 如果本地有 token，就自动放到 Header 里
        const token = localStorage.getItem('token')
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`
        }
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

// 3. 响应拦截器
service.interceptors.response.use(
    (response) => {
        const res = response.data
        // 假设你的后端约定：code === 0 为成功 (如果你的后端是用 code 200 表示成功，请在这里改成 200)
        if (res.code !== 0) {
            ElMessage.error(res.message || '系统错误')
            return Promise.reject(new Error(res.message || 'Error'))
        } else {
            return res
        }
    },
    (error) => {
        ElMessage.error(error.message || '网络请求失败')
        return Promise.reject(error)
    }
)

export default service