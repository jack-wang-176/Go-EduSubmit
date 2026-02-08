<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { getHomeworkSubmissions, reviewSubmission } from '../../api/submission'
import { ElMessage } from 'element-plus'

const route = useRoute()
const homeworkId = Number(route.params.id) // 从 URL 获取作业 ID

// 表格数据
const tableData = ref([])
const loading = ref(false)

// 批改弹窗相关
const dialogVisible = ref(false)
const currentSub = reactive({
  id: 0,
  score: 0,
  comment: '',
  is_excellent: false
})

// 获取列表数据
const fetchData = async () => {
  loading.value = true
  try {
    const res: any = await getHomeworkSubmissions(homeworkId, {
      page: 1,
      page_size: 100 // 偷懒：一次性拉取100条，暂时不做分页
    })
    tableData.value = res.data.list
  } catch (error) {
    console.error("获取提交列表失败", error)
  } finally {
    loading.value = false
  }
}

// 打开批改弹窗 (点击“批改”按钮时触发)
const handleReview = (row: any) => {
  currentSub.id = row.id
  // 回显数据：如果之前批改过，就显示旧的分数；没批改过，给个默认值
  currentSub.score = row.score !== null ? row.score : 80
  currentSub.comment = row.comment || ''
  currentSub.is_excellent = row.is_excellent || false

  dialogVisible.value = true
}

// 提交批改结果
const submitReview = async () => {
  try {
    await reviewSubmission({
      id: currentSub.id,
      score: currentSub.score,
      comment: currentSub.comment,
      is_excellent: currentSub.is_excellent
    })
    ElMessage.success('批改完成！')
    dialogVisible.value = false // 关闭弹窗
    fetchData() // 🔄 刷新列表，显示最新分数
  } catch (error) {
    console.error("批改失败", error)
  }
}

// 页面加载时获取数据
onMounted(() => {
  fetchData()
})
</script>

<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>提交列表 (作业ID: {{ homeworkId }})</span>
        </div>
      </template>

      <el-table :data="tableData" v-loading="loading" border>
        <el-table-column prop="student.nickname" label="学生姓名" width="120" />
        <el-table-column prop="student.department_label" label="部门" width="100" />

        <el-table-column label="提交内容" min-width="200">
          <template #default="scope">
            <div class="content-text">{{ scope.row.content }}</div>
          </template>
        </el-table-column>

        <el-table-column prop="submitted_at" label="提交时间" width="180" />

        <el-table-column label="分数" width="120">
          <template #default="scope">
            <el-tag v-if="scope.row.score !== null" type="success" effect="dark">
              {{ scope.row.score }} 分
            </el-tag>
            <el-tag v-else type="info">未批改</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="优秀作业" width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.is_excellent" type="warning" effect="plain">Excellent</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="操作" width="100" fixed="right">
          <template #default="scope">
            <el-button type="primary" size="small" @click="handleReview(scope.row)">
              批改
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" title="批改作业" width="400px">
      <el-form label-position="top">
        <el-form-item label="给个分吧 (0-100)">
          <el-input-number v-model="currentSub.score" :min="0" :max="100" />
        </el-form-item>

        <el-form-item label="评语">
          <el-input
              v-model="currentSub.comment"
              type="textarea"
              :rows="3"
              placeholder="写点鼓励的话吧..."
          />
        </el-form-item>

        <el-form-item label="设为优秀作业">
          <el-switch
              v-model="currentSub.is_excellent"
              active-text="是"
              inactive-text="否"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitReview">提交结果</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.page-container {
  padding: 20px;
}
.content-text {
  white-space: pre-wrap; /* 保留换行 */
  max-height: 100px;
  overflow-y: auto; /* 内容太多出滚动条 */
}
</style>