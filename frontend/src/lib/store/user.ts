import { writable } from 'svelte/store'
import type { user } from '$lib/interfaces/user'

export const userStore = writable<user | null>(null)
