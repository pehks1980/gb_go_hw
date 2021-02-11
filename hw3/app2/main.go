package main

import (
	"fmt"
	"math"
)

/*
Написать приложение, которое ищет все простые числа от 0 до N включительно.
	Число N должно быть задано из стандартного потока ввода.
*/

// функция возвращет true если нет ни одного делителя
func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}

	limit := uint64(math.Sqrt(float64(n)))

	for i:= uint64(2); i<= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var aN uint64

	fmt.Printf("Введите, пожалуйста N ")

	_, err := fmt.Scanln(&aN)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		return
	}

	//слайс - список простых чисел
	var numArr []uint64

	for i:= uint64(2); i<= aN; i++ {
		if isPrime(i) {
			//fmt.Printf("Число %d простое\n", i)
			numArr = append(numArr, i)
		}
	}

	for _, prime := range numArr {
		fmt.Printf("Число %d простое\n", prime)
	}

}
