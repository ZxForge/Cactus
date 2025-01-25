import { createRouter, createWebHistory } from 'vue-router';

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
            path: '/users',
            name: 'users',
            component: () => import('@/views/core/users/UsersPage.vue')
        },
        {
            path: '/users/create_user',
            name: 'create_user',
            component: () => import('@/views/core/users/UserCreate.vue')
        },
        {
            path: '/tokens',
            name: 'tokens',
            component: () => import('@/views/core/tokens/TokensPage.vue')
        },
        {
            path: '/tokens/create_token',
            name: 'create_token',
            component: () => import('@/views/core/tokens/TokenCreate.vue')
        },
    ]
});

export default router
