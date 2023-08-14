export const stringToDate = (date: string) => {
	const dateObj = new Date(date)
	const dateInMs = dateObj.getTime() + dateObj.getTimezoneOffset() * 60 * 1000
	return new Date(dateInMs)
}
