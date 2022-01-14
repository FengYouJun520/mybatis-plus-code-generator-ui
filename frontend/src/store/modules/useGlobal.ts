import { GlobalConfig } from '@/types'
import { defineStore } from 'pinia'

interface GlobalConfigState {
  global: GlobalConfig
}

export const useGlobal = defineStore('global-config', {
  state: (): GlobalConfigState => {
    return {
      global: {
        fileOverride: false,
        disableOpenDir: true,
        outputDir: 'D://',
        author: 'FengYouJun',
        enableKotlin: false,
        enableSwagger: false,
        dateType: 'TIME_PACK',
        commentDate: 'yyyy-MM-dd',
      },
    }
  },
  actions: {
    clear() {
      this.$reset()
      localStorage.removeItem('global-config')
    },
  },
  getters: {
    getGlobal: (state) => {
      return state.global
    },
  },
  persist: true,
})
