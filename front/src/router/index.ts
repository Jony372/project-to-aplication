import { createRouter, createWebHistory } from "vue-router";
import TableComponent from "../components/TableComponent.vue";
import StatisticsComponent from "../components/StatisticsComponent.vue";

const routes = [
  {
    path: '/index',
    component: TableComponent
  },
  {
    path: '/statistics',
    component: StatisticsComponent
  },
  {
    path: '',
    redirect: '/index'
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/index'
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router;