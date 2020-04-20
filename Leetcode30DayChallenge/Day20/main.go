/*
Return the root node of a binary search tree that matches
the given preorder traversal.

(Recall that a binary search tree is a binary tree where
for every node, any descendant of node.left has a
value < node.val, and any descendant of node.right has a
value > node.val.  Also recall that a preorder traversal
displays the value of the node first, then traverses
node.left, then traverses node.right.)

Example 1:
Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]

Note:

1 <= preorder.length <= 100
The values of preorder are distinct.
*/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := []int{8, 5, 1, 7, 10, 12}
	fmt.Println(bstFromPreorder(test))
}

func bstFromPreorder(preorder []int) *TreeNode {
	v := preorder[0]
	node := &TreeNode{v, nil, nil}
	index, left := bstFromPreorderRec(preorder, 1, node)
	node.Left = left
	_, right := bstFromPreorderRec(preorder, index, nil)
	node.Right = right
	return node
}
func bstFromPreorderRec(preorder []int, start int, ancestor *TreeNode) (int, *TreeNode) {
	if start >= len(preorder) {
		return start, nil
	}
	v := preorder[start]
	if ancestor != nil && v > ancestor.Val {
		return start, nil
	}
	node := &TreeNode{v, nil, nil}
	index, left := bstFromPreorderRec(preorder, start+1, node)
	node.Left = left
	next, right := bstFromPreorderRec(preorder, index, ancestor)
	node.Right = right
	return next, node
}
