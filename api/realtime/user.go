package realtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
)

func User(data json.RawMessage) (structures.EventParams, error) {
	var user = structures.User{}
	var err error = nil

	err = json.Unmarshal(data, &user)
	if err != nil {
		err = errors.New(types.ErrorMessageMissingData)
	}

	fmt.Println(user.UserId)

	if err == nil {
		data, _ = json.Marshal(user)
	}

	return structures.EventParams{Event: types.ResponseUser, Data: data}, err
}
