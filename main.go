package main

import (
	"fmt"
	"github.com/gobwas/ws"
	"log"
	"net/http"
	"rabotyaga-go-backend/events"
)

func main() {
	fmt.Println("Server started!")
	err := http.ListenAndServe("0.0.0.0:3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Connection")
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Panicln("Upgrade HTTP error")
			return
		}

		var e = events.Events{conn}

		go e.New()
	}))

	if err != nil {
		log.Panicln("Server error starting")
		return
	}
}
