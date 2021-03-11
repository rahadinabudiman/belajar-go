package main

import "fmt"

func main() {
	bilangan := 1
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i == j {
				fmt.Print(bilangan)
			} else {
				fmt.Print(0)
			}
		}
		bilangan = bilangan + 1
		fmt.Println()
	}
}
