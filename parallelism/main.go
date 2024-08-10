package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	ninjas := []string{"Ninja1", "Ninja2", "Ninja3", "Ninja4"}
	for _, ninja := range ninjas {
		attack(ninja)
	}
	fmt.Println("this operation took " + time.Since(startTime).String())
}

func attack(ninja string) {
	time.Sleep(time.Second)
	fmt.Println("Attacking", ninja)
}
