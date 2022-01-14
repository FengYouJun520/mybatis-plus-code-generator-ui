<template>
  <n-space vertical :size="24" style="padding: 24px">
    <m-basic-config />
    <n-space vertical :size="24">
      <n-tabs default-value="global-config" type="card">
        <n-tab-pane name="global-config" tab="全局配置">
          <m-global-config />
        </n-tab-pane>
        <n-tab-pane name="package-config" tab="包配置">
          <m-package-config />
        </n-tab-pane>
        <n-tab-pane name="template-config" tab="模板配置">
          <m-template-config />
        </n-tab-pane>
        <n-tab-pane disabled name="injection-config" tab="注入配置">
        </n-tab-pane>
        <n-tab-pane name="strategy-config" tab="策略配置">
          <n-grid responsive="screen" :x-gap="24" cols="1 m:2">
            <n-grid-item>
              <n-button type="error" @click="strategyState.clear">
                重置
              </n-button>
              <n-form
                label-align="left"
                label-placement="left"
                :label-width="140"
                :model="strategyState"
              >
                <n-grid responsive="screen" :x-gap="24" cols="1 s:2">
                  <n-form-item-gi label="开启大写命名：">
                    <n-switch
                      disabled
                      v-model:value="strategyState.strategy.enableCapitalMode"
                    />
                  </n-form-item-gi>
                  <n-form-item-gi label="开启跳过视图：">
                    <n-switch
                      disabled
                      v-model:value="strategyState.strategy.enableSkipView"
                    />
                  </n-form-item-gi>
                  <n-form-item-gi label="禁用 sql 过滤：">
                    <n-switch
                      disabled
                      v-model:value="strategyState.strategy.disableSqlFilter"
                    />
                  </n-form-item-gi>
                  <n-form-item-gi label="启用 schema：">
                    <n-switch
                      disabled
                      v-model:value="strategyState.strategy.enableSchema"
                    />
                  </n-form-item-gi>
                </n-grid>
                <n-form-item>
                  <n-button
                    type="primary"
                    secondary
                    :disabled="appState.tablesOptions.length <= 0"
                    @click="selectAll"
                  >
                    全选
                  </n-button>
                </n-form-item>
                <n-form-item label="要生成的表列表：">
                  <n-select
                    clearable
                    multiple
                    v-model:value="strategyState.strategy.addIncludes"
                    :options="appState.tablesOptions"
                    :render-label="renderTablesOptions"
                    :render-tag="renderMultipleSelectTag"
                  />
                </n-form-item>
              </n-form>
            </n-grid-item>
            <n-grid-item>
              <n-tabs default-value="entity-config" size="small">
                <n-tab-pane name="entity-config" tab="Entity配置">
                  <strategy-entity />
                </n-tab-pane>
                <n-tab-pane name="controller配置-config" tab="Controller配置">
                  <strategy-controller />
                </n-tab-pane>
                <n-tab-pane name="service-config" tab="Service配置">
                  <strategy-service />
                </n-tab-pane>
                <n-tab-pane name="mapper-config" tab="Mapper配置">
                  <strategy-mapper />
                </n-tab-pane>
              </n-tabs>
            </n-grid-item>
          </n-grid>
        </n-tab-pane>
      </n-tabs>
      <n-button
        type="primary"
        block
        @click="execute"
        :disabled="appState.executeDisable"
        :loading="loadding"
      >
        执行
      </n-button>
    </n-space>
  </n-space>
</template>
<script setup lang="ts">
import { useApp } from '@/store/modules/useApp'
import { useBasic } from '@/store/modules/useBasic'
import { useGlobal } from '@/store/modules/useGlobal'
import { usePackage } from '@/store/modules/usePackage'
import { useStrategy } from '@/store/modules/useStrategy'
import { useTemplate } from '@/store/modules/useTemplate'
import {
  NButton,
  NForm,
  NFormItem,
  NFormItemGi,
  NGrid,
  NGridItem,
  NSelect,
  NSpace,
  NSwitch,
  NTabPane,
  NTabs,
  NTag,
  useMessage,
} from 'naive-ui'
import { h, ref } from 'vue'
import {
  ConfigContext,
  DataSourceConfig,
  GlobalConfig,
  PackageConfig,
  StrategyConfig,
  TemplateConfig,
} from '~wails/go/models'
import MBasicConfig from './components/BasicConfig.vue'
import MGlobalConfig from './components/GlobalConfig.vue'
import MPackageConfig from './components/PackageConfig.vue'
import MTemplateConfig from './components/TemplateConfig.vue'
import { TablesOptions } from './data'
import StrategyController from './strategy/StrategyController.vue'
import StrategyEntity from './strategy/StrategyEntity.vue'
import StrategyMapper from './strategy/StrategyMapper.vue'
import StrategyService from './strategy/StrategyService.vue'
const message = useMessage()

const loadding = ref(false)
const appState = useApp()
const basicState = useBasic()
const globalState = useGlobal()
const packageState = usePackage()
const templateState = useTemplate()
const strategyState = useStrategy()

const renderTablesOptions = (option: TablesOptions) => {
  return h(
    'div',
    {
      style: {
        display: 'flex',
        alignItems: 'center',
        columnGap: '32px',
        justifyContent: 'space-between',
      },
    },
    [h('div', null, option.label), h('div', option.comment)]
  )
}

const renderMultipleSelectTag = ({ option, handleClose }: any) => {
  return h(
    NTag,
    {
      closable: true,
      onClose: (e) => {
        e.stopPropagation()
        handleClose()
      },
    },
    h('span', null, option.label)
  )
}

const selectAll = () => {
  strategyState.setAddIncludes()
}

const execute = () => {
  if (
    !strategyState.strategy.addIncludes ||
    !strategyState.strategy.addIncludes?.length
  ) {
    message.warning('请在策略配置中选择要生成的表！')
    return
  }

  const configContext: ConfigContext = ConfigContext.createFrom({
    dataSource: DataSourceConfig.createFrom(basicState.dataSource),
    globalConfig: GlobalConfig.createFrom(globalState.global),
    packageConfig: PackageConfig.createFrom(packageState.package),
    templateConfig: TemplateConfig.createFrom(templateState.template),
    strategyConfig: StrategyConfig.createFrom(strategyState.strategy),
  })

  loadding.value = true

  // @ts-ignore
  window.go.codegen.Manager.CodeGenerate(configContext)
    .then((msg: string) => {
      message.success(msg)
    })
    .catch((err: string) => {
      message.error(err)
    })
    .finally(() => {
      loadding.value = false
    })
}
</script>

<style lang="scss" scoped></style>
