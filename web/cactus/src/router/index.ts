import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'dashboard',
            component: () => import('@/views/core/dashboard/DashboardPage.vue')
        },
        {
            path: '/processes',
            name: 'processes',
            component: () => import('@/views/core/type_worker/TypesWorkerPage.vue')
        },
        {
            path: '/tokens',
            name: 'tokens',
            component: () => import('@/views/core/tokens/TokensPage.vue')
        },
        {
            path: '/email',
            name: 'email',
            component: () => import('@/views/core/dashboard/DashboardPage.vue')
        },
        {
            path: '/telegram',
            name: 'telegram',
            component: () => import('@/views/core/dashboard/DashboardPage.vue')
        },
        {
            path: '/sms',
            name: 'sms',
            component: () => import('@/views/core/dashboard/DashboardPage.vue')
        },
        {
            path: '/push',
            name: 'push',
            component: () => import('@/views/core/dashboard/DashboardPage.vue')
        },
        {
            path: '/oneC',
            name: 'oneC',
            component: () => import('@/views/core/dashboard/DashboardPage.vue')
        },
        {
            path: '/users',
            name: 'users',
            component: () => import('@/views/core/users/UsersPage.vue')
        },
        {
            path: '/users/createnewuser',
            name: 'createnewuser',
            component: () => import('@/views/core/users/CreateNewUser.vue')
        },
        {
            path: '/edituser',
            name: 'edituser',
            component: () => import('@/views/core/users/EditUserPage.vue')
        },
        {
            path: '/deleteuser',
            name: 'deleteuser',
            component: () => import('@/views/core/users/DeleteUserPage.vue')
        },
    ]
})

export default router
