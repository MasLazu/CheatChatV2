import { writable } from 'svelte/store'
import type { cacheChat } from '$lib/interfaces/chaceChat'

export const cacheChatStore = writable<cacheChat[]>([])
