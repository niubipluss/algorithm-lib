// FILE: Binary-Search-Tree-to-Greater-Sum-Tree/BinarySearchTreetoGreaterSumTree.go
// AUTH: ydhmkpsryk321@gmail.com
// DES: 2024-02-01 10:36:36
package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	sum := 0
	dfs1038(root, &sum)
	return root
}

func dfs1038(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs1038(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	dfs1038(root.Left, sum)
}
