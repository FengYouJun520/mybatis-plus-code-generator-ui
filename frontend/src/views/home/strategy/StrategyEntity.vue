<template>
  <n-space vertical :size="24">
    <n-button type="error" @click=""> 重置 </n-button>
    <n-form
      :label-width="160"
      label-align="left"
      label-placement="left"
      :model="strategyState.strategy.entity"
    >
      <n-grid responsive="screen" cols="1 s:2" :x-gap="24">
        <n-form-item-gi label="禁用SerialVersionUID：">
          <n-switch
            v-model:value="
              strategyState.strategy.entity.disableSerialVersionUID
            "
          />
        </n-form-item-gi>
        <n-form-item-gi label="使用链式模式：">
          <n-switch
            v-model:value="strategyState.strategy.entity.enableChainModel"
          />
        </n-form-item-gi>
        <n-form-item-gi label="使用lombok：">
          <n-switch
            v-model:value="strategyState.strategy.entity.enableLombok"
          />
        </n-form-item-gi>
        <n-form-item-gi label="使用record：">
          <n-switch
            v-model:value="strategyState.strategy.entity.enableActiveRecord"
          />
        </n-form-item-gi>
        <n-form-item-gi label="使用静态常量字段：">
          <n-switch
            v-model:value="strategyState.strategy.entity.enableColumnConstant"
          />
        </n-form-item-gi>
        <n-form-item-gi label="移除Boolean类型is前缀：">
          <n-switch
            v-model:value="strategyState.strategy.entity.enableRemoveIsPrefix"
          />
        </n-form-item-gi>
        <!-- <n-form-item-gi label="使用@TableField注解：">
          <n-switch
            disabled
            v-model:value="
              strateEntity.strategy.entity.enableTableFieldAnnotation
            "
          />
        </n-form-item-gi> -->
      </n-grid>

      <n-form-item label="父类包名：">
        <n-input v-model:value="strategyState.strategy.entity.superClass" />
      </n-form-item>
      <n-form-item label="乐观锁属性名：">
        <n-input
          v-model:value="strategyState.strategy.entity.versionPropertyName"
        />
      </n-form-item>
      <n-form-item label="乐观锁字段名：">
        <n-input
          v-model:value="strategyState.strategy.entity.versionColumnName"
        />
      </n-form-item>
      <n-form-item label="逻辑删除属性名：">
        <n-input
          v-model:value="strategyState.strategy.entity.logicDeletePropertyName"
        />
      </n-form-item>
      <n-form-item label="逻辑删除字段名：">
        <n-input
          v-model:value="strategyState.strategy.entity.logicDeleteColumnName"
        />
      </n-form-item>
      <n-form-item label="实体命名策略：">
        <n-select
          v-model:value="strategyState.strategy.entity.naming"
          :options="[
            {
              label: '下划线转驼峰',
              value: 'underline_to_camel',
            },
          ]"
        />
      </n-form-item>
      <n-form-item label="添加忽略字段：">
        <n-dynamic-tags
          v-model:value="strategyState.strategy.entity.addIgnoreColumns"
        />
      </n-form-item>
      <n-grid cols="1">
        <n-form-item-gi label="自动填充字段：">
          <n-dynamic-input
            v-model:value="strategyState.strategy.entity.addTableFills"
            :on-create="onCreate"
            #="{ index, value }"
          >
            <n-input
              placeholder="字段"
              @keydown.enter.prevent
              v-model:value="value.name"
            />
            <div style="height: 34px; line-height: 34px; margin: 0 8px">=</div>
            <n-select
              placeholder="类型"
              @keydown.enter.prevent
              v-model:value="value.value"
              :options="fillOptions"
            />
          </n-dynamic-input>
        </n-form-item-gi>
      </n-grid>
      <n-form-item label="全局主键类型：">
        <n-select
          v-model:value="strategyState.strategy.entity.idType"
          :options="idTypeOptions"
        />
      </n-form-item>
      <n-form-item label="类名格式：">
        <n-input v-model:value="strategyState.strategy.entity.formatFileName" />
      </n-form-item>
    </n-form>
  </n-space>
</template>
<script setup lang="ts">
import { useStrategy } from '@/store/modules/useStrategy'
import {
  NButton,
  NDynamicInput,
  NDynamicTags,
  NForm,
  NFormItem,
  NFormItemGi,
  NGrid,
  NInput,
  NSelect,
  NSpace,
  NSwitch,
  SelectOption,
} from 'naive-ui'

const strategyState = useStrategy()

const idTypeOptions: SelectOption[] = [
  {
    label: 'NONE',
    value: 'NONE',
  },
  {
    label: 'INPUT',
    value: 'INPUT',
  },
  {
    label: 'AUTO',
    value: 'AUTO',
  },
  {
    label: 'ASSIGN_ID',
    value: 'ASSIGN_ID',
  },
  {
    label: 'ASSIGN_UUID',
    value: 'ASSIGN_UUID',
    disabled: true,
  },
]

const fillOptions: SelectOption[] = [
  {
    label: 'DEFAULT',
    value: 'DEFAULT',
  },
  {
    label: 'INSERT',
    value: 'INSERT',
  },
  {
    label: 'UPDATE',
    value: 'UPDATE',
  },
  {
    label: 'INSERT_UPDATE',
    value: 'INSERT_UPDATE',
  },
]

const onCreate = () => {
  return {
    name: '',
    value: 'DEFAULT',
  }
}
</script>

<style lang="scss" scoped></style>
