import { PackageConfig } from '@/types'
import { defineStore } from 'pinia'

interface PackageState {
  package: PackageConfig
}

export const usePackage = defineStore('package-config', {
  state: (): PackageState => {
    return {
      package: {
        parent: 'com.baomidou',
        entity: 'entity',
        service: 'service',
        serviceImpl: 'service.impl',
        mapper: 'mapper',
        mapperXml: 'mapper.xml',
        controller: 'controller',
        moduleName: '',
        other: '',
        pathInfo: [],
      },
    }
  },
  actions: {
    clear() {
      this.$reset()
      localStorage.removeItem('package-config')
    },
  },
  getters: {
    getPackage: (state) => {
      return state.package
    },
  },
  persist: true,
})
