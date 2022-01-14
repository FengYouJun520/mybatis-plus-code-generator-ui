<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <home />
    </n-message-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import Home from '@/views/home/index.vue'
import { dateZhCN, NConfigProvider, NMessageProvider, zhCN } from 'naive-ui'
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  height: 100%;
}

html,
body {
  height: 100%;
  width: 100%;
}
</style>
