package main

import (
	"fmt"
	"time"
)

func main() {
	go count("R4ha")
	count("lucu")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
