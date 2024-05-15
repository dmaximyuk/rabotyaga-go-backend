package main

import (
	"rabotyaga-go-backend/api/realtime"
	"rabotyaga-go-backend/api/realtime/user"
	"rabotyaga-go-backend/mysql/database"
	"rabotyaga-go-backend/server"
	"rabotyaga-go-backend/types"
)

func main() {
	database.New(database.Options{
		Database:       "dev",
		Username:       "root",
		Host:           "localhost",
		MaxConnections: 10,
		Port:           3306,
		Password:       "admin",
	})

	s := server.Init()

	s.On(types.RequestPing, realtime.Ping)
	s.On(types.RequestStartApp, realtime.StartApp)
	s.On(types.RequestUserGet, user.Get)

	s.Listen()
}
