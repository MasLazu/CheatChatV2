import { writable } from 'svelte/store'
import type { contact } from '$lib/interfaces/contact'

export const contactsStore = writable<contact[]>([])
