package structures

import (
	"encoding/json"
	"rabotyaga-go-backend/types"
)

type EventParams struct {
	Event types.EventType `json:"event"`
	Data  json.RawMessage `json:"data,omitempty"`
}

type User struct {
	UserId   uint   `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

type Error struct {
	Msg string `json:"msg"`
}
