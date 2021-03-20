package main

import "github.com/qjvtenkroode/mqtt2cribl/internal/mqtt2cribl"

func main() {
	client := mqtt2cribl.StartReceiving()

	c := make(chan struct{})
	<-c

	client.Disconnect(100)
}
