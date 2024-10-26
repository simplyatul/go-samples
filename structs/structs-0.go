package structs

import (
	"log"
)

type Item struct {
	itemName        string
	itemDescription string
	cost            float32
}

type localItem struct {
	itemName        string
	itemDescription string
	cost            float32
}

func Struct_0() {
	log.Println("----------------------------")

	itPtr := &Item{
		itemName:        "Pencil",
		itemDescription: "Pencil",
		cost:            10.2,
	}
	log.Printf("Type of itPtr => %T", itPtr) // Type of itPtr => *structs.Item

	itLoPtr := &localItem{}
	log.Printf("Type of itLoPtr => %T", itLoPtr) // Type of itLoPtr => *structs.localItem

	anonymousStruct := struct {
		f1 string
		f2 int
	}{
		f1: "Field 1",
		f2: 10,
	}
	ptrToanonymousStruct := &anonymousStruct

	log.Printf("Type of anonymousStruct => %T", anonymousStruct)           // Type of anonymousStruct => struct { f1 string; f2 int }
	log.Printf("Type of ptrToanonymousStruct => %T", ptrToanonymousStruct) // Type of ptrToanonymousStruct => *struct { f1 string; f2 int }

	type anonymousFieldStruct struct {
		string
		int
	}

	aFS := anonymousFieldStruct{string: "Name", int: 20}
	log.Printf("Type of anonymousFieldStruct => %T", aFS) // Type of anonymousFieldStruct => structs.anonymousFieldStruct

	ptrToaFS := &aFS
	log.Printf("Type of ptrToaSF => %T", ptrToaFS) // Type of ptrToaSF => *structs.anonymousFieldStruct
}
