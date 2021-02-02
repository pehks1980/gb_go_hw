package main

import (
	"fmt"
	"log"
)

/*
Задачка:
    1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
    2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
*/

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

// вычисляет число фиббоббанначчи
func getFibbo(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return getFibbo(n-1) + getFibbo(n-2)
}

//вычисляет число фиббоббанначчи c кешированием предварит вычислений
func getFibboMap(n int64, cache *map[int]int64) int64 {
	if val, ok := (*cache)[int(n)]; ok {
		fmt.Println("!got cache for n=", n)
		return val
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return getFibboMap(n-1, cache) + getFibboMap(n-2, cache)
}

func main() {
	n := getDigit("Введите N: (40 для наглядности): ")

	fmt.Println("Вычисление обычной рекурсией:")
	for i := 0; i <= int(n); i++ {
		fmt.Printf("Result n = %d значение = %d\n", i, getFibbo(int64(i)))
	}

	mapResult := make(map[int]int64, n)
	fmt.Println("Вычисление рекурсией и кешированием:")
	for i := 0; i <= int(n); i++ {
		mapResult[i] = getFibboMap(int64(i), &mapResult)
		fmt.Printf("Result n = %d значение = %d\n", i, mapResult[i])
	}

	fmt.Println("Содержимое cache:")
	for k, v := range mapResult {
		fmt.Printf("Result n = %d значение = %d\n", k, v)
	}

}
