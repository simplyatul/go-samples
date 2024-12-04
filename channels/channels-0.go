package channels

import (
	"log"
	"runtime"
	"time"
)

// Ref: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/02.7.html

func ChannelsTestOne() {
	log.Println("Here")
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				log.Println(v)
			case <-time.After(5 * time.Second):
				log.Println("timeout")
				o <- true
				break
			}
		}
	}()
	log.Printf("No of Go routines: %d", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	log.Println("Here")
	log.Printf("No of Go routines: %d", runtime.NumGoroutine())
	c <- 2
	<-o
	log.Printf("No of Go routines: %d", runtime.NumGoroutine())
}
