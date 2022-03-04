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
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}


	bubbleSortRecursive (ar)

	fmt.Println(ar)
}


func bubbleSortRecursive (ar []int) {
	if len(ar) == 1 {
		return
	}
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i+1], ar[i] = ar[i], ar[i+1]
		}
	}
	bubbleSortRecursive(ar[:len(ar)-1])
}