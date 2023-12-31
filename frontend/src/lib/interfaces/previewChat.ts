export interface previewChat {
	chatId: number
	senderEmail: string
	message: string
	createdAt: Date
	//null if not group
	groupId: number | null
	groupName: string | null
	//null if not personal
	email: string | null
}
