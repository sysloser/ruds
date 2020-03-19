package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//arg := os.Args[1]
	countdown := 5
	for {
		fmt.Println("Countdown", countdown)
		countdown = countdown - 1
		if countdown == 0 {
			fmt.Println("Countdown", countdown)
			time.Sleep(1 * time.Second)
			fmt.Println("BOOM")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}
}
