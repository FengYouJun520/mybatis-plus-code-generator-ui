<template>
  <n-space vertical :size="24">
    <n-button type="error" @click="globalState.clear">重置</n-button>
    <n-form
      :label-width="140"
      label-align="left"
      label-placement="left"
      :model="globalState"
    >
      <n-form-item label="作者名称：">
        <n-input v-model:value="globalState.global.author" />
      </n-form-item>
      <n-grid cols="1 s:2 m:3 l:4" :x-gap="24" responsive="screen">
        <n-form-item-gi label="覆盖已生成文件：">
          <n-switch v-model:value="globalState.global.fileOverride" />
        </n-form-item-gi>
        <n-form-item-gi label="禁止打开输出目录：">
          <n-switch v-model:value="globalState.global.disableOpenDir" />
        </n-form-item-gi>
        <n-form-item-gi label="开启 kotlin 模式：">
          <n-switch disabled />
        </n-form-item-gi>
        <n-form-item-gi label="开启 swagger 模式：">
          <n-switch v-model:value="globalState.global.enableSwagger" />
        </n-form-item-gi>
        <n-form-item-gi label="时间策略：">
          <n-select
            v-model:value="globalState.global.dateType"
            :options="dateTypeOptions"
          />
        </n-form-item-gi>
        <n-form-item-gi label="注释日期格式：">
          <n-input disabled v-model:value="globalState.global.commentDate" />
        </n-form-item-gi>
      </n-grid>
      <n-form-item label="指定输出目录：">
        <n-input v-model:value="globalState.global.outputDir" />
      </n-form-item>
    </n-form>
  </n-space>
</template>
<script setup lang="ts">
import { useGlobal } from '@/store/modules/useGlobal'
import {
  NButton,
  NForm,
  NFormItem,
  NFormItemGi,
  NGrid,
  NInput,
  NSelect,
  NSpace,
  NSwitch,
} from 'naive-ui'
import { dateTypeOptions } from '../data'

const globalState = useGlobal()
</script>

<style lang="scss" scoped></style>
