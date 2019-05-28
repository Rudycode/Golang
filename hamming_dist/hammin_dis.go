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

	result, str := hammingDifference(dna1, dna2)
	fmt.Println(str, result)
}

func hammingDifference(a, b string) (uint8, string) {
	firstDNA := make([]rune, len(a))
	secondDNA := make([]rune, len(b))
	var count uint8
	str := "The Hamming distance is"
	err :=  "Error! Strings have different length!"

	if len(a) != len(b){
		return 0, err
	}
	for key, value := range a {
		firstDNA [key] = value
	}
	for key, value := range b {
		secondDNA [key] = value
		if firstDNA[key] != secondDNA[key] {
			count++
		}
	}
	return count, str
}
