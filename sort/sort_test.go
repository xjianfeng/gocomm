package sort

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	fmt.Println("=======================")
	m := SortMap{
		"1": 2,
		"5": 1,
		"3": "S",
		"2": "1"}
	s := m.SortStrKeyUrlEncoded(false)
	fmt.Printf("%v\n", s)
}
