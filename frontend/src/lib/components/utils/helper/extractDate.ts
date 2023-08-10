export const extractDate = (chatDate: Date): string => {
	const currentDate = new Date()
	if (
		chatDate.getFullYear() === currentDate.getFullYear() &&
		chatDate.getMonth() === currentDate.getMonth()
	) {
		if (chatDate.getDate() === currentDate.getDate()) {
			return (
				chatDate?.getHours().toString().padStart(2, '0') +
				':' +
				chatDate?.getMinutes().toString().padStart(2, '0')
			)
		} else if (currentDate.getDate() - chatDate.getDate() < 7) {
			return chatDate.toLocaleDateString('en-US', { weekday: 'long' })
		}
	}
	return (
		chatDate.getDate().toString().padStart(2, '0') +
		'/' +
		chatDate.getMonth().toString().padStart(2, '0') +
		'/' +
		chatDate.getFullYear().toString()
	)
}
