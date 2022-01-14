import { DataSourceConfig } from '@/types'
import { defineStore } from 'pinia'

interface BasicState {
  dataSource: DataSourceConfig
}

export const useBasic = defineStore('basic-config', {
  state: (): BasicState => {
    return {
      dataSource: {
        typ: 'mysql',
        database: 'blog',
        username: 'root',
        password: 'root',
        host: 'localhost',
        port: 3306,
      },
    }
  },
  actions: {
    clear() {
      this.$reset()
      localStorage.removeItem('basic-config')
    },
  },
  getters: {
    getDataSource: (state) => {
      return state.dataSource
    },
  },
  persist: true,
})
