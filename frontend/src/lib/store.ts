import { writable } from 'svelte/store'
import type { user } from './interfaces/user'
import type { group } from './interfaces/group'
import type { contact } from './interfaces/contact'
import type { previewChat } from './interfaces/previewChat'

export const userStore = writable<user | null>(null)

export const groupsStore = writable<group[]>([])

export const contactsStore = writable<contact[]>([])

export const wsConnStore = writable<WebSocket | null>(null)

export const previewChatStore = writable<previewChat[]>([])
