package itemServer

import (
	"log"
)

func CreateItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			log.Println("item is ", item)
		}
	}()
	return out
}
