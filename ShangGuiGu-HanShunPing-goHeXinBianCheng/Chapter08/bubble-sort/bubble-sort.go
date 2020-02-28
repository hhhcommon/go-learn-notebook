package main

import "fmt"

func BubbleSort(arr *[6]int) {
	tmp := 0
	array := *arr
	fmt.Printf("Befor sort: %+v\n", array)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				tmp = array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}

	}
	arr = &array
	fmt.Printf("Success Sorted!  Sorted array is :%+v\n", *arr)
}

func main() {
	arr := [6]int{45, 3, 56, 22, 67, 889}
	BubbleSort(&arr)
}
