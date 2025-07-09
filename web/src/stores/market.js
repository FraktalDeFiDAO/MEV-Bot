import { defineStore } from "pinia";

const API = import.meta.env.VITE_API_URL || "";

export const useMarketStore = defineStore("market", {
  state: () => ({ pools: [], tokens: [] }),
  actions: {
    async refresh() {
      this.tokens = await fetch(`${API}/tokens`).then((r) => r.json());
      this.pools = await fetch(`${API}/pools`).then((r) => r.json());
    },
    async addToken(addr) {
      await fetch(`${API}/tokens`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ address: addr }),
      });
      this.refresh();
    },
    async addPool(address) {
      await fetch(`${API}/pools`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ address }),
      });
      this.refresh();
    },
  },
});
