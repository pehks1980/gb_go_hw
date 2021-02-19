package fibbo_bm_test

import (
	"testing"
)

/*
Задачка:
    1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
    2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.
*/

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
func getFibboMap(n uint64, cache map[uint64]uint64) uint64 {
	if val, ok := cache[n]; ok {
		//fmt.Println("!got cache for n=", n)
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

/// функция из методички
func F(n uint8) uint64 {
	var f1, f2 uint64
	f2 = 1

	var i uint8
	for ; i < n-1; i++ {
		sum := f1 + f2
		f1 = f2
		f2 = sum
	}

	return f2
}

//result is result of test - must be
var result uint64

//////////////////////тест 1 бенчмарк вычисления числа фибобаначи рекурсией 30 чисел
func BenchmarkFibboRecursive(b *testing.B) {
	var res uint64

	for i := 0; i < b.N; i++ {
		res = getFibbo(30)
	}

	result = res
}

//////////////////////тест 2 бенчмарк вычисления числа фибобаначи рекурсией+кэш мапа 30 чисел
func BenchmarkFibboRecursiveCache(b *testing.B) {
	var res uint64

	mapResult := make(map[uint64]uint64, 30)

	for i := 0; i < b.N; i++ {
		mapResult[30] = getFibboMap(30, mapResult)
		res = mapResult[30]
	}

	result = res
}

//////////////////////тест 3 бенчмарк вычисления числа фибобаначи суммированием из методички 30 чисел
func BenchmarkFibboSum(b *testing.B) {
	var res uint64

	for i := 0; i < b.N; i++ {
		res = F(30)
	}

	result = res
}


