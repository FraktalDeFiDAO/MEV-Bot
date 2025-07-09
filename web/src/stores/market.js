import { defineStore } from 'pinia'

export const useMarketStore = defineStore('market', {
  state: () => ({ pools: [], tokens: [] }),
  actions: {
    async refresh() {
      this.tokens = await fetch('/tokens').then(r => r.json())
      this.pools = await fetch('/pools').then(r => r.json())
    },
    async addToken(addr) {
      await fetch('/tokens', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ address: addr }) })
      this.refresh()
    },
    async addPool(address, token0, token1) {
      await fetch('/pools', { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify({ address, token0, token1 }) })
      this.refresh()
    },
  }
})
