package sum

import "fmt"

func SumOfArray(numbers[] int) int{
	total := 0
	size := len(numbers)
	for i:=0;i<size;i++{
		total+=numbers[i]
	}
	return total
}

func SumMultiple(slices...[]int) []int{
	// numOfSlice := len(slices)
	fmt.Println(len(slices))
	// newSlice := make([]int,numOfSlice)
	newSlice := []int{}
	for _,num := range slices{
		fmt.Println(num)
		newSlice=append(newSlice, SumOfArray(num))
	}
	return newSlice
}