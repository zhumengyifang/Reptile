package main

import (
	"goReptile/src/main/Reptile"
	"time"
)

func main() {
	for {
		Reptile.GetBlogVisitCount("")
		time.Sleep(time.Minute * 30)
	}
}
