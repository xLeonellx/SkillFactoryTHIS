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


	bubbleSortReversed (ar)

fmt.Println(ar)
}


func bubbleSortReversed (ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		swapped := false
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] < ar [j] {
				swapped = true
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
		if !swapped {
			break
		}
	}
}