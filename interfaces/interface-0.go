package interfaces

import (
	"log"
)

type shape interface {
	area() float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return c.radius * c.radius
}

func CheckInterfaces_0() {
	c := circle{radius: 33.3}
	var s shape
	log.Printf("Type %T Value %v", s, s) // Type <nil> Value <nil>
	s = c
	log.Printf("Type %T Value %v Area %v", s, s, s.area()) // Type interfaces.circle Value {33.3} Area 1108.8899

}

/*
 * =========================================================================
 */
func assert(i interface{}) {
	v, ok := i.(int)
	log.Println(v, ok)

}

func CheckInterfaces_1() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}

/*
 * =========================================================================
 */

func myType(i interface{}) {
	switch v := i.(type) {
	case string:
		log.Printf("string value %s\n", i.(string))
		log.Printf("Type = %T, value = %v\n", v, v)
	case int:
		log.Printf("int value %d\n", i.(int))
	default:
		log.Printf("Unknown type\n")
	}
}

func CheckInterfaces_2() {
	myType("XYZ")
	myType(11)
	myType(5.2)
}
