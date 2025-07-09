import { defineStore } from 'pinia'

export const useMarketStore = defineStore('market', {
  state: () => ({ pools: [], tokens: [] }),
})
