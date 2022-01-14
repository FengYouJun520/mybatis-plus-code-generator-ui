<template>
  <n-space vertical :size="24">
    <n-button type="error" @click="packageState.clear"> 重置 </n-button>
    <n-form
      :label-width="140"
      label-align="left"
      label-placement="left"
      :model="packageState"
    >
      <n-grid responsive="screen" cols="1 s:2 m:3" :x-gap="24">
        <n-form-item-gi label="父包名：">
          <n-input v-model:value="packageState.package.parent" />
        </n-form-item-gi>
        <n-form-item-gi label="父模块名：">
          <n-input v-model:value="packageState.package.moduleName" />
        </n-form-item-gi>
        <n-form-item-gi label="Entity 包名：">
          <n-input v-model:value="packageState.package.entity" />
        </n-form-item-gi>
        <n-form-item-gi label="Service 包名：">
          <n-input v-model:value="packageState.package.service" />
        </n-form-item-gi>
        <n-form-item-gi label="Service Impl 包名：">
          <n-input v-model:value="packageState.package.serviceImpl" />
        </n-form-item-gi>
        <n-form-item-gi label="Mapper 包名：">
          <n-input v-model:value="packageState.package.mapper" />
        </n-form-item-gi>
        <n-form-item-gi label="Mapper XML 包名：">
          <n-input v-model:value="packageState.package.mapperXml" />
        </n-form-item-gi>
        <n-form-item-gi label="Controller 包名：">
          <n-input v-model:value="packageState.package.controller" />
        </n-form-item-gi>
        <n-form-item-gi label="自定义文件包名 包名：">
          <n-input disabled />
        </n-form-item-gi>
      </n-grid>
    </n-form>
  </n-space>
</template>
<script setup lang="ts">
import { usePackage } from '@/store/modules/usePackage'
import { NButton, NForm, NFormItemGi, NGrid, NInput, NSpace } from 'naive-ui'

const packageState = usePackage()
</script>

<style lang="scss" scoped></style>
