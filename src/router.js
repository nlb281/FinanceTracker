import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import RecordsPage from '@/views/RecordsPage.vue'

const routes = [
  { path: '/', name: 'home', component: HomeView },
  {
    path: '/transactions',
    name: 'transactions',
    component: RecordsPage,
    meta: { mode: 'transactions' },
  },
  { path: '/categories', name: 'categories', component: RecordsPage, meta: { mode: 'categories' } },
]

export default createRouter({
  history: createWebHistory(),
  routes,
})
