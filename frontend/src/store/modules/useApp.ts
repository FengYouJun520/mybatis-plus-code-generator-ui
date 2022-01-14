import { TablesOptions } from '@/views/home/data'
import { defineStore } from 'pinia'
import { DatabaseOptions } from '~wails/go/models'

interface AppState {
  executeDisable: boolean
  tablesOptions: TablesOptions[]
}

export const useApp = defineStore('app', {
  state: (): AppState => {
    return {
      executeDisable: true,
      tablesOptions: [],
    }
  },
  actions: {
    setOptions(options: DatabaseOptions[]) {
      this.tablesOptions = []
      options.forEach((option) => {
        this.tablesOptions.push({
          label: option.name,
          value: option.name,
          comment: option.comment,
        })
      })
    },
    clearOptions() {
      this.tablesOptions = []
    },
  },
})
