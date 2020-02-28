package main

import "fmt"

func BinarySearch(arr *[5]int, startIndex int, endIndex int, resultValue int) {
	
	array := *arr
	if endIndex < startIndex {
		fmt.Println("没找到")
		return
	}

	middleIndex := (startIndex + endIndex) / 2

	if array[middleIndex] > resultValue {
		BinarySearch(arr, 0, middleIndex-1, resultValue)
	} else if array[middleIndex] < resultValue {
		BinarySearch(arr, middleIndex+1, endIndex, resultValue)
	} else {
		fmt.Println("Success! Result`s index  is ", middleIndex)
		return
	}

}

func main() {
	// 有序数列，从大到小
	arr := [5]int{2, 34, 67, 888, 1024}
	BinarySearch(&arr, 0, 4, 1024)
}
