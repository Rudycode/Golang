package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 017 10 1312  || 5168 7423 3010 5911

func main() {

	cardNumber := verifyCardNumber("017 10 1312")

	if cardNumber {
		fmt.Println("Verification passed! Congratulations!")
	} else {
		fmt.Println("Verification failed! Try again!")
	}

}

func verifyCardNumber(card string) bool {
	var filteredCardNumber string
	var resulrVerify []int

	re := regexp.MustCompile("\\D")
	filteredCardNumber = re.ReplaceAllString(card, "")

	fmt.Println(len(filteredCardNumber), filteredCardNumber)
	for key, value := range filteredCardNumber {

		valueAsInt, err := strconv.Atoi(string(value))
		// конвертим со строки в инт, обрабатываем ошибку
		if err != nil {
			return false
		}

		if len(filteredCardNumber)%2 == 0 {
			// pair len card
			if key%2 == 0 {
				doubleValueAsInt := valueAsInt * 2
				if doubleValueAsInt > 9 {
					doubleValueAsInt -= 9
				}
				resulrVerify = append(resulrVerify, doubleValueAsInt)
			} else {
				resulrVerify = append(resulrVerify, valueAsInt)
			}

		} else {
			// unpair len card
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
	}

	return secondStageVerify(resulrVerify)
}

// second stage of luhn algorithm, check sum, 
func secondStageVerify(card []int) bool {
	fmt.Println(card)

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
