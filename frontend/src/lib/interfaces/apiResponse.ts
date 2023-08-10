export interface apiResponse {
	code: number
	status: string
}

export interface messageApiResponse extends apiResponse {
	data: {
		message: string
	}
}
