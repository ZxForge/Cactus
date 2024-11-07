import { defineStore } from 'pinia'

export const useCounterStore = defineStore('emails', {
    state: () => (
        { 
            emails: [
            ]
        }
    ),
    getters: {
        get: (state) => state.c,
    },
    actions: {
        add: (state) => state.c,
        delete: (state) => state.c,
    },
})