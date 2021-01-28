import {createRouter, createWebHistory} from 'vue-router';

interface RoutesType {
    path: string;
    name: string;
    component: () => {};
    meta?: {
        index?: number;
        keepAlive?: boolean;
    };
    children?: RoutesType[];
}

const routes: RoutesType[] = [
    {
        path: '/',
        name: 'Home',
        component: () => import(/* webpackChunkName: "home" */ '../views/home/Home.vue'),
        children: [
            {
                path: '',
                name: 'Recommend',
                component: () => import(/* webpackChunkName: "recommend" */ '../views/home/Home.vue'),
                meta: {
                    index: 1,
                    keepAlive: false
                }
            },
            {
                path: 'attention/:id',
                name: 'Attention',
                component: () => import(/* webpackChunkName: "attention" */ '../views/home/Home.vue'),
                meta: {
                    index: 0,
                    keepAlive: false
                }
            },
            {
                path: 'hotList',
                name: 'HotList',
                component: () => import(/* webpackChunkName: "hostList" */ '../views/home/Home.vue'),
                meta: {
                    index: 2,
                    keepAlive: false
                }
            },
            {
                path: 'pneumonia',
                name: 'Pneumonia',
                component: () => import(/* webpackChunkName: "pneumonia" */ '../views/home/Home.vue'),
                meta: {
                    index: 4,
                    keepAlive: false
                }
            }
        ]
    },
    {
        path: '/vip',
        name: 'Vip',
        component: () => import(/* webpackChunkName: "vip" */ '../views/home/Home.vue'),
    },
    {
        path: '/self',
        name: 'Self',
        component: () => import(/* webpackChunkName: "self" */ '../views/home/Home.vue'),
    },
    {
        path: '/notify',
        name: 'Notify',
        component: () => import(/* webpackChunkName: "notify" */ '../views/home/Home.vue'),
    }
]

const router = createRouter({
    // createWebHistory 第一个参数为以前路由的base
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
