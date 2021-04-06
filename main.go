package main

import "github.com/kmiloparra/detector-gen-mutante-go/servidor"

func main() {



	server := servidor.NewServer(":5000")
	server.Listen()
}
