import { contactsStore } from '$lib/store/contact'
import { get } from 'svelte/store'

export const nameIfSaved = (email: string) => {
	let contact = [...get(contactsStore)].filter((contact) => contact.email === email)[0]
	return contact ? contact.name : email
}
