import { writable } from 'svelte/store'
import type { user } from './interfaces/user'
import type { group } from './interfaces/group'
import type { contact } from './interfaces/contact'

export const userStore = writable<user | null>(null)

export const groupsStore = writable<group[]>([])

export const contactsStore = writable<contact[]>([])
