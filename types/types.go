package types

type EventSend struct {
	Event string  `json:"event"`
	Data  *string `json:"data"`
}

const (
	ResponseEventPong = "pong"
)
