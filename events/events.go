package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/api/realtime"
	"rabotyaga-go-backend/structures"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utlis"
	"time"
)

type Events struct {
	Conn net.Conn
}

func (e *Events) Listener(msg []byte) ([]byte, error) {
	var err error = nil
	var message structures.EventParams

	if len(msg) < 20 {
		err = errors.New(types.ErrorMessageMsgLength)
	}

	if err = json.Unmarshal(msg, &message); err != nil {
		err = errors.New(types.ErrorMessageParseData)
	}

	if len(message.Event) < 4 {
		err = errors.New(types.ErrorMessageMissingEvent)
	}

	if err == nil {
		switch message.Event {
		case types.RequestStartApp:
			message, err = realtime.Ping()
			break
		case types.RequestPing:
			message, err = realtime.Ping()
			break
		case types.RequestBalance:
			message, err = realtime.Balance()
			break
		case types.RequestUser:
			message, err = realtime.User(message.Data)
			break
		default:
			err = errors.New(types.ErrorMessageMissingEvent)
			break
		}
	}

	if err != nil {
		errorData, _ := json.Marshal(structures.Error{Msg: err.Error()})
		message = structures.EventParams{
			Event: types.ResponseError,
			Data:  errorData,
		}
	}

	msg = utlis.EventStructToByte(message)
	return msg, err
}

func (e *Events) New() {
	defer func() {
		if err := e.Conn.Close(); err != nil {
			logInfo := fmt.Sprintf("[ %d ]: connection closed, address: %s", time.Now().Unix(), e.Conn.RemoteAddr())
			fmt.Println(logInfo)
		}
	}()

	for {
		msg, op, err := wsutil.ReadClientData(e.Conn)
		if err != nil {
			break
		}

		msg, err = e.Listener(msg)

		err = wsutil.WriteServerMessage(e.Conn, op, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
