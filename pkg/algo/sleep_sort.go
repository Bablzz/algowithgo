package algo

import "time"

func SleepSort(arr []int) []int {
	out := make([]int, 0, len(arr))
	ch := make(chan int)

	for _, val := range arr {
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}(val)
	}

	for i := 0; i < len(arr); i++ {
		val := <-ch
		out = append(out, val)
	}

	return out
}
