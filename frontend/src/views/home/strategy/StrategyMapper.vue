<template>
  <n-space vertical :size="24">
    <n-button type="error" @click="strategyState.clear"> 重置 </n-button>
    <n-form
      :label-width="200"
      label-align="left"
      label-placement="left"
      :model="strategyState.strategy.mapper"
    >
      <n-form-item label="开启@Mapper注解：">
        <n-switch
          v-model:value="strategyState.strategy.mapper.enableMapperAnnotation"
        />
      </n-form-item>
      <n-form-item label="设置通用父类：">
        <n-input v-model:value="strategyState.strategy.mapper.superClass" />
      </n-form-item>
      <n-form-item label="格式化mapper文件名称：">
        <n-input
          v-model:value="strategyState.strategy.mapper.formatMapperFileName"
        />
      </n-form-item>
      <n-form-item label="格式化xml实现类文件名称：">
        <n-input
          v-model:value="strategyState.strategy.mapper.formatXmlFileName"
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
