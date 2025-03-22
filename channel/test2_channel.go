package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Println("len(ch)=", len(ch), ",cap(ch)=", cap(ch))

	go func() {
		defer fmt.Println("func end")

		for i := 1; i < 20; i++ {
			ch <- i
			fmt.Println("sendData:", i, "len(ch):", len(ch), ",cap(ch):", cap(ch))
		}

	}()

	// time.Sleep(2 * time.Second)

	for i := 0; i < 19; i++ {
		num := <-ch
		time.Sleep(1 * time.Second)
		fmt.Println("num:", num)
	}

	fmt.Println("main end")
}
