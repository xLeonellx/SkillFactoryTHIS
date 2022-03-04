package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


func bubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

func bubbleSortWithBreak (ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		swapped := false
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar [j] {
				swapped = true
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
		if !swapped {
			break
		}
	}
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

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		var minIndex = i
		for j := i+1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
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

func insertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}
	for i := 1; i < len(ar); i++ {
		for j := i; j>0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
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

func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar) - 1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left + 1:])

	return
}