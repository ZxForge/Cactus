import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', {
    state: () => ({ c: 0 }),
    getters: {
        count: (state) => state.c,
    },
    actions: {
        increment() {
            this.c++
        },
    },
})