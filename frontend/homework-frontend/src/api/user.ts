import request from '../utils/request'

// 登录接口
export const login = (data: any) => {
    return request({
        url: '/user/login', // 对应你 Go 后端的路由
        method: 'post',
        data // 发送用户名和密码
    })
}