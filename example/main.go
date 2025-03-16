package main

import (
	"time"

	"github.com/jimmicro/argo"
)

func main() {
	testRgo()
	time.Sleep(time.Second)
}

func testRgo() {
	argo.Go(func() {
		panic("test")
	})
}
