import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'dashboard',
            component: () => import('@/views/core/dashboard/DashboardPage.vue'),
        },
        {
            path: '/processes',
            name: 'processes',
            component: () => import('@/views/core/processes/ProcessesPage.vue'),
        },
        {
            path: '/tokens',
            name: 'tokens',
            component: () => import('@/views/core/tokens/TokensPage.vue'),
        }
    ]
})

export default router
