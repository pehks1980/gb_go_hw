package main

import (
	"fmt"
)

/*  С клавиатуры вводится трехзначное число.
    Выведите цифры, соответствующие количество
	сотен, десятков и единиц  в этом числе.
*/

func main() {
	var aNumber uint16

	fmt.Printf("Введите, пожалуйста число:")

	_, err := fmt.Scanf("%d", &aNumber)
	if err != nil {
		// error here
		fmt.Printf("Error enter")
		return
	}

	fmt.Printf("Число сотен: %d, десятков: %d, единиц: %d\n",
		(aNumber%1000)/100, (aNumber%100)/10, aNumber%10)
}
