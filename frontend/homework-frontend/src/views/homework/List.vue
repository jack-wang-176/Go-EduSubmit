<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
// ✅ 引入 ElMessageBox 和 deleteHomework
import { getHomeworkList, createHomework, deleteHomework } from '../../api/homework'
import { ElMessage, ElMessageBox } from 'element-plus'

const router = useRouter()

// 表格 loading 状态
const loading = ref(false)
const tableData = ref([])

// ✅ 新增：当前用户角色 (用于控制按钮显示)
const userRole = ref('')

// 默认选中部门
const currentDepartment = ref('Backend')

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

// 分页数据
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

// === 👇 发布作业逻辑 ===
const dialogVisible = ref(false)
const createLoading = ref(false)

const form = reactive({
  title: '',
  description: '',
  department: 'Backend',
  deadline: '',
  allow_late: false
})

const handleOpenDialog = () => {
  dialogVisible.value = true
}

const handleCreate = async () => {
  if (!form.title || !form.deadline) {
    ElMessage.warning('标题和截止时间必填')
    return
  }
  createLoading.value = true
  try {
    await createHomework({
      title: form.title,
      description: form.description,
      department: form.department,
      deadline: form.deadline,
      allow_late: form.allow_late
    })
    ElMessage.success('发布成功！')
    dialogVisible.value = false
    fetchData()
    // 重置
    form.title = ''
    form.description = ''
    form.deadline = ''
  } catch (error) {
    console.error("发布失败", error)
  } finally {
    createLoading.value = false
  }
}
// === 👆 发布逻辑结束 ===

// === 👇 新增：删除作业逻辑 ===
const handleDelete = (id: number) => {
  ElMessageBox.confirm(
      '确定要删除这个作业吗？删除后所有学生的提交记录也会一并消失，不可恢复！',
      '高危操作警告',
      {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning',
      }
  ).then(async () => {
    try {
      await deleteHomework(id)
      ElMessage.success('删除成功')
      fetchData() // 刷新列表
    } catch (error) {
      console.error(error)
    }
  }).catch(() => {
    // 取消删除
  })
}
// === 👆 删除逻辑结束 ===

// 获取数据方法
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
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  // ✅ 获取用户角色 (假设你在 Login.vue 里存的是 'role')
  // 这里的判断逻辑是：如果是 'admin' 或者是数字 '2' (取决于你后端返回啥)
  userRole.value = localStorage.getItem('role') || 'student'
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

          <el-button
              v-if="userRole === 'admin' || userRole === '2'"
              type="primary"
              @click="handleOpenDialog"
          >
            发布作业 (管理员)
          </el-button>
        </div>
      </template>

      <el-table :data="tableData" style="width: 100%" v-loading="loading" border>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="title" label="作业标题" />
        <el-table-column prop="description" label="内容" show-overflow-tooltip />
        <el-table-column prop="deadline" label="截止时间" width="180" />

        <el-table-column label="操作" width="220">
          <template #default="scope">
            <el-button
                link
                type="primary"
                size="small"
                @click="router.push(`/homework/${scope.row.id}`)"
            >
              详情
            </el-button>

            <el-button
                v-if="userRole === 'admin' || userRole === '2'"
                link
                type="warning"
                size="small"
                @click="router.push(`/homework/${scope.row.id}/submissions`)"
            >
              批改
            </el-button>

            <el-button
                v-if="userRole === 'admin' || userRole === '2'"
                link
                type="danger"
                size="small"
                @click="handleDelete(scope.row.id)"
            >
              删除
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

    <el-dialog
        v-model="dialogVisible"
        title="发布新作业"
        width="500px"
    >
      <el-form label-width="80px">
        <el-form-item label="标题">
          <el-input v-model="form.title" placeholder="请输入作业标题" />
        </el-form-item>

        <el-form-item label="所属部门">
          <el-select v-model="form.department" placeholder="请选择">
            <el-option
                v-for="item in departmentOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="截止时间">
          <el-date-picker
              v-model="form.deadline"
              type="datetime"
              placeholder="选择截止时间"
              value-format="YYYY-MM-DD HH:mm:ss"
              style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="允许补交">
          <el-switch v-model="form.allow_late" />
        </el-form-item>

        <el-form-item label="作业描述">
          <el-input
              v-model="form.description"
              type="textarea"
              :rows="4"
              placeholder="请输入作业的具体要求..."
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleCreate" :loading="createLoading">
            确认发布
          </el-button>
        </span>
      </template>
    </el-dialog>
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