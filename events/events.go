package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"net"
	"time"
)

type EventParams struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func routing(msg []byte) ([]byte, error) {
	var params EventParams

	if len(msg) < 20 {
		return []byte("You're a damn hacker!"), errors.New(fmt.Sprintf("the connected user transmits incorrect parameters (%d)", len(msg)))
	}

	if err := json.Unmarshal(msg, &params); err != nil {
		return []byte(""), err
	}

	switch params.Event {
	case "test":
		return []byte("Work!"), nil
	default:
		return []byte(""), errors.New("event error")
	}
}

func Listen(conn net.Conn) {
	defer func() {
		logInfo := fmt.Sprintf("[ %d ]: connection closed, address: %s", time.Now().Unix(), conn.RemoteAddr())

		err := conn.Close()
		fmt.Println(logInfo)

		if err != nil {
			return
		}
	}()

	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			break
		}

		msg, err = routing(msg)
		if err != nil {
			fmt.Println(err)
			break
		}

		if err := wsutil.WriteServerMessage(conn, op, msg); err != nil {
			break
		}
	}
}
