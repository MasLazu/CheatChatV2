import { writable } from 'svelte/store'
import type { cacheChat } from '$lib/interfaces/chaceChat'

export const currentChatStore = writable<cacheChat | null>(null)
