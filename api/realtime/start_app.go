package realtime

import (
	"encoding/json"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
)

func StartApp(data json.RawMessage) (structures.EventParams, error) {
	return structures.EventParams{Event: types.ResponseStartApp}, nil
}
