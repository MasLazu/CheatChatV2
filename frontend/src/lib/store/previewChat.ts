import { writable } from 'svelte/store'
import type { previewChat } from '$lib/interfaces/previewChat'

export const previewChatStore = writable<previewChat[]>([])
