package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello"
		ch <- "World"

		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}
}
