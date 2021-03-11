package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i == j {
				fmt.Print(0)
			} else {
				fmt.Print(1)
			}
		}
		fmt.Println()
	}
}
