package signals

import (
	"log"
	"os"
	"os/signal"
)

func catchCtrlCInGoroutine() {
	log.Println("Press CTRL+C to stop")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	data := <-sigCh

	log.Printf("Got data %v of type %T\n", data, data)
	log.Println("Existing....")
}

func CatchCtrlC() {
	log.Println("Press CTRL+C to stop")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	go catchCtrlCInGoroutine()
	data := <-sigCh

	log.Printf("Got data %v of type %T\n", data, data)
	log.Println("Existing....")
}
