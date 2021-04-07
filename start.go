package main

import "github.com/kmiloparra/detector-gen-mutante-go/controller"

func main() {

	server := controller.NewServer(":5000")
	server.Listen()
}
