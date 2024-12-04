package goroutines

import (
	"log"
	"runtime"
	"time"
)

func say(word string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // Let the scheduler execute other goroutines and comes back at some point.
		log.Printf(word)
	}
}

func RunTimePkg() {
	log.Printf("No of CPU Cores: %d", runtime.NumCPU())
	go say("hi")
	log.Printf("No of Go routines: %d", runtime.NumGoroutine()) // Outputs: 2 (main is a goroutine)
	say("Boss")
	time.Sleep(1 * time.Millisecond)
	log.Printf("No of Go routines: %d", runtime.NumGoroutine())
}
