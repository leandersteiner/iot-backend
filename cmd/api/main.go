package main

import (
	"fmt"
	"github.com/leandersteiner/iot-backend/internal/server"
	"log"
)

func main() {
	s := server.NewServer()

	waitChan := make(chan interface{})

	go func(waitChan chan<- interface{}) {
		err := s.Run("8080")
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		waitChan <- struct{}{}
	}(waitChan)

	fmt.Println("Listening on localhost:8080")
	<-waitChan
}
