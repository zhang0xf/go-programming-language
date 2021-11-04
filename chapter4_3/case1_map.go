package chapter4_3

import "fmt"

func Map() {
	ages1 := make(map[string]int) // mapping from strings to ints
	fmt.Printf("%T\n", ages1)

	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Printf("%T\n", ages2)

	ages3 := make(map[string]int)
	ages3["alice"] = 31
	ages3["charlie"] = 34

	ages3["alice"] = 32
	fmt.Println(ages3["alice"]) // "32"

	delete(ages3, "alice") // remove element ages["alice"]

	ages3["bob"] = ages3["bob"] + 1 // happy birthday!

	ages3["bob"] += 1

	ages3["bob"]++

	// _ = &ages3["bob"] // compile error: cannot take address of map element
	// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。

	for name, age := range ages3 {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
