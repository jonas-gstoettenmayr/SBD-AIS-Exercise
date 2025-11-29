package main

import (
	"time"
	"exc8/server"
	"exc8/client"
)

func main() {
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err!= nil {
			println("Womp womp")
			return
		}
	}()
	time.Sleep(1 * time.Second)
	// todo start client
	if clnt, err := client.NewGrpcClient(); err == nil {
		clnt.Run()
	}
	println("Orders complete!")
}
