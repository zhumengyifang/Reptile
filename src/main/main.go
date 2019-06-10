package main

import (
	"fmt"
	"goReptile/src/main/Reptile"
	"time"
)

func main() {
	for {
		fmt.Println("Start")

		ch := make(chan int)
		Reptile.GetBlogVisitCount("")
		go loading(ch)
		time.Sleep(time.Minute * 15)
		ch <- 1
	}
}

func loading(ch chan int, ) {
	k := 0
	n := 0
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
				if k == 10 {
					k = 0
					n += 1
					fmt.Print(n)
				} else {
					fmt.Print(".")
					k += 1
				}
			}
		}
	}()
}
