package main

import (
	"gin-todo-app/server"
	"gin-todo-app/services"
	"os"
)

func main() {
	app, client, err := server.ConnectFirestore()
	if err != nil {
		services.LogError(err)
		os.Exit(1)
	}

	router := server.InitializeRouter(*app, *client)
	router.Run(":8080")
}
