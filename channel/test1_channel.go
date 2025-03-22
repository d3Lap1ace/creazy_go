package main

import "fmt"

func main1() {
	ch := make(chan string)

	go func() {
		fmt.Println("sendMsg string")

		ch <- "goland"
	}()

	fmt.Println("Receiving string")
	value := <-ch
	fmt.Println("Recaiving string is:", value)
}
