<template>
  <n-space vertical :size="24">
    <n-button type="error" @click="templateState.clear"> 重置 </n-button>
    <n-form label-align="left" label-placement="left" :label-width="140">
      <n-grid cols="1 s:2 m:3 l:4" :x-gap="24" responsive="screen">
        <n-form-item-gi label="禁用所有模板">
          <n-switch disabled />
        </n-form-item-gi>
        <n-form-item-gi label="禁用模板">
          <n-select disabled />
        </n-form-item-gi>
      </n-grid>
    </n-form>
  </n-space>
</template>
<script setup lang="ts">
import { useTemplate } from '@/store/modules/useTemplate'
import {
  NButton,
  NForm,
  NFormItemGi,
  NGrid,
  NSelect,
  NSpace,
  NSwitch,
} from 'naive-ui'

const templateState = useTemplate()
</script>

<style lang="scss" scoped></style>
