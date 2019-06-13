package main

import (
	"fmt"
	"goReptile/src/Config"
	"goReptile/src/main/Reptile"
	"time"
)

func main() {
	config := Config.GetConfig()
	fmt.Println(config)
	for {
		fmt.Println("Start")
		ch := make(chan int)
		Reptile.GetBlogVisitCount(config.URL)
		go loading(ch)
		time.Sleep(time.Minute * time.Duration(config.ExecutionInterval))
		ch <- 1
	}
}

func loading(ch chan int, ) {
	fmt.Print("Loading")
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println()
				fmt.Println("Loading Over")
				return
			default:
				time.Sleep(time.Second)
				fmt.Print(".")
			}
		}
	}()
}
