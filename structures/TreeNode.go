package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

// 将数组转成二叉树
func Ints2TreeNode(arr []int) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && arr[i] != NULL {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && arr[i] != NULL {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Ints2Node(arr []int) *Node {
	n := len(arr)
	if n == 0 {
		return nil
	}

	root := &Node{Val: arr[0]}

	queue := make([]*Node, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && arr[i] != NULL {
			node.Left = &Node{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && arr[i] != NULL {
			node.Right = &Node{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func Preorder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	*vals = append(*vals, root.Val)
	Preorder(root.Left, vals)
	Preorder(root.Right, vals)
}
