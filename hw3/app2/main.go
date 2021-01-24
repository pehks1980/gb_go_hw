package main

import (
	"fmt"
	"math"
	"os"
)

/*
Написать приложение, которое ищет все простые числа от 0 до N включительно.
	Число N должно быть задано из стандартного потока ввода.
*/

// функция возвращет true если нет ни одного делителя
func isPrime(n int32) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	limit := int32(math.Sqrt(float64(n)))

	i := int32(2)

	for i <= limit {
		if n%i == 0 {
			return false
		}
		i++

	}
	return true
}

func main() {
	var aN int32

	fmt.Printf("Введите, пожалуйста N ")

	_, err := fmt.Scanln(&aN)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		os.Exit(1)
	}

	//слайс - список простых чисел
	var numArr []int32
	var i int32 = 2

	for i <= aN {
		if isPrime(i) == true {
			//fmt.Printf("Число %d простое\n", i)
			numArr = append(numArr, i)
		}
		i++
	}

	for i := range numArr {
		fmt.Printf("Число %d простое\n", numArr[i])
	}

}
