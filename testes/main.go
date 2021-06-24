package main

import "fmt"

func main() {
	array1 := [...]int{55, 12, 59, 82, 94, 10, 23, 12, 21, 55, 01, 20, 41, 24, 23, 10, 104, 50, 192}
	array2 := [...]int{20, 45, 23, 50, 61, 42, 10, 102, 49, 29, 32, 59, 12, 32, 55, 23, 21, 20, 39, 12}

	arrayUniao := make([]int, 0)
	arrayUniao = retornaUniao(array1, array2)

	fmt.Printf("Uniao = %d.\n", arrayUniao)

	arrayInterseccao := make([]int, 0)
	arrayInterseccao = retornaInterseccao(array1, array2)

	fmt.Printf("Interseccao = %d.\n", arrayInterseccao)
}

func retornaUniao(firstArray [19]int, secondArray [20]int) []int {
	arrayComRepeticao := make([]int, 0)

	for _, elem1 := range firstArray {
		arrayComRepeticao = append(arrayComRepeticao, elem1)
	}

	for _, elem2 := range secondArray {
		arrayComRepeticao = append(arrayComRepeticao, elem2)
	}

	arraySemRepeticao := removeRepeticao(arrayComRepeticao)

	return arraySemRepeticao
}

func retornaInterseccao(firstArray [19]int, secondArray [20]int) []int {
	arrayComRepeticao := make([]int, 0)

	for _, elem1 := range firstArray {
		for _, elem2 := range secondArray {
			if elem1 == elem2 {
				arrayComRepeticao = append(arrayComRepeticao, elem1)
			}
		}
	}

	arraySemRepeticao := removeRepeticao(arrayComRepeticao)

	return arraySemRepeticao
}

func removeRepeticao(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if i != j {
				if slice[i] == slice[j] {
					slice = append(slice[:j], slice[j+1:]...)
				}
			}
		}
	}
	return slice
}
