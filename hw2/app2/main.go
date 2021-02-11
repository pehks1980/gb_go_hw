package main

import (
	"fmt"
	"github.com/pehks1980/gb_go_hw/hw2/app2/circle"
)

/*Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
Площадь круга должна вводиться пользователем с клавиатуры.*/

func main() {
	var circle_area float64

	fmt.Printf("Введите, пожалуйста площадь круга: ")

	_, err := fmt.Scanf("%g", &circle_area)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		return
	}

	if circle_area > 0 {
		// вычисление диаметра и длины окружности по заданной площади круга
		diameter, circle_len := circle.CircleParams(circle_area)

		fmt.Printf("Диаметр окружности равен: %.2f\n", diameter)
		fmt.Printf("Длина окружности равна: %.2f\n", circle_len)
	} else {
		fmt.Printf("Площадь круга должна быть >0 ")
	}

}
