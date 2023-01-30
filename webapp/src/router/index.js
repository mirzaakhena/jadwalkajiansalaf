import {createRouter, createWebHistory} from 'vue-router'

const routes = [

    {
        path: '/',
        component: () => import('../pages/PageWithSidebar.vue'),
        children: [
            {
                path: '/admin',
                component: () => import('../pages/admin/ViewTable.vue'),
            },
        ],
    },

]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router