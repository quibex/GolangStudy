package main

import (
	"fmt"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		fmt.Print("a")
		c.Wait()
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second)

	c.L.Lock()
	done = true
	c.L.Unlock()

	fmt.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	var m sync.Mutex
	cond := sync.NewCond(&m)

	go read("Reader 1", cond)
	go read("Reader 2", cond)
	go read("Reader 3", cond)
	write("Writer", cond)

	time.Sleep(4 * time.Second)
}
