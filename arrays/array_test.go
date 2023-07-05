package arrays

import (
	"testing"
	"fmt"
)

func TestArray(t *testing.T) {
    
	arr := CreateArray()

	//test for empty array

	length := arr.Length()

	if length != 0 {
		t.Errorf("Expected length 0, got %d", length)
	}

    //test adding value
	err := arr.AddValue(0, 40)

	if err != nil {
		t.Errorf("")
	}
    



}