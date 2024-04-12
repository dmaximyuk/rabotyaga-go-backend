package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/api/realtime"
	"rabotyaga-go-backend/utlis"
	"time"
)

type Events struct {
	Conn net.Conn
}

type EventParams struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func (e *Events) Listener(msg []byte) ([]byte, error) {
	var params EventParams

	if len(msg) < 20 {
		return []byte("You're a damn hacker!"), errors.New(fmt.Sprintf("the connected user transmits incorrect parameters (%d)", len(msg)))
	}

	if err := json.Unmarshal(msg, &params); err != nil {
		return []byte(""), err
	}

	switch params.Event {
	case "ping":
		data := realtime.Ping()
		message := utlis.EventStructToByte(data)
		return message, nil
	default:
		return []byte(""), errors.New("event error")
	}
}

func (e *Events) New() {
	defer func() {
		logInfo := fmt.Sprintf("[ %d ]: connection closed, address: %s", time.Now().Unix(), e.Conn.RemoteAddr())

		err := e.Conn.Close()
		fmt.Println(logInfo)

		if err != nil {
			return
		}
	}()

	for {
		msg, op, err := wsutil.ReadClientData(e.Conn)
		if err != nil {
			break
		}

		msg, err = e.Listener(msg)
		if err != nil {
			fmt.Println(err)
			break
		}

		if err := wsutil.WriteServerMessage(e.Conn, op, msg); err != nil {
			break
		}
	}
}
