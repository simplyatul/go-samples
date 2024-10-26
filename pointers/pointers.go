package pointers

import (
	"log"
)

func change(val *int) {
	*val = 55
}

func Checkptr_0() {
	a := 58
	log.Println("value of a before function call is", a) // 58
	b := &a
	change(b)
	log.Println("value of a after function call is", a) // 55
	log.Printf("Type of b %T\n", b)                     // *int
	c := &b
	log.Printf("Type of c %T", c) // **int
}

func Checkptr_1() {
	sls := make([]int, 2, 5)
	sls[0] = 2
	sls[1] = 5

	arr := [3]int{12, 78, 50}

	log.Printf("Type of sls %T\n", sls)                                            // sls []int
	log.Printf("Type of arr %T\n", arr)                                            // a [3]int
	log.Printf("slice contents: %v, len %d and cap %d\n", sls, len(sls), cap(sls)) // slice contents: [2 5], len 2 and cap 5
	log.Printf("Address store in slice %p\n", sls)                                 // 0xc000022100
	log.Printf("Address of slice variable %p\n", &sls)                             // 0xc000010030
	log.Printf("Address of 0th element of slice %p\n", &sls[0])                    // 0xc000022100
	log.Println("-------------------------------")
	log.Printf("array contents: %v, len %d, cap %d\n", arr, len(arr), cap(arr)) // array contents: [12 78 50], len 3, cap 3
	log.Printf("Address store in arr %p\n", arr)                                // Address store in arr %!p([3]int=[12 78 50])
	log.Printf("Address of arr variable %p\n", &arr)                            // Address of arr variable 0xc00013a000
	log.Printf("Address of 0th element of arr %p\n", &arr[0])                   // Address of 0th element of arr 0xc00013a000

}
