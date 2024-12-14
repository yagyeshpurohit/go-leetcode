package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func insertLevelOrder(arr []interface{}, root *TreeNode, i int) *TreeNode {

	if i >= len(arr) || arr[i] == nil {
		return nil
	}
	//create node
	root = newNode(arr[i].(int))
	//insert left child
	root.Left = insertLevelOrder(arr, root.Left, 2*i+1)
	//insert right child
	root.Right = insertLevelOrder(arr, root.Right, 2*i+2)

	return root
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrder(root.Right)

}

//func levelOrder(root *TreeNode) {
//	if root != nil {
//		nodeQueue := []*TreeNode{root}
//		for len(nodeQueue) > 0 {
//			currNode := nodeQueue[0]
//			if currNode == nil {
//				fmt.Printf("null ")
//			} else {
//				fmt.Printf("%d ", currNode.Val)
//			}
//			nodeQueue = nodeQueue[1:]
//			if currNode != nil {
//				nodeQueue = append(nodeQueue, currNode.Left)
//				nodeQueue = append(nodeQueue, currNode.Right)
//			}
//			//if currNode.Left != nil {
//			//	nodeQueue = append(nodeQueue, currNode.Left)
//			//}
//			//if currNode.Right != nil {
//			//	nodeQueue = append(nodeQueue, currNode.Right)
//			//}
//		}
//	}
//
//}

func createTree(arr []interface{}) *TreeNode {
	var root *TreeNode
	root = insertLevelOrder(arr, root, 0)
	return root
}

//https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

//https://leetcode.com/problems/binary-tree-level-order-traversal/description/
func levelOrder(root *TreeNode) [][]int {
	var levelOrderSeq [][]int
	if root == nil {
		return levelOrderSeq
	}
	queue := []*TreeNode{root}
	//while queue is not empty
	for len(queue) > 0 {
		currQueueLen := len(queue) // length of queue at any given level
		var sameLevelNodes []int   // for storing all the nodes at a given level
		//for a particular level
		for i := 1; i <= currQueueLen; i++ {
			currNode := queue[0] //node visited (front of queue)
			queue = queue[1:]    //remove the visited node from the front of the queue
			sameLevelNodes = append(sameLevelNodes, currNode.Val)
			//no null children to be stored in the queue
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}
		levelOrderSeq = append(levelOrderSeq, sameLevelNodes)
	}
	return levelOrderSeq
}

//https://leetcode.com/problems/path-sum-ii/
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var pathNodes []int
	pathSumHelper(root, targetSum, pathNodes, &result)
	return result
}

func pathSumHelper(node *TreeNode, targetSum int, pathNodes []int, result *[][]int) {
	if node == nil {
		return
	}
	targetSum -= node.Val
	pathNodes = append(pathNodes, node.Val)
	// if leaf node
	if node.Left == nil && node.Right == nil {
		if targetSum == 0 {
			*result = append(*result, pathNodes)
			fmt.Println(result)
			return
		}
	}
	pathSumHelper(node.Left, targetSum, pathNodes, result)
	pathSumHelper(node.Right, targetSum, pathNodes, result)

	//remove the traversed node before returning
	pathNodes = pathNodes[:len(pathNodes)-1]
}
