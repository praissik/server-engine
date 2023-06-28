package engine

import (
	"log"
	"time"
)

func Heartbeat(duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println(".")
			}
		}
	}()
}
