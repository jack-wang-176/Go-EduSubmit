<script setup lang="ts">
import { reactive, ref } from 'vue'
import { User, Lock } from '@element-plus/icons-vue'
import { login } from '../api/user'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)

// 表单数据
const loginForm = reactive({
  username: '',
  password: ''
})

// 登录逻辑
const handleLogin = async () => {
  if (!loginForm.username || !loginForm.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }

  loading.value = true
  try {
    const res: any = await login(loginForm)
    ElMessage.success('登录成功')

    // 保存 Token
    // 注意：这里要根据你后端实际返回的结构来取值
    localStorage.setItem('token', res.data.access_token)

    // 打印一下，确认存上了
    console.log("Token stored:", res.data.access_token)

    // ==========================================
    // 🚀 核心修改：登录成功后，跳转到首页
    // ==========================================
    // 这里的 '/' 会被路由重定向到 '/homework' (我们在 router/index.ts 里配过的)
    router.push('/')

  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <h2>Maple 系统登录</h2>
        </div>
      </template>

      <el-form :model="loginForm" size="large">
        <el-form-item>
          <el-input
              v-model="loginForm.username"
              placeholder="用户名"
              :prefix-icon="User"
          />
        </el-form-item>

        <el-form-item>
          <el-input
              v-model="loginForm.password"
              placeholder="密码"
              type="password"
              show-password
              :prefix-icon="Lock"
              @keyup.enter="handleLogin"
          />
        </el-form-item>

        <el-form-item>
          <el-button
              type="primary"
              :loading="loading"
              style="width: 100%"
              @click="handleLogin"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}
.login-card {
  width: 400px;
}
.card-header {
  text-align: center;
}
</style>