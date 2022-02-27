package util

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandInt(1, 0))
	}
}

func TestChoose(t *testing.T) {
	lastNames := []interface{}{}
	lastName := RandChoose(lastNames)
	if lastName == nil {
		lastName = ""
	}

	fmt.Println(lastName)
}
