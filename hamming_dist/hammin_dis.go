package main

import "fmt"

/*

GAGCCTACTAACGGGAT
CATCGTAATGACGGCCT
^ ^ ^  ^ ^    ^^

*/

func main() {

	dna1 := "GAGCCTACTAACGGGAT"
	dna2 := "CATCGTAATGACGGCCT"


	result := hammingDifference(dna1, dna2)
	fmt.Printf("The Hamming distance is %d\n", result)
}

func hammingDifference(a, b string) uint8 {
	firstDNA := make([]rune, len(a))
	secondDNA := make([]rune, len(b))
	var count uint8

	fmt.Println(len(a), len(b))

	for key, value := range a {
		firstDNA [key] = value
	}
	for key, value := range b {
		secondDNA [key] = value
		if firstDNA[key] != secondDNA[key] {
			count++
		}
	}

	fmt.Println(firstDNA, secondDNA)

	return count
}
