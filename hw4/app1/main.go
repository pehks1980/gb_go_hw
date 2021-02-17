package main

import "fmt"

//функция сортирует список int
func mySortIns(a []int) []int {
	// движемся от 2 элемента в массиве до конца массива
	for i := 1; i < len(a); i++ {
		// берем сравнваемое значение из массива
		tmp := a[i]
		idx := i
		// тут (idx - 0) кусок в который будет вставляться tmp
		// в цикле идет раздвижение (на одну позицию вправо) в этом куске до того места в который
		// будет вставлена tmp по ранжиру
		for idx > 0 && a[idx-1] > tmp {
			a[idx] = a[idx-1]
			idx--
		}
		a[idx] = tmp
		// в одной итерации цикла for кусок массива (idx - 0) - отсортирован
		// чем длиннее массив тем будет дольше выполняться внутренний цикл для вставки по ранжиру
	}
	return a
}

func main() {

	s1 := []int{5, 2, 6, 3, 1, 4} // unsorted
	//sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)

	s2 := mySortIns(s1)

	fmt.Println(s2)

}


