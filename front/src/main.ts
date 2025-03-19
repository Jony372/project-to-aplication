import { createApp, onMounted } from 'vue'
import './style.css'
import App from './App.vue'
import { initFlowbite } from 'flowbite'
import { createPinia } from 'pinia'
import router from './router'
import VueApexCharts from 'vue3-apexcharts'

onMounted(() => {
  initFlowbite()
})
createApp(App)
  .use(createPinia())
  .use(router)
  .use(VueApexCharts)
  .mount('#app')
