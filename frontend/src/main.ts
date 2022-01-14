import 'vfonts/FiraCode.css'
import 'virtual:svg-icons-register'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { setupPinia } from './store'

const app = createApp(App)
setupPinia(app)
app.use(router)
app.mount('#app')
