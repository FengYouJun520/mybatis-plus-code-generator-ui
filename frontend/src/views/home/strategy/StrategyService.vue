<template>
  <n-space vertical :size="24">
    <n-button type="error" @click="strategyState.clear"> 重置 </n-button>
    <n-form
      :label-width="200"
      label-align="left"
      label-placement="left"
      :model="strategyState.strategy.service"
    >
      <n-form-item label="格式化service接口文件名称：">
        <n-input
          v-model:value="strategyState.strategy.service.formatServiceFileName"
        />
      </n-form-item>
      <n-form-item label="格式化service实现类文件名称：">
        <n-input
          v-model:value="
            strategyState.strategy.service.formatServiceImplFileName
          "
        />
      </n-form-item>
      <n-form-item label="service父类接口：">
        <n-input
          v-model:value="strategyState.strategy.service.superServiceClass"
        />
      </n-form-item>
      <n-form-item label="service实现类父类接口：">
        <n-input
          v-model:value="strategyState.strategy.service.superServiceImplClass"
        />
      </n-form-item>
    </n-form>
  </n-space>
</template>
<script setup lang="ts">
import { useStrategy } from '@/store/modules/useStrategy'
import { NButton, NForm, NFormItem, NInput, NSpace, NSwitch } from 'naive-ui'

const strategyState = useStrategy()
</script>

<style lang="scss" scoped></style>
