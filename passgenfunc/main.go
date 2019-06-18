package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
a function with closure, in which a map with unique keys is generated, with given a length. 
If the key/value already exists during generation, the generated key is increased by an order 
and this is repeated until the map is ready, the length specified.
*/

const DefaultString = "0123456789" // constant where we will take a random number
type emptyStruct struct{}

// generate random integer
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// returns the function() map[string]string where a unique key-value of length "X" was generated
func generationPass(x int) func() map[string]string {
	generatedLen := x
	generatedMap := make(map[string]string, generatedLen)
	return func() map[string]string{
		min := 0
		max := len(DefaultString)
		var resultValue string
		var degree int
		 for {
			resultValue = string(DefaultString[randInt(min, max)])
			for i := 0; i < degree; i++ {
				resultValue += string(DefaultString[randInt(min, max)])
			}
			_, ok := generatedMap[resultValue]
			if ok {
				degree++
			}
			generatedMap[resultValue] = resultValue

			if len(generatedMap) >= generatedLen {
				break
			}
		}
		return generatedMap
	}
}

func main() {
	t := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	 firstMap := generationPass(300)
	fmt.Println("In Main1", len(firstMap()))
	fmt.Println(time.Since(t))
}
