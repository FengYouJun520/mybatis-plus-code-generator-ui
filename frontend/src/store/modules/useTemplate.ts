import { TemplateConfig } from '@/types'
import { defineStore } from 'pinia'

interface TemplateConfigState {
  template: TemplateConfig
}

export const useTemplate = defineStore('template-config', {
  state: (): TemplateConfigState => {
    return {
      template: {},
    }
  },
  actions: {
    clear() {
      this.$reset()
      localStorage.removeItem('template-config')
    },
  },
  getters: {
    getTemplate: (state) => {
      return state.template
    },
  },
  persist: true,
})
