<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getExcellentSubmissions } from '../../api/submission'

const tableData = ref([])
const loading = ref(false)

const fetchData = async () => {
  loading.value = true
  try {
    // 默认拉取第一页，20条
    const res: any = await getExcellentSubmissions({ page: 1, pageSize: 20 })
    tableData.value = res.data.list
  } catch (error) {
    console.error("获取优秀作业失败", error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>

<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span style="font-weight: bold; color: #E6A23C; font-size: 18px;">
            🏆 优秀作业展示墙 (Hall of Fame)
          </span>
        </div>
      </template>

      <el-table :data="tableData" v-loading="loading" stripe border>
        <el-table-column label="大神昵称" width="150">
          <template #default="scope">
            <strong>{{ scope.row.student?.nickname || '神秘大神' }}</strong>
          </template>
        </el-table-column>

        <el-table-column label="作业题目" min-width="200">
          <template #default="scope">
            {{ scope.row.homework?.title }}
          </template>
        </el-table-column>

        <el-table-column label="得分" width="100">
          <template #default="scope">
            <span style="color: #F56C6C; font-weight: bold; font-size: 16px;">
              {{ scope.row.score }}
            </span>
          </template>
        </el-table-column>

        <el-table-column prop="comment" label="老师点评" show-overflow-tooltip />

        <el-table-column prop="created_at" label="入选时间" width="180" />
      </el-table>
    </el-card>
  </div>
</template>

<style scoped>
.page-container { padding: 20px; }
</style>