package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	ch := make(chan string)

	ninjas := []string{"Ninja1", "Ninja2", "Ninja3", "Ninja1", "Ninja2", "Ninja3"}
	for _, ninja := range ninjas {
		wg.Add(1)
		go attack(ninja, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("this operation took " + time.Since(startTime).String())
}

func attack(ninja string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	time.Sleep(time.Second)
	ch <- ninja
}
