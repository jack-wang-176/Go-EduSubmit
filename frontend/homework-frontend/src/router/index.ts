import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
// 引入 Layout
import Layout from '../layout/index.vue'

// 1. 定义路由表
const routes: Array<RouteRecordRaw> = [
    {
        path: '/login',
        name: 'Login',
        component: () => import('../views/Login.vue')
    },
    {
        path: '/',
        component: Layout, // 父路由使用 Layout (带导航栏的壳子)
        redirect: '/homework', // 访问根路径自动跳到作业列表
        children: [
            {
                path: 'homework', // 实际路径是 /homework
                name: 'HomeworkList',
                component: () => import('../views/homework/List.vue')
            },
            {

                path: 'homework/:id',
                name: 'HomeworkDetail',
                component: () => import('../views/homework/Detail.vue')
            },

            {

                path: 'homework/:id/submissions',
                name: 'HomeworkSubmissions',
                component: () => import('../views/homework/Submissions.vue')
            }
        ]
    }
]

// 2. 创建路由器
const router = createRouter({
    history: createWebHistory(),
    routes
})

// 3. 路由守卫
// 注意这里用了 _from (前面加了下划线)，这样 TS 就不会报错说“变量未使用了”
router.beforeEach((to, _from, next) => {
    const token = localStorage.getItem('token')

    if (to.path !== '/login' && !token) {
        next('/login')
    } else {
        next()
    }
})

export default router