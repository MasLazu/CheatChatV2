package web

type WebsocketReqRes struct {
	MessageType int8 `json:"type"`
	Body        any  `json:"body"`
}
