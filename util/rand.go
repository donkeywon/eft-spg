package util

import (
	"github.com/bytedance/sonic/ast"
	"math/rand"
	"time"
)

func RandInt(min int, max int) int {
	if max <= min {
		return min
	}

	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func RandIntNode(min *ast.Node, max *ast.Node) int {
	if min == nil || max == nil || !min.Exists() || !max.Exists() {
		return -1
	}

	min1, err := min.Int64()
	if err != nil {
		return -1
	}

	max1, err := max.Int64()
	if err != nil {
		return -1
	}

	return RandInt(int(min1), int(max1))
}

func RandChoose(arr []interface{}) interface{} {
	if len(arr) == 0 {
		return nil
	}

	return arr[RandInt(0, len(arr))]
}

func RandChooseNode(arrNode *ast.Node) *ast.Node {
	arr, err := arrNode.ArrayUseNode()
	if err != nil {
		return nil
	}

	if len(arr) == 0 {
		return nil
	}

	return &arr[RandInt(0, len(arr))]
}
