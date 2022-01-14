import { createPinia } from 'pinia'
import PiniaPluginPersist from 'pinia-plugin-persistedstate'
import { App } from 'vue'

const pinia = createPinia()
pinia.use(PiniaPluginPersist)

export function setupPinia(app: App) {
  app.use(pinia)
}

export default pinia
