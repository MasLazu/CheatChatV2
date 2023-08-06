export interface previewChat {
	chat_id: number
	sender_email: string
	message: string
	created_at: Date
	//null if not group
	group_id: number | null
	group_name: string | null
	//null if not personal
	email: string | null
}
