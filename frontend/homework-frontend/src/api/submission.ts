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