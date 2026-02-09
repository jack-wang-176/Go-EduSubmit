<script setup lang="ts">
import { ref, reactive } from 'vue'
// 1. 修复：删除了未使用的 'Message' 图标
import { User, Lock, Avatar, School } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { login, register } from '../api/user'
import { useRouter } from 'vue-router'

const router = useRouter()

// 控制是登录还是注册状态
const isRegister = ref(false)

// 表单数据
const form = reactive({
  username: '',
  password: '',
  confirmPassword: '',
  nickname: '',
  department: '',
  role: 'student'
})

// 部门选项
const departmentOptions = [
  { label: '后端 (Golang)', value: 'Backend' },
  { label: '前端 (Web)', value: 'Frontend' },
  { label: 'Android', value: 'Android' },
  { label: 'iOS', value: 'IOS' },
  { label: 'SRE (运维)', value: 'Sre' },
  { label: '产品', value: 'Product' },
  { label: '设计', value: 'Design' }
]

// 切换登录/注册模式
const toggleMode = () => {
  isRegister.value = !isRegister.value
  // 重置表单
  form.username = ''
  form.password = ''
  form.confirmPassword = ''
  form.nickname = ''
  form.department = ''
  form.role = 'student'
}

// 提交表单
const handleSubmit = async () => {
  // 基本校验
  if (!form.username || !form.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }

  if (isRegister.value) {
    // === 注册逻辑 ===
    if (form.password !== form.confirmPassword) {
      ElMessage.warning('两次输入的密码不一致')
      return
    }
    if (!form.nickname) {
      ElMessage.warning('请输入昵称')
      return
    }
    if (!form.department) {
      ElMessage.warning('请选择部门')
      return
    }

    try {
      await register({
        username: form.username,
        password: form.password,
        nickname: form.nickname,
        department: form.department,
        role: form.role
      })
      ElMessage.success('注册成功，请登录')
      toggleMode()
    } catch (error) {
      console.error('注册失败:', error)
    }

  } else {
    // === 登录逻辑 ===
    try {
      const res = await login({
        username: form.username,
        password: form.password
      })
      ElMessage.success('登录成功')

      localStorage.setItem('token', res.data.token)
      // 注意：有的后端返回结构可能是 res.data.role 而不是 res.data.user.role，请根据实际情况调整
      localStorage.setItem('role', res.data.user ? res.data.user.role : res.data.role)
      localStorage.setItem('nickname', res.data.user ? res.data.user.nickname : res.data.nickname)

      // 2. 修复：给 router.push 加上 await，消除警告
      await router.push('/')
    } catch (error) {
      console.error('登录失败:', error)
    }
  }
}
</script>

<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <span>{{ isRegister ? '注册新账号' : '登录 Maple 系统' }}</span>
        </div>
      </template>

      <el-form :model="form" size="large">
        <el-form-item>
          <el-input
              v-model="form.username"
              placeholder="用户名"
              :prefix-icon="User"
          />
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

        <template v-if="isRegister">
          <el-form-item>
            <el-input
                v-model="form.confirmPassword"
                placeholder="确认密码"
                type="password"
                show-password
                :prefix-icon="Lock"
            />
          </el-form-item>

          <el-form-item>
            <el-input
                v-model="form.nickname"
                placeholder="你的昵称 (比如: 小登007)"
                :prefix-icon="Avatar"
            />
          </el-form-item>

          <el-form-item>
            <el-select
                v-model="form.department"
                placeholder="选择你的部门"
                style="width: 100%"
            >
              <template #prefix>
                <el-icon><School /></el-icon>
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
            <el-radio-group v-model="form.role" style="width: 100%; display: flex; justify-content: space-around;">
              <el-radio label="student" size="large" border>👨‍🎓 学生</el-radio>
              <el-radio label="admin" size="large" border>👩‍🏫 管理员</el-radio>
            </el-radio-group>
          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" class="submit-btn" @click="handleSubmit">
            {{ isRegister ? '立即注册' : '登录' }}
          </el-button>
        </el-form-item>

        <div class="toggle-link">
          <el-button link type="primary" @click="toggleMode">
            {{ isRegister ? '已有账号? 去登录' : '注册新账号' }}
          </el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.login-container {
  height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  background-size: cover;
}

.login-card {
  width: 450px;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  max-width: 90%;
}

.card-header {
  text-align: center;
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  padding: 10px 0;
}

.submit-btn {
  width: 100%;
  font-size: 16px;
  padding: 20px 0;
}

.toggle-link {
  text-align: center;
  margin-top: -10px;
}

/* 3. 修复：使用 :deep() 穿透组件样式，解决选择器未使用的警告 */
:deep(.el-input__icon) {
  font-size: 18px;
}
</style>