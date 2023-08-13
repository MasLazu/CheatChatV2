import { writable } from 'svelte/store'

export const websocketStore = writable<WebSocket | null>(null)
