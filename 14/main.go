package main

import (
	"fmt"
	"hash/fnv"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {

	var x, y int
	var value string
	var duplicate []string
	m := make(map[uint32]string)
	fmt.Println("Enter first array size:")
	fmt.Scan(&x)

	fmt.Println("Enter second array size:")
	fmt.Scan(&y)

	fmt.Println("Enter first array:")
	for i := 0; i < x; i++ {
		fmt.Scan(&value)
		key := hash(value)
		m[key] = value
	}
		fmt.Println("Enter second array:")
		for i := 0; i < y; i++ {
			fmt.Scan(&value)
			key := hash(value)
			_,ok := m[key]
				if ok {
						duplicate = append(duplicate, value)
					}
					m[key]=value
			}

			fmt.Println(duplicate)

}
