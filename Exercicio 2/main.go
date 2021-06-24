package main

import (
	"fmt"
	"sort"
)

func uniao(origem []int, destino []int) []int {
	var validacao bool

	for _, elem1 := range origem {
		validacao = true
		for _, elemUniao := range destino {
			if elem1 == elemUniao {
				validacao = false
			}
		}
		if validacao {
			destino = append(destino, elem1)
		}
	}

	return destino
}

func interseccao(slice1 []int, slice2 []int) []int {
	sliceInterseccao := make([]int, 0)
	var validacao bool
	for _, elem1 := range slice1 {
		for _, elem2 := range slice2 {
			validacao = false
			if elem1 == elem2 {
				validacao = true
				for _, elemIntersec := range sliceInterseccao {
					if elem1 == elemIntersec {
						validacao = false
					}
				}
			}

			if validacao {
				sliceInterseccao = append(sliceInterseccao, elem1)
			}
		}
	}

	return sliceInterseccao
}

func main() {
	array1 := [...]int{55, 12, 59, 82, 94, 10, 23, 12, 21, 55, 01, 20, 41, 24, 23, 10, 104, 50, 192}
	array2 := [...]int{20, 45, 23, 50, 61, 42, 10, 102, 49, 29, 32, 59, 12, 32, 55, 23, 21, 20, 39, 12}

	sliceUniao := make([]int, 0)
	sliceInterseccao := make([]int, 0)

	sliceUniao = uniao(array1[:], sliceUniao)
	sliceUniao = uniao(array2[:], sliceUniao)
	sliceInterseccao = interseccao(array1[:], array2[:])
	sort.Ints(sliceUniao)
	sort.Ints(sliceInterseccao)

	fmt.Printf("Exercicio 2 - Uniao e Interseccao \n")
	fmt.Println()
	fmt.Printf("Array1: %d \n", array1)
	fmt.Printf("Array2: %d \n", array2)
	fmt.Printf("Uniao : %d \n", sliceUniao)
	fmt.Printf("Interseccao: %d \n", sliceInterseccao)

}
