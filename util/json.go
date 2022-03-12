package util

import "github.com/bytedance/sonic/ast"

func JsonNodeKeys(n *ast.Node) []string {
	nl, _ := n.Len()
	ks := make([]string, 0, nl)
	n.ForEach(func(path ast.Sequence, node *ast.Node) bool {
		ks = append(ks, *path.Key)
		return true
	})
	return ks
}
