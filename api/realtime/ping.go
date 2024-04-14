package realtime

import (
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
)

func Ping() (structures.EventParams, error) {
	return structures.EventParams{Event: types.ResponsePong}, nil
}
