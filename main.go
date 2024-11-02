package main

import (
	"To-Do-App/handlers"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	sigchan := make(chan os.Signal, 1)

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		<-sigchan
		handlers.Exit()
	}()
	
	handlers.Run()
}










