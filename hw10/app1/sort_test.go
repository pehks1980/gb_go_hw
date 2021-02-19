package test_sort

import (
	"reflect"
	"testing"
)


//функция сортирует список int
func mySortIns(a []int) {
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
	//return a
}

////////////////////////////тест 1 - табличное тестирование
func TestSplit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "insert sort positive", input: []int{5, 2, 6, 3, 1, 4}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "insert sort negative", input: []int{5, -2, 6, -3, -1, 4}, want: []int{-3, -2, -1, 4, 5, 6}},
		{name: "insert sort +,- 0s", input: []int{-5, 2, 0, 3, -1, 4}, want: []int{-5, -1, 0, 2, 3, 4}},
	}

	for _, tc := range tests {
		mySortIns(tc.input)
		if !reflect.DeepEqual(tc.want, tc.input) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, tc.input)
		}
	}
}
