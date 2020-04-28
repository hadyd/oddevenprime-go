package main

import (
	"fmt"
)

func main() {
	array := [30]int{21, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("Bilangan Genap  : %+v, Jumlah = %d\n", genap(array), len(genap(array)))
	fmt.Printf("Bilangan Ganjil : %+v, Jumlah = %d\n", ganjil(array), len(ganjil(array)))
	fmt.Printf("Bilangan Prima  : %+v, Jumlah = %d\n", prima(array), len(prima(array)))
}

func genap(array [30]int) []int {
	var bilanganGenap []int
	for _, i := range array {
		if i != 0 {
			if i%2 == 0 {
				bilanganGenap = append(bilanganGenap, i)
			}
		}
	}

	return bilanganGenap
}

func ganjil(array [30]int) []int {
	var bilanganGanjil []int
	for _, i := range array {
		if i%2 != 0 {
			bilanganGanjil = append(bilanganGanjil, i)
		}
	}
	return bilanganGanjil
}

func prima(array [30]int) []int {
	var bilanganPrima []int
	for a := 1; a < len(array); a++ {
		i := 0
		for b := 1; b < len(array); b++ {
			if a%b == 0 {
				i++
			}
		}
		if (i == 2) && (a != 1) {
			bilanganPrima = append(bilanganPrima, a)
		}
	}
	return bilanganPrima
}
