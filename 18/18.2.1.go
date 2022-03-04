package countersemaphore

import (
	"fmt"
	"sync"
)

func main() {
	const routinesCount = 5
	const printNumberTimes = 10
	var wg sync.WaitGroup
	wg.Add(routinesCount)
	for routine := 1; routine <= routinesCount; routine++ {
		go func(routine int) {
			for i := 0; i < printNumberTimes; i++ {
				fmt.Println(routine)
			}
			wg.Done()
		}(routine)
	}
	wg.Wait()
}