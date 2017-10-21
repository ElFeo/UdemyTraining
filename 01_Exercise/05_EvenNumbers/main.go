package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("even number: ",i)
		}
		if i%3 == 0 {
			fmt.Println("Divisible by 3: ",i)
		}

	}
}
