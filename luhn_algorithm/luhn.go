package main

import (
	"fmt"
	"regexp"
	"strconv"
)



// second stage of luhn algorithm, check sum
func secondStageVerify(card []int) bool {
	var cardSumm int
	for _, value := range card {
		cardSumm += value
	}
	if cardSumm%10 == 0 {
		return true
	} else {
		return false
	}

}

func lenCardPaired(card string) bool {
	// Received a card without spaces and other symbols except digits, even number of digits
	var resulrVerify []int

	for key, value := range card {

		valueAsInt, err := strconv.Atoi(string(value))
		// конвертим со строки в инт, обрабатываем ошибку
		if err != nil {
			return false
		}
		if key%2 == 0 {
			doubleValueAsInt := valueAsInt * 2
			if doubleValueAsInt > 9 {
				doubleValueAsInt -= 9
			}
			resulrVerify = append(resulrVerify, doubleValueAsInt)
		} else {
			resulrVerify = append(resulrVerify, valueAsInt)
		}
	}

	return secondStageVerify(resulrVerify)
}

func lenCardUnpaired(card string) bool {

	// Received a card without spaces and other symbols except digits, an odd number of digits
	var resulrVerify []int

	for key, value := range card {

		valueAsInt, err := strconv.Atoi(string(value))
		if err != nil {
			return false
		}
		if key%2 != 0 {

			doubleValueAsInt := valueAsInt * 2
			if doubleValueAsInt > 9 {
				doubleValueAsInt -= 9
			}
			resulrVerify = append(resulrVerify, doubleValueAsInt)
		} else {
			resulrVerify = append(resulrVerify, valueAsInt)
		}
	}

	return secondStageVerify(resulrVerify)
}

func verifyCardNumber(card string) bool {
	var filteredCardNumber string

	re := regexp.MustCompile("\\D")
	filteredCardNumber = re.ReplaceAllString(card, "")

	if len(filteredCardNumber)%2 == 0 {
		return lenCardPaired(filteredCardNumber)
	} else {
		return lenCardUnpaired(filteredCardNumber)
	}

}

func main() {

	cardNumber := verifyCardNumber("017 10 1312")

	if cardNumber {
		fmt.Println("Verification passed! Congratulations!")
	} else {
		fmt.Println("Verification failed! Try again!")
	}

}
