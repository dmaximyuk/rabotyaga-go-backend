package realtime

import "rabotyaga-go-backend/types"

func Ping() types.EventSend {
	return types.EventSend{Event: types.ResponseEventPong}
}
