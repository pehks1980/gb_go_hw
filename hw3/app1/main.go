package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

// функция берет число с клавиатуры, проверяет данные на валидность
func getDigit(prompt string) int32 {
	var a int32
	fmt.Println(prompt)
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatal("ошибко %v", err)
	}
	return a
}

//вычисляет факториал рекурсией
func getFact(n int64) int64 {
	if n == 0 {
		return 1
	}
	return getFact(n-1) * n
}

/*
калькулятор: больше операций и валидации данных
*/
func main() {
	var a, b int32
	var op string
	var res int32

label:

	opNumber := getDigit("Введите номер операции:  (0 - выход) 1-'+' 2-'+' 3-'*' 4-'/' 5-'square root' 6-'factorial'")
	switch opNumber {
	case 0:
		fmt.Println("До свидания!")
		os.Exit(0)
	case 1:
		a = getDigit("Введите число A:")
		b = getDigit("Введите число Б:")
		op = "+"
	case 2:
		a = getDigit("Введите число A:")
		b = getDigit("Введите число Б:")
		op = "-"
	case 3:
		a = getDigit("Введите число A:")
		b = getDigit("Введите число Б:")
		op = "*"
	case 4:
		a = getDigit("Введите число A:")
		b = getDigit("Введите число Б:")
		op = "/"
	case 5:
		a = getDigit("Введите число A:")
		op = "sqrt"
	case 6:
		a = getDigit("Введите число A:")
		op = "fact"
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	switch op {
	case "+":
		res = a + b
		fmt.Printf("Результат выполнения операции: %d\n", res)
	case "-":
		res = a - b
		fmt.Printf("Результат выполнения операции: %d\n", res)
	case "*":
		res = a * b
		fmt.Printf("Результат выполнения операции: %d\n", res)
	case "/":
		if b != 0 {
			resFl := float32(a) / float32(b)
			fmt.Printf("Результат выполнения операции: %.4f\n", resFl)
		} else {
			fmt.Println("На 0 делить нельзя!")
			//os.Exit(1)
		}
	case "sqrt":
		if a > 0 {
			resFl := math.Sqrt(float64(a))
			fmt.Printf("Результат выполнения операции: %.4f\n", resFl)
		} else {
			fmt.Println("Число должно быть больше 0")
			//os.Exit(1)
		}
	case "fact":
		if a > 0 && a < 26 {
			res := getFact( int64(a) )
			fmt.Printf("Результат выполнения операции: %d\n", res)
		} else {
			fmt.Println("Число должно быть в диапазоне 1-25 ")
			//os.Exit(1)
		}

	}

	goto label
}
