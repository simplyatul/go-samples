package main

import (
	_ "fmt" // w/o _, compiler will throw an error. Ok to use in development code. Must be avoided in production grade code
	_ "go-samples/goroutines"
	"go-samples/namaste"
	_ "go-samples/pointers"
	_ "go-samples/signals"

	"log"
)

func main() {
	// Set log flags to include date, time with milliseconds, short file name, and line number
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	namaste.SayNamaste()
	// pointers.Checkptr_1()

	// structs.Struct_0()
	//signals.CatchCtrlC()
	// goroutines.RunTimePkg()
	// channels.ChannelsTestOne()

	//interfaces.CheckInterfaces_0()
	// interfaces.CheckInterfaces_1()
	//interfaces.CheckInterfaces_2()
}
