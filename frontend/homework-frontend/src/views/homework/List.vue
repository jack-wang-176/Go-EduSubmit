<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { getHomeworkList } from '../../api/homework'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
const router = useRouter()
// 表格数据
const tableData = ref([])
const loading = ref(false)

// ✅ 1. 定义当前选中的部门
// 注意："Golang" 必须是你后端 model.Depart map 里真正存在的 Key！
// 如果你的 Map Key 是 "WEB" 或 "Java"，请改这里。
const currentDepartment = ref('Backend')

// ✅ 2. 修改选项列表
// value 必须严格匹配后端 Go 代码里的 Map Key (注意大小写！)
const departmentOptions = [
  { label: '后端 (Golang)', value: 'Backend' },   // 对应后端 "Backend"
  { label: '前端 (Web)', value: 'Frontend' },     // 对应后端 "Frontend"
  { label: 'Android', value: 'Android' },         // 对应后端 "Android"
  { label: 'iOS', value: 'IOS' },                 // 对应后端 "IOS" (注意全大写)
  { label: 'SRE (运维)', value: 'Sre' },          // 对应后端 "Sre"
  { label: '产品', value: 'Product' },            // 对应后端 "Product"
  { label: '设计', value: 'Design' }              // 对应后端 "Design"
]

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const fetchData = async () => {
  loading.value = true
  try {
    const res: any = await getHomeworkList({
      page: pagination.page,
      pageSize: pagination.pageSize,
      department: currentDepartment.value
    })

    tableData.value = res.data.list
    pagination.total = res.data.total

  } catch (error) {
    console.error("获取失败", error)
    ElMessage.error('查询失败：请检查部门名称是否正确')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
})

const handlePageChange = (newPage: number) => {
  pagination.page = newPage
  fetchData()
}


const handleDepartmentChange = () => {
  pagination.page = 1
  fetchData()
}
</script>

<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <div class="left-panel">
            <span>作业列表</span>

            <el-select
                v-model="currentDepartment"
                placeholder="选择部门"
                style="width: 150px; margin-left: 20px"
                @change="handleDepartmentChange"
            >
              <el-option
                  v-for="item in departmentOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
              />
            </el-select>
          </div>

          <el-button type="primary">发布作业</el-button>
        </div>
      </template>

      <el-table :data="tableData" style="width: 100%" v-loading="loading" border>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="title" label="作业标题" />
        <el-table-column prop="description" label="内容" show-overflow-tooltip />
        <el-table-column prop="deadline" label="截止时间" width="180" />

        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button
                link
                type="primary"
                size="small"
                @click="router.push(`/homework/${scope.row.id}`)"
            >
              详情
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
            background
            layout="prev, pager, next"
            :total="pagination.total"
            :page-size="pagination.pageSize"
            @current-change="handlePageChange"
        />
      </div>
    </el-card>
  </div>
</template>

<style scoped>


.page-container {
  padding: 20px;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.left-panel {
  display: flex;
  align-items: center;
}
.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>