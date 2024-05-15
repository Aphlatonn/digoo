package main

import (
	_ "Digoo/cmd/commands"
	_ "Digoo/cmd/engine"
	_ "Digoo/cmd/events"
)

func main() {
	// Keep the app running indefinitely
	<-make(chan struct{})
}
