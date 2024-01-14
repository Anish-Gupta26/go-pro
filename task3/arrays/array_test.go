package sum

import (
	"reflect"
	"testing"
)

//Total sum
func TestSum(t *testing.T){
	num := [] int{1,2,3,4,5}
	want := 15
	got := SumOfArray(num)

	if(want!=got){
		t.Errorf("Error want %d, got %d, for numbers %v",want,got,num)
	}else{
		t.Log("Seccess")
	}
}

//Sum of multiple slices
func TestSumMultiple(t *testing.T){
	got := SumMultiple([]int{1,2,3},[]int{1,2,3})
	want := []int{6,6}

	if !reflect.DeepEqual(got,want){
		t.Errorf("Error want %d, got %d",want,got)
	}else{
		t.Log("Seccess")
	}

}
