package namaste

import (
	"fmt"
	"log"
)

func SayNamaste() string {
	s := "Namaste"
	fmt.Println(s)
	log.Println(s)
	var b1 bool = true
	fmt.Printf("Bool: %t\n", b1)
	return s
}
