package sorts


import (
	"sort"
	"testing"
	"fmt"
)


func TestArraySort(T *testing.T) {
	fruits := SortInts([]int{2, 1, 3})
	sort.Sort(fruits)
	fmt.Println(fruits)
	if fruits[0] != 1 {
		T.Error("sort failed")
	}


	intSice := []int{2, 1, 3}
	sort.Ints(intSice)
	if intSice[0] != 1 {
		T.Error("sort int slice failed")
	}
}