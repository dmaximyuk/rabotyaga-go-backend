package realtime

import (
	"encoding/json"
	"errors"
	"rabotyaga-go-backend/mysql"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
)

func User(data json.RawMessage) (structures.EventParams, error) {
	var userData = new(structures.User)
	var err error = nil

	err = json.Unmarshal(data, &userData)
	if err != nil {
		err = errors.New(types.ErrorMessageMissingData)
	}

	if err == nil {
		user, _ := mysql.USER_GET_BY_UID(userData.UserId)
		data, _ = json.Marshal(user)
	}

	return structures.EventParams{Event: types.ResponseUser, Data: data}, err
}
