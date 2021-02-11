package main

import (
	"fmt"
)

/* 	Напишите программу для вычисления площади прямоугольника.
Длины сторон прямоугольника должны вводиться пользователем с клавиатуры
*/

// Функция вычисляет площадь прямоугольника
func area(length float64, height float64) float64 {
	return length * height
}

func main() {
	var length, height float32

	fmt.Printf("Введите, пожалуйста длину: ")

	_, err := fmt.Scanf("%g", &length)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		return
	}

	fmt.Printf("Введите, пожалуйста ширину: ")
	_, err = fmt.Scanln(&height)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		return
	}

	result := area(float64(length), float64(height))

	fmt.Printf("Площадь прямоугольника равна: %.2f\n", result)

}
