<script setup lang="ts">
import { useRouter } from 'vue-router'
import { SwitchButton, User } from '@element-plus/icons-vue'

const router = useRouter()

const handleLogout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">
        <div class="logo">Maple 作业管理系统</div>
        <div class="user-info">
          <el-dropdown>
            <span class="el-dropdown-link">
              <el-icon><User /></el-icon> 我的账户
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="handleLogout" :icon="SwitchButton">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>
.header {
  background-color: #fff;
  border-bottom: 1px solid #dcdfe6;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  padding: 0 20px;
}
.logo {
  font-size: 20px;
  font-weight: bold;
  color: #409EFF;
}
.el-dropdown-link {
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 5px;
}
body {
  margin: 0;
  background-color: #f5f7fa;
}
</style>