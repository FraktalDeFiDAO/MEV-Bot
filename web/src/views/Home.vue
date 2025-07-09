<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Known Pools</h1>
    <ul>
      <li v-for="p in pools" :key="p">{{ p }}</li>
    </ul>
    <h1 class="text-2xl font-bold mt-6 mb-4">Known Tokens</h1>
    <ul>
      <li v-for="t in tokens" :key="t">{{ t }}</li>
    </ul>
  </div>
</template>

<script setup>
import { useMarketStore } from '../stores/market'
const store = useMarketStore()
const pools = store.pools
const tokens = store.tokens
</script>
