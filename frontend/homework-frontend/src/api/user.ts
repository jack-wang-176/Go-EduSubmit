import request from '../utils/request'

// 登录接口
export const login = (data: any) => {
    return request({
        url: '/user/login', // 对应你 Go 后端的路由
        method: 'post',
        data // 发送用户名和密码
    })
}

interface RegisterData {
    username: string
    password: string
    nickname: string
    department: string // 枚举值: Backend, Frontend 等 (注意后端大小写敏感)
    role: string       // student (1) 或 admin (2)
}

// 新增：注册接口
export const register = (data: RegisterData) => {
    return request({
        url: '/user/register', // 对应后端 POST /user/register
        method: 'post',
        data
    })
}