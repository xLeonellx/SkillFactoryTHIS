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

	selectionSortDouble(ar)

	fmt.Println(ar)
}

func selectionSortDouble(ar []int) {
	maxPosition := len(ar) - 1
	minPosition := 0

	for minPosition <= maxPosition {
		var minIndex = minPosition
		var maxIndex = maxPosition
		for j := minPosition; j <= maxPosition; j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}

		ar[minPosition], ar[minIndex] = ar[minIndex], ar[minPosition]
		if maxIndex == minPosition {
			maxIndex = minIndex
		}
		ar[maxPosition], ar[maxIndex] = ar[maxIndex], ar[maxPosition]

		minPosition++
		maxPosition--
	}
}