package realtime

import (
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
)

func Balance() (structures.EventParams, error) {
	return structures.EventParams{Event: types.ResponseBalance}, nil
}
