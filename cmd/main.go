package main

import "ACAutomatic/internal/apiserver"

func main() {
	server, err := apiserver.CreateServer()
	if err != nil {
		panic(err)
	}
	server.Run(":8989")
}
