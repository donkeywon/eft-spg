package eft

import (
	"github.com/bytedance/sonic/ast"
	"github.com/donkeywon/eft-spg/util"
)

func getWeightedInventoryItem(equipPool *ast.Node) (string, int) {
	var itemKeys []string
	var weights []int

	equipPool.ForEach(func(path ast.Sequence, node *ast.Node) bool {
		itemKeys = append(itemKeys, *path.Key)
		w, _ := node.Int64()
		weights = append(weights, int(w))

		return true
	})

	return weightedRandom(itemKeys, weights)
}

func weightedRandom(items []string, weights []int) (string, int) {
	if len(items) != len(weights) || len(items) == 0 {
		return "", 0
	}

	cumulativeWeights := make([]int, len(weights), len(weights))
	cumulativeWeights[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		cumulativeWeights[i] = weights[i] + cumulativeWeights[i-1]
	}

	maxCumulativeWeight := cumulativeWeights[len(cumulativeWeights)-1]
	randNum := maxCumulativeWeight * util.RandInt(0, 100) / 100

	for i := 0; i < len(items); i++ {
		if cumulativeWeights[i] >= randNum {
			return items[i], i
		}
	}

	return "", 0
}
