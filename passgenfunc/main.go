package main

import (
	"fmt"
	"math/rand"
	"time"
)

const DefaultString = "0123456789"
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func generationPass(x int) func() map[string]string {
	generatedLen := x
	generatedMap := make(map[string]string, generatedLen)
	return func() map[string]string {
		min := 0
		max := len(DefaultString)
		var m string
		var sum int
		for {
			m = string(DefaultString[randInt(min, max)])
			for i := 0; i < sum; i++ {
				m += string(DefaultString[randInt(min, max)])
			}
			_, ok := generatedMap[m]
			if ok {
				sum++
			}
			generatedMap[m] = m

			if len(generatedMap) >= (generatedLen - 1) {
				break
			}
		}
		return generatedMap
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	first := generationPass(300)
	fmt.Println("In Main1", first(), len(first()))
}
