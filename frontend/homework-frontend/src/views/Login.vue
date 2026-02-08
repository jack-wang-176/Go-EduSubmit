<script setup lang="ts">
import { reactive, ref } from 'vue'
// 引入图标
import { User, Lock, Postcard, House, check } from '@element-plus/icons-vue'
// ✅ 引入 register
import { login, register } from '../api/user'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

const router = useRouter()
const loading = ref(false)

// ✅ 核心变量：控制当前是登录还是注册模式
const isRegister = ref(false)

// 定义表单数据 (包含注册所需的所有字段)
const form = reactive({
  username: '',
  password: '',
  confirmPassword: '', // 前端校验用，不发给后端
  nickname: '',
  department: '',
  role: 'student' // 默认为 student (后端会自动转成枚举)
})

// 部门列表 (注意：后端是大小写敏感的，value 必须和后端 Map 的 Key 一致)
const departmentOptions = [
  { label: '后端 (Backend)', value: 'Backend' },
  { label: '前端 (Frontend)', value: 'Frontend' },
  { label: 'Android', value: 'Android' },
  { label: 'iOS', value: 'IOS' },
  { label: 'SRE (运维)', value: 'Sre' },
  { label: '产品', value: 'Product' },
  { label: '设计', value: 'Design' }
]

// 切换模式时清空表单，避免混淆
const toggleMode = () => {
  isRegister.value = !isRegister.value
  // 重置表单
  form.username = ''
  form.password = ''
  form.confirmPassword = ''
  form.nickname = ''
  form.department = ''
}

// 提交逻辑 (合并了登录和注册)
const handleSubmit = async () => {
  // 1. 基础校验 (登录注册都得填)
  if (!form.username || !form.password) {
    ElMessage.warning('用户名和密码必填')
    return
  }

  loading.value = true

  try {
    if (isRegister.value) {
      // === 注册模式 ===

      // 额外校验
      if (form.password !== form.confirmPassword) {
        ElMessage.error('两次输入的密码不一致')
        loading.value = false
        return
      }
      if (!form.nickname || !form.department) {
        ElMessage.warning('注册需要填写昵称和部门')
        loading.value = false
        return
      }

      // 调用注册接口
      await register({
        username: form.username,
        password: form.password,
        nickname: form.nickname,
        department: form.department,
        role: form.role
      })

      ElMessage.success('注册成功，请登录')
      // 注册成功后，自动切回登录模式
      isRegister.value = false

    } else {
      // === 登录模式 ===
      const res: any = await login({
        username: form.username,
        password: form.password
      })

      ElMessage.success('登录成功')

      // 存储 Token
      localStorage.setItem('token', res.data.access_token)
      localStorage.setItem('refresh_token', res.data.refresh_token)

      // 如果后端返回了 user 信息，也可以存一下角色
      if (res.data.user) {
        localStorage.setItem('role', res.data.user.role) // 1 或 2
        localStorage.setItem('department', res.data.user.department)
      }

      router.push('/')
    }
  } catch (error) {
    console.error("操作失败", error)
    // 错误处理通常由 axios 拦截器统一弹出 ElMessage
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
          <h2>{{ isRegister ? '注册新账号' : 'Maple 系统登录' }}</h2>
        </div>
      </template>

      <el-form :model="form" size="large">

        <el-form-item>
          <el-input
              v-model="form.username"
              placeholder="用户名 (账号)"
              :prefix-icon="User"
          />
        </el-form-item>

        <el-form-item v-if="isRegister">
          <el-input
              v-model="form.nickname"
              placeholder="你的昵称 (比如: 小登007)"
              :prefix-icon="Postcard"
          />
        </el-form-item>

        <el-form-item v-if="isRegister">
          <el-select
              v-model="form.department"
              placeholder="选择你的部门"
              style="width: 100%"
          >
            <template #prefix>
              <el-icon><House /></el-icon>
            </template>
            <el-option
                v-for="item in departmentOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-input
              v-model="form.password"
              placeholder="密码"
              type="password"
              show-password
              :prefix-icon="Lock"
          />
        </el-form-item>

        <el-form-item v-if="isRegister">
          <el-input
              v-model="form.confirmPassword"
              placeholder="确认密码"
              type="password"
              show-password
              :prefix-icon="Lock"
          />
        </el-form-item>

        <el-form-item>
          <el-button
              type="primary"
              :loading="loading"
              style="width: 100%"
              @click="handleSubmit"
          >
            {{ isRegister ? '立即注册' : '登录' }}
          </el-button>
        </el-form-item>

        <div class="toggle-link">
          <el-link type="primary" @click="toggleMode">
            {{ isRegister ? '已有账号？去登录' : '没有账号？去注册' }}
          </el-link>
        </div>

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
  /* 搞点渐变背景，看起来高级点 */
  background-image: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}
.login-card {
  width: 400px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}
.card-header {
  text-align: center;
}
.toggle-link {
  text-align: center;
  margin-top: -10px;
}
</style>