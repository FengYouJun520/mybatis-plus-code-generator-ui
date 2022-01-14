import { StrategyConfig } from '@/types'
import { defineStore } from 'pinia'
import { useApp } from './useApp'

interface StrategyConfigState {
  strategy: StrategyConfig
}

export const useStrategy = defineStore('strategy-config', {
  state: (): StrategyConfigState => {
    return {
      strategy: {
        enableCapitalMode: false,
        enableSkipView: false,
        disableSqlFilter: true,
        addIncludes: [],
        enableSchema: false,
        entity: {
          disableSerialVersionUID: true,
          enableColumnConstant: false,
          enableChainModel: false,
          enableLombok: false,
          enableRemoveIsPrefix: false,
          enableTableFieldAnnotation: false,
          enableActiveRecord: false,
          naming: 'underline_to_camel',
          idType: 'AUTO',
          formatFileName: '%s',
        },
        controller: {
          enableRestStyle: false,
          enableHyphenStyle: false,
          formatFileName: '%sController',
        },
        service: {
          formatServiceFileName: '%sService',
          formatServiceImplFileName: 'I%sServiceImpl',
        },
        mapper: {
          enableMapperAnnotation: false,
          formatXmlFileName: '%sMapper',
          formatMapperFileName: '%sMapper',
        },
      },
    }
  },
  actions: {
    clear() {
      this.$reset()
      localStorage.removeItem('strategy-config')
    },
    setAddIncludes() {
      const appState = useApp()
      this.strategy.addIncludes = []
      appState.tablesOptions.forEach((option) => {
        this.strategy.addIncludes.push(option.value)
      })
    },
    clearAddIncludes() {
      this.strategy.addIncludes = []
    },
  },
  getters: {
    getStrategy: (state) => {
      return state.strategy
    },
  },
  persist: {
    paths: [
      'strategy.enableCapitalMode',
      'strategy.enableSkipView',
      'strategy.disableSqlFilter',
      'strategy.enableSchema',
      'strategy.entity',
      'strategy.controller',
      'strategy.service',
      'strategy.mapper',
    ],
  },
})
