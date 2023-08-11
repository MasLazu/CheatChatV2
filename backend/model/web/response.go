package web

type JsonResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
