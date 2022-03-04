package main

import (
"fmt"
"math/rand"
"time"
)

func init() {
rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
ar := make([]int, 50)
for i := range ar {
ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
}

	ar = mergeSort(ar)

	fmt.Println(ar)
}


func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	middle := len(ar)/2

sortedAr := make([]int, 0, len(ar))
left, right := mergeSort(ar[:middle]), mergeSort(ar[middle:])

var i, j = 0, 0
for i < len(left) && j < len(right) {
if left[i] > right[j] {
sortedAr = append(sortedAr, right[j])
j++
} else {
sortedAr = append(sortedAr, left[i])
i++
}
}

sortedAr = append(sortedAr, left[i:]...)
sortedAr = append(sortedAr, right[j:]...)

return sortedAr
}