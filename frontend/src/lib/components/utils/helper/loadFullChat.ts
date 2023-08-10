import axios, { AxiosError } from 'axios'
import type { AxiosResponse } from 'axios'
import type { apiResponse } from '$lib/interfaces/apiResponse'
import { cacheChatStore } from '$lib/components/store/cacheChat'
import { groupsStore } from '$lib/components/store/group'
import { currentChatStore } from '$lib/components/store/currentChat'
import { get } from 'svelte/store'
import type { messageApiResponse } from '$lib/interfaces/apiResponse'

interface fullChatResponse extends apiResponse {
	data: {
		id: number
		sender_email: string
		message: string
		created_at: string
	}[]
}

//try catch needed
export const loadFullChat = async (key: number | string) => {
	if (typeof key === 'number') {
		let cache = [...get(cacheChatStore)].filter((cacheChat) => cacheChat.groupId === key)[0]
		let groupName = [...get(groupsStore)].filter((group) => group.id === key)[0]?.name
		if (!cache) {
			try {
				const result: AxiosResponse<fullChatResponse> = await axios.get(
					`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/chats/group/${key}`,
					{ withCredentials: true }
				)
				cache = {
					groupId: key,
					groupName: groupName,
					email: null,
					chats: result.data.data.map((chat) => {
						return {
							id: chat.id,
							senderEmail: chat.sender_email,
							message: chat.message,
							createdAt: new Date(chat.created_at)
						}
					})
				}
			} catch (err) {
				const errors = err as AxiosError<messageApiResponse>
				if (errors.response?.status === 404) {
					cache = {
						groupId: key,
						groupName: groupName,
						email: null,
						chats: []
					}
				} else {
					throw new Error('Failed to get chat data')
				}
			}
			cacheChatStore.update((prev) => [...prev, cache])
		}
		currentChatStore.set(cache)
	} else if (typeof key === 'string') {
		let cache = [...get(cacheChatStore)].filter((cacheChat) => cacheChat.email === key)[0]
		if (!cache) {
			try {
				const result: AxiosResponse<fullChatResponse> = await axios.get(
					`http://${import.meta.env.VITE_BACKEND_DOMAIN}/api/login/chats/personal/${key}`,
					{ withCredentials: true }
				)
				cache = {
					groupId: null,
					groupName: null,
					email: key,
					chats: result.data.data.map((chat) => {
						return {
							id: chat.id,
							senderEmail: chat.sender_email,
							message: chat.message,
							createdAt: new Date(chat.created_at)
						}
					})
				}
			} catch (err) {
				const errors = err as AxiosError<messageApiResponse>
				if (errors.response?.status === 404) {
					cache = {
						groupId: null,
						groupName: null,
						email: key,
						chats: []
					}
				} else {
					throw new Error('Failed to get chat data')
				}
			}
			cacheChatStore.update((prev) => [...prev, cache])
		}
		currentChatStore.set(cache)
	}
	//console.log(get(cacheChatStore))
	//console.log(get(currentChatStore))
}
