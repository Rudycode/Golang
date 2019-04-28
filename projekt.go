package main

import "fmt"

func FizzBuzz(x int64) int64 {
	for i := x; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Printf("Fizz")
		}
		if i % 5 == 0 {
			fmt.Printf("Buzz")
		}
		if i % 3 != 0 && i % 5 != 0 {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
 	return x
}

func main() {
	fmt.Println(FizzBuzz(1))

}