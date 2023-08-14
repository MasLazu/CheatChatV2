import { cacheChatStore } from '$lib/store/cacheChat'
import { currentChatStore } from '$lib/store/currentChat'
import { previewChatStore } from '$lib/store/previewChat'
import { userStore } from '$lib/store/user'
import { get } from 'svelte/store'
import { stringToDate } from './helper/stringToDate'

export const websocketManager = (data: any) => {
	// console.log(data)
	// const date = new Date(data.created_at)
	// const CreatedAt = date.getTime() + date.getTimezoneOffset() * 60 * 1000

	// new chat data
	const chatData = {
		id: data.id,
		senderEmail: data.sender_email,
		message: data.message,
		createdAt: new Date(data.created_at)
	}

	//update cache
	cacheChatStore.update((prev) => {
		let cache = [...prev].findIndex(
			(cacheChat) =>
				(cacheChat.groupId && cacheChat.groupId === data.group_id) ||
				(cacheChat.email &&
					(cacheChat.email === data.sender_email || cacheChat.email === data.receiver_email))
		)

		if (cache === -1) {
			prev.push({
				groupId: data.group_id,
				groupName: data.group_name,
				email: data.group_id
					? null
					: data.sender_email === get(userStore)?.email
					? data.receiver_email
					: data.sender_email,
				chats: [chatData]
			})
		} else {
			prev[cache].chats.push(chatData)
		}

		return prev
	})

	//update current chat if open
	if (
		(get(currentChatStore)?.groupId && get(currentChatStore)?.groupId === data.group_id) ||
		(get(currentChatStore)?.email &&
			(get(currentChatStore)?.email === data.receiver_email ||
				get(currentChatStore)?.email === data.sender_email))
	) {
		currentChatStore.set(
			get(cacheChatStore)[
				get(cacheChatStore).findIndex(
					(cacheChat) =>
						(data.group_id && cacheChat.groupId === data.group_id) ||
						(cacheChat.email &&
							(cacheChat.email === data.sender_email || cacheChat.email === data.receiver_email))
				)
			]
		)
	}

	//update preview chat
	previewChatStore.update((prev) => {
		const index = get(previewChatStore).findIndex(
			(previewChat) =>
				(previewChat.groupId && previewChat.groupId === data.group_id) ||
				(previewChat.email &&
					(previewChat.email === data.sender_email || previewChat.email === data.receiver_email))
		)

		const previewChat = {
			email: data.group_id
				? null
				: data.sender_email === get(userStore)?.email
				? data.receiver_email
				: data.sender_email,
			groupName: data.group_id ? prev[index]?.groupName : null,
			groupId: data.group_id ?? null,
			message: data.message,
			createdAt: new Date(data.created_at),
			chatId: data.id,
			senderEmail: data.sender_email
		}

		if (index === -1) {
			prev.push(previewChat)
		} else {
			prev[index] = previewChat
		}

		return [...prev].sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime())
	})
}
