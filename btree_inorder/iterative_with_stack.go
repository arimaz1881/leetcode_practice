package main

import "fmt"


type Stack struct {
	items []*TreeNode
}

func (s *Stack) push(node *TreeNode) {
	s.items = append(s.items, node)
}

func (s *Stack) pop() *TreeNode {
	if len(s.items) == 0 {
		return nil
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := Stack{}
	curr := root

	for curr != nil || len(stack.items) > 0 {
		if curr != nil {
			stack.push(curr)
			curr = curr.Left
		} else {
			curr = stack.pop()
			result = append(result, curr.Val)
			curr = curr.Right
		}
	}
	return result
}

func main() {
	binaryTree := BinaryTree{}

	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, val := range values {
		binaryTree.Insert(val)
	}

	fmt.Println("In-Order Traversal: ", inorderTraversal(binaryTree.Root))
}


// ------------------------------------------ BTree implementation ---------------------------------------------
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (bt *BinaryTree) Insert(val int) {
	bt.Root = insertNode(bt.Root, val)
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}

	if val < root.Val {
		root.Left = insertNode(root.Left, val)
	} else {
		root.Right = insertNode(root.Right, val)
	}

	return root
}

