/*
Given a binary tree, you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes
in a tree. This path may or may not pass through the root.

Example:
Given a binary tree

          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].
*/
package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 1
	height(root)
	return ans - 1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	L := height(node.Left)
	R := height(node.Right)
	ans = max(ans, L+R+1)
	return max(L, R) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
