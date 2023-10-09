package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.RWMutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.Lock()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func (c *Counter) GetValue(wg *sync.WaitGroup) {
	defer wg.Done()

	c.m.RLock() //блокирует для записи другими горутинами, но разрешает чтение
	defer c.m.RUnlock()
	fmt.Println("Get value:", c.value)
	//time.Sleep(4000 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}

	wg.Add(4)

	go c.Update(10, &wg)
	go c.GetValue(&wg)
	go c.GetValue(&wg)
	go c.GetValue(&wg)

	wg.Wait()
}
