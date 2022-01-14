<template>
  <n-space vertical :size="24">
    <n-button type="error" @click="strategyState.clear"> 重置 </n-button>
    <n-form
      :label-width="200"
      label-align="left"
      label-placement="left"
      :model="strategyState.strategy.controller"
    >
      <n-form-item label="开启驼峰转连字符：">
        <n-switch
          v-model:value="strategyState.strategy.controller.enableHyphenStyle"
        />
      </n-form-item>
      <n-form-item label="开启@RestController 控制器：">
        <n-switch
          v-model:value="strategyState.strategy.controller.enableRestStyle"
        />
      </n-form-item>
      <n-form-item label="父类包名：" :label-width="140">
        <n-input v-model:value="strategyState.strategy.controller.superClass" />
      </n-form-item>
      <n-form-item label="格式化文件名称：" :label-width="140">
        <n-input
          v-model:value="strategyState.strategy.controller.formatFileName"
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
