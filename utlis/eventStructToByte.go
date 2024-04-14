package utlis

import (
	"encoding/json"
	"rabotyaga-go-backend/structures"
)

func EventStructToByte(data structures.EventParams) []byte {
	marshal, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return marshal
}
