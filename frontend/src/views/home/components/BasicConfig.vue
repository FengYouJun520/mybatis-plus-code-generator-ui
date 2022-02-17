<template>
  <n-form
    label-align="left"
    label-placement="left"
    label-width="140"
    :model="basicState"
  >
    <n-grid responsive="screen" cols="1 s:2" :x-gap="24">
      <n-form-item-gi label="数据库类型：">
        <n-select
          v-model:value="basicState.dataSource.typ"
          :options="databaseTypeOptions"
        />
      </n-form-item-gi>
      <n-form-item-gi label="数据库名称：">
        <n-input v-model:value="basicState.dataSource.database" />
      </n-form-item-gi>
      <n-form-item-gi label="Host：">
        <n-input
          type="text"
          placeholder="请输入数据库Host"
          v-model:value="basicState.dataSource.host"
        />
      </n-form-item-gi>
      <n-form-item-gi label="Port：">
        <n-input-number
          type="text"
          :min="3306"
          :max="65535"
          :show-button="false"
          placeholder="请输入数据库"
          v-model:value="basicState.dataSource.port"
        />
      </n-form-item-gi>
    </n-grid>
    <n-form-item label="数据库账号：">
      <n-input
        type="text"
        placeholder="请输入数据库账号"
        v-model:value="basicState.dataSource.username"
      />
    </n-form-item>
    <n-form-item label="数据库密码：">
      <n-input
        type="password"
        placeholder="请输入数据库密码"
        show-password-on="click"
        v-model:value="basicState.dataSource.password"
      />
    </n-form-item>
    <n-form-item style="display: flex; justify-content: center">
      <n-space :size="24">
        <n-button type="error" @click="basicState.clear">重置</n-button>
        <n-button type="primary" @click="testConnection"> 测试连接 </n-button>
      </n-space>
    </n-form-item>
  </n-form>
</template>
<script setup lang="ts">
import { useApp } from '@/store/modules/useApp'
import { useBasic } from '@/store/modules/useBasic'
import { useStrategy } from '@/store/modules/useStrategy'
import {
  NButton,
  NForm,
  NFormItem,
  NFormItemGi,
  NGrid,
  NInput,
  NInputNumber,
  NSelect,
  NSpace,
  useMessage,
} from 'naive-ui'
import { ref } from 'vue'
import { DatabaseOptions } from '~wails/go/models'
import { databaseTypeOptions } from '../data'

const message = useMessage()
const basicState = useBasic()
const strategy = useStrategy()
const appState = useApp()

const testConnection = () => {
  window.go.main.App.PingDb(basicState.dataSource)
    .then((options: DatabaseOptions[] | Error) => {
      appState.setOptions(options as DatabaseOptions[])
      message.success('连接成功')
      appState.executeDisable = false
    })
    .catch((err: string) => {
      strategy.clearAddIncludes()
      appState.clearOptions()
      message.error(err)
    })
}
</script>

<style lang="scss" scoped></style>
