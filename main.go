package main

import "fmt"

func main() {

	a := make([]int, 3)
	fmt.Println("len of a:", len(a))
	// len of a: 3
	fmt.Println("cap of a:", cap(a))
	// cap of a: 3
	fmt.Println("appending 4 to b from a")
	// appending 4 to b from a
	b := append(a, 4)
	fmt.Println("b:", b)
	// b: [0 0 0 4]
	fmt.Println(cap(b))
	fmt.Println("addr of b:", &b[0])
	// addr of b: 0x44a0c0
	fmt.Println("appending 5 to c from a")
	// appending 5 to c from a
	c := append(a, 5)
	fmt.Println("addr of c:", &c[0])
	// addr of c: 0x44a180
	fmt.Println(cap(c))
	fmt.Println("a:", a)
	// a: [0 0 0]
	fmt.Println("b:", b)
	// b: [0 0 0 4]
	fmt.Println("c:", c)
	// c: [0 0 0 5]
}
