package util

import (
	"fmt"
	"testing"
)

func TestArrayContains(t *testing.T) {
	a := []interface{}{"abc", "def", "qwe", "asd", 1, 2, 3}

	fmt.Println(ArrayContains(a, 1))
}
