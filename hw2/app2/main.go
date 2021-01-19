package main

import (
	"fmt"
	"github.com/pehks1980/gb_go_hw/hw2/app2/circle"
	"os"
)

/*Напишите программу, вычисляющую диаметр(наверное, скорее радиус) и длину окружности по заданной площади круга.
Площадь круга должна вводиться пользователем с клавиатуры.*/
func main() {
	var circle_area float64

	fmt.Printf("Введите, пожалуйста площадь круга: ")

	_, err := fmt.Scanf("%g", &circle_area)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		os.Exit(128)
	}
	// compute radius & circle_length in one func
	diameter, circle_len := circle.CircleParams(circle_area)

	fmt.Printf("Диаметр окружности равен: %.2f\n", diameter)
	fmt.Printf("Длина окружности равна: %.2f\n", circle_len)
}
