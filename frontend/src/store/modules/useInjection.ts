import { InjectionConfig } from '@/types'
import { defineStore } from 'pinia'

interface InjectionConfigState {
  injection: InjectionConfig
}

export const useInjection = defineStore('injection-config', {
  state: (): InjectionConfigState => {
    return {
      injection: {},
    }
  },
  actions: {
    clear() {
      this.$reset()
      localStorage.removeItem('injection-config')
    },
  },
  getters: {
    getInjection: (state) => {
      return state.injection
    },
  },
  persist: true,
})
