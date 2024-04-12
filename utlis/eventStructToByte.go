package utlis

import (
	"encoding/json"
	"rabotyaga-go-backend/types"
)

func EventStructToByte(data types.EventSend) []byte {
	marshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return marshal
}
