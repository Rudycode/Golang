package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	first := generationPass()
	fmt.Println("In Main1", first())
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func generationPass() func() map[int]int {
	generatedLen := 300
	generatedMap := make(map[int]int, generatedLen)
	return func() map[int]int {
		min := 0
		max := 10
		var m int
		count := 0
		var duplicatedValue []int
		for i := 1; i <= generatedLen; i++ {
			a, ok := generatedMap[m]
			if ok {
				fmt.Println(a)
				max *= 10
				count++
				m = randInt(min, max)
				generatedMap[i] = m
				duplicatedValue = append(duplicatedValue, a)
			} else {
				m = randInt(min, max)
				generatedMap[i] = m
			}
		}
		return generatedMap
	}
}
