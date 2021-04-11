import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: () => import("../views/dashboard/index.vue")
  },
]

export default createRouter({
  history: createWebHistory(),
  routes
})
