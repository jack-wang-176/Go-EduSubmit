<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
// 引入我们刚才定义的 API
import { getHomeworkDetail } from '../../api/homework'
import { submitHomework } from '../../api/submission'
import { ElMessage } from 'element-plus'

const route = useRoute() // 用于获取当前 URL 里的参数 (比如 id)
const router = useRouter() // 用于跳转页面 (如果有需要)

// 定义接口类型，方便 TypeScript 提示
interface Submission {
  id: number
  score: number | null // 分数可能是 null
  is_excellent: boolean
}

interface HomeworkDetail {
  id: number
  title: string
  description: string
  deadline: string
  my_submission?: Submission // 这个字段可能为空（如果还没提交）
}

// 响应式数据
const detail = ref<HomeworkDetail | null>(null)
const content = ref('') // 用户输入的作业内容
const loading = ref(false)
const submitting = ref(false)

// 1. 获取详情逻辑
const fetchDetail = async () => {
  // 从 URL 路径 /homework/3 中拿到 "3"
  const id = Number(route.params.id)
  if (!id) return

  loading.value = true
  try {
    const res: any = await getHomeworkDetail(id)
    // 后端返回的数据赋值给 detail
    detail.value = res.data
  } catch (error) {
    console.error("获取详情失败", error)
  } finally {
    loading.value = false
  }
}

// 2. 提交作业逻辑
const handleSubmit = async () => {
  if (!content.value) {
    ElMessage.warning('请填写提交内容')
    return
  }
  if (!detail.value) return

  submitting.value = true
  try {
    await submitHomework({
      homework_id: detail.value.id,
      content: content.value
    })
    ElMessage.success('提交成功！')
    // 关键点：提交成功后，重新获取详情，这样界面会自动变成“已提交”状态
    content.value = '' // 清空输入框
    fetchDetail()
  } catch (error) {
    console.error("提交失败", error)
  } finally {
    submitting.value = false
  }
}

// 页面加载完成后，立刻获取数据
onMounted(() => {
  fetchDetail()
})
</script>

<template>
  <div class="detail-container" v-loading="loading">
    <el-card v-if="detail">
      <template #header>
        <div class="card-header">
          <h2>{{ detail.title }}</h2>
          <el-tag type="danger">截止时间: {{ detail.deadline }}</el-tag>
        </div>
      </template>

      <div class="description">
        <h3>作业要求：</h3>
        <p style="white-space: pre-wrap;">{{ detail.description }}</p>
      </div>

      <el-divider />

      <div class="submission-area">
        <h3>我的提交：</h3>

        <div v-if="detail.my_submission">
          <el-alert
              title="你已经提交过这份作业了"
              type="success"
              :description="detail.my_submission.score !== null ? `得分: ${detail.my_submission.score}` : '等待老师批改中...'"
              show-icon
              :closable="false"
          />
        </div>

        <div v-else>
          <el-input
              v-model="content"
              type="textarea"
              :rows="5"
              placeholder="请输入作业内容或 GitHub 链接..."
          />
          <div style="margin-top: 20px; text-align: right">
            <el-button type="primary" @click="handleSubmit" :loading="submitting">
              提交作业
            </el-button>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.detail-container {
  padding: 20px;
  max-width: 900px;
  margin: 0 auto;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>