import request from '../utils/request'

// 定义前端调用的参数结构
interface HomeworkQuery {
    page: number
    pageSize: number // 前端代码习惯用驼峰
    department: string // ✅ 必填：后端必须依靠这个字符串去查 Map
}

export const getHomeworkList = (params: HomeworkQuery) => {
    return request({
        url: '/homework',
        method: 'get',
        // 🚀 核心修改：在这里手动组装参数名，适配你的后端
        params: {
            page: params.page,
            page_size: params.pageSize, // 把前端的 pageSize 映射给后端的 page_size
            department: params.department // ✅ 把部门字符串传过去
        }
    })
}
export const getHomeworkDetail = (id: number) => {
    return request({
        url: `/homework/${id}`, // 对应后端 GET /homework/:id
        method: 'get'
    })
}
interface CreateHomeworkData {
    title: string
    description: string
    department: string // 后端需要枚举值，如 'Backend'
    deadline: string   // 格式 '2006-01-02 15:04:05'
    allow_late: boolean
}

// 新增：发布作业方法
export const createHomework = (data: CreateHomeworkData) => {
    return request({
        url: '/homework', // 对应后端 POST /homework
        method: 'post',
        data
    })
}