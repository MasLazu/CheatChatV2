import type { chat } from './cacheChat'

export interface cacheChat {
	//null if not group
	groupId: number | null
	groupName: string | null
	//null if not personal
	email: string | null
	chats: chat[]
}
