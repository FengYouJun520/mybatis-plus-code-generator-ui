<template>
  <svg aria-hidden="true">
    <use :href="symbolId" :fill="color" />
  </svg>
</template>

<script lang="ts" setup>
import { computed } from 'vue'

interface SvgIconProps {
  prefix?: string
  name: string
  color?: string
}

const props = withDefaults(defineProps<SvgIconProps>(), {
  color: '#333',
  prefix: 'icon',
})

const symbolId = computed(() => `#${props.prefix}-${props.name}`)
</script>
