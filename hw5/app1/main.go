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
func getDigit(prompt string) (uint64, error) {
	var a uint64
	fmt.Println(prompt)
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println("Ошибка:", err)
		return 0, err
	}
	return a, nil
}

// вычисляет число фиббоббанначчи
func getFibbo(n uint64) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return getFibbo(n-1) + getFibbo(n-2)
}

//вычисляет число фиббоббанначчи c кешированием предварит вычислений
func getFibboMap(n uint64, cache *map[uint64]uint64) uint64 {
	if val, ok := (*cache)[n]; ok {
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
	n, err := getDigit("Введите N: (40 для наглядности): ")

	if err != nil {
		fmt.Println("Oшибка ввода.")
		return
	}

	fmt.Println("Вычисление обычной рекурсией:")
	for i := uint64(0); i <= n; i++ {
		fmt.Printf("Result n = %d значение = %d\n", i, getFibbo(i))
	}

	mapResult := make(map[uint64]uint64, n)
	fmt.Println("Вычисление рекурсией и кешированием:")
	for i := uint64(0); i <= n; i++ {
		mapResult[i] = getFibboMap(i, &mapResult)
		fmt.Printf("Result n = %d значение = %d\n", i, mapResult[i])
	}

	fmt.Println("Содержимое cache:")
	for k, v := range mapResult {
		fmt.Printf("Result n = %d значение = %d\n", k, v)
	}

}
