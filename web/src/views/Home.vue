<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Known Pools</h1>
    <ul>
      <li v-for="p in pools" :key="p.address">{{ p.address }} - {{ p.token0 }} / {{ p.token1 }}</li>
    </ul>
    <form @submit.prevent="addPool" class="my-2">
      <input v-model="poolAddr" placeholder="Pool" class="border mr-1" />
      <input v-model="poolT0" placeholder="Token0" class="border mr-1" />
      <input v-model="poolT1" placeholder="Token1" class="border mr-1" />
      <button class="border px-2">Add Pool</button>
    </form>
    <h1 class="text-2xl font-bold mt-6 mb-4">Known Tokens</h1>
    <ul>
      <li v-for="t in tokens" :key="t">{{ t }}</li>
    </ul>
    <form @submit.prevent="addToken" class="mt-2">
      <input v-model="tokenAddr" placeholder="Token" class="border mr-1" />
      <button class="border px-2">Add Token</button>
    </form>
  </div>
</template>

<script setup>
import { useMarketStore } from '../stores/market'
import { ref, onMounted } from 'vue'
const store = useMarketStore()
const pools = store.pools
const tokens = store.tokens
const poolAddr = ref('')
const poolT0 = ref('')
const poolT1 = ref('')
const tokenAddr = ref('')
onMounted(() => {
  store.refresh()
})
function addPool() {
  store.addPool(poolAddr.value, poolT0.value, poolT1.value)
  poolAddr.value = poolT0.value = poolT1.value = ''
}
function addToken() {
  store.addToken(tokenAddr.value)
  tokenAddr.value = ''
}
</script>
