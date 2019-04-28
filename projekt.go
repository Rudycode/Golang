package main

import "fmt"

func FizzBuzz(x int64) string {
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
 	return string(x)
}

func main() {
	fmt.Printf(FizzBuzz(1))

}