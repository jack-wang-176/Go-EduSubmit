// src/api/submission.ts
import request from '../utils/request'

interface SubmitData {
    homework_id: number
    content: string
}

// 提交作业
export const submitHomework = (data: SubmitData) => {
    return request({
        url: '/submission',
        method: 'post',
        data
    })
}

export const getHomeworkSubmissions = (homeworkId: number, params: any) => {
    return request({
        url: `/submission/homework/${homeworkId}`, // 对应后端 GET /submission/homework/:id
        method: 'get',
        params
    })
}

// 新增：批改作业的数据结构
interface ReviewData {
    id: number // 提交记录的 ID
    score: number
    comment: string
    is_excellent: boolean
}

// 新增：批改作业方法
export const reviewSubmission = (data: ReviewData) => {
    // 注意：后端路由是 PUT /submission/:id/review
    return request({
        url: `/submission/${data.id}/review`,
        method: 'put',
        data: {
            score: data.score,
            comment: data.comment,
            is_excellent: data.is_excellent
        }
    })
}
export const getExcellentSubmissions = (params: any) => {
    return request({
        url: '/submission/excellent', // 对应后端 GET /submission/excellent
        method: 'get',
        params
    })
}