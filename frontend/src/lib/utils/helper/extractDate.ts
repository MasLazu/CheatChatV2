export const extractDate = (chatDate: Date): string => {
	const currentDate = new Date()
	const currentDay = currentDate.getDate()
	const currentMonth = currentDate.getMonth()
	const currentYear = currentDate.getFullYear()

	const chatMinute = chatDate.getMinutes()
	const chatHour = chatDate.getHours()
	const chatDay = chatDate.getDate()
	const chatMonth = chatDate.getMonth()
	const chatYear = chatDate.getFullYear()

	if (chatYear === currentYear) {
		if (chatMonth === currentMonth) {
			if (chatDay === currentDay) {
				return `${chatHour}:${chatMinute < 10 ? '0' : ''}${chatMinute}`
			} else if (currentDay - chatDay < 7) {
				return chatDate.toLocaleDateString('en-US', { weekday: 'short' })
			}
		}
		return `${chatDay < 10 ? '0' : ''}${chatDay}/${chatMonth + 1 < 10 ? '0' : ''}${chatMonth + 1}`
	}
	return `${chatDay < 10 ? '0' : ''}${chatDay}/${
		chatMonth + 1 < 10 ? '0' : ''
	}${chatMonth}/${chatYear}`
}
