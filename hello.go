package main

import "fmt"
import "time"

func main() {

	var ops uint64

	for i := 0; i < 1000; i++ {
		go func() {
			ops++
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("ops:", ops)
}
