import { writable } from 'svelte/store'
import type { group } from '$lib/interfaces/group'

export const groupsStore = writable<group[]>([])
