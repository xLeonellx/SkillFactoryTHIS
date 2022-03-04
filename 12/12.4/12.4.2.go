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

	selectionSortByMax(ar)

	fmt.Println(ar)
}


func selectionSortByMax(ar []int) {
	for i := len(ar) - 1; i >= 0; i-- {
		var maxIndex = i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
	}
}