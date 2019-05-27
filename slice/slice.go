package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x, _ := range a {
		a[x] = make([]uint8, dx)
		for y, _ := range a[x] {
			a[x][y] = uint8(x^y)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}