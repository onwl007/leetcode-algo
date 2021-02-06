# 二叉树

## 二叉树前中后序遍历模板

```go
func traversal(root *TreeNode) {
  // 前序遍历
  traversal(root.Left)
  // 中序遍历
  traversal(root.Right)
  // 后序遍历
}
```

## 前序遍历

按照 `根节点-左子树-右子树`的访问顺序访问一颗二叉树

- 递归解法

```go
func preorderTraversal(root *TreeNode) (vals []int) {
    preorder(&vals, root)
    return
}

func preorder(vals *[]int, root *TreeNode) {
    if root == nil {
        return
    }
    *vals = append(*vals, root.Val)
    preorder(vals, root.Left)
    preorder(vals, root.Right)
}
```

- 迭代解法

```go
func preorderIter(root *TreeNode) []int {
    vals := []int{}
    stack := []*TreeNode{}
    for len(stack) > 0 || root != nil {
        for root != nil {
            vals = append(vals, root.Val)
            stack = append(stack, root.Right)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
    }
    return vals
}
```

## 中序遍历

按照 `左子树-根节点-右子树`的访问顺序访问二叉树

- 递归解法

```go
func inorderTraversal(root *TreeNode) (vals []int) {
    inorder(&vals, root)
    return
}

func inorder(vals *[]int, root *TreeNode) {
    if root == nil {
        return
    }
    inorder(vals, root.Left)
    *vals = append(*vals, root.Val)
    inorder(vals, root.Right)
}
```

- 迭代解法

```go
func inorderIter(root *TreeNode) []int {
    vals := []int{}
    stack := []*TreeNode{}
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        vals = append(vals, root.Val)
        root = root.Right
    }
    return vals
}
```

## 后序遍历

按照 `左子树-右子树-根节点`的访问顺序访问二叉树

- 递归解法

```go
func postorderTraversal(root *TreeNode) (vals []int) {
    postorder(&vals, root)
    return
}

func postorder(vals *[]int, root *TreeNode) {
    if root == nil {
        return
    }
    postorder(vals, root.Left)
    postorder(vals, root.Right)
    *vals = append(*vals, root.Val)
}
```

## 层序遍历

按照层次，一层一层的遍历二叉树。
利用队列，将根节点入队，判断根节点左子树和右子树不为空，依次入队。将根节点元素添加的数组中，然后出队。

```go
func levelOrder(root *TreeNode) (vals []int) {
    if root == nil {
        return
    }
    queue := []*TreeNode{root}
    for len(queue) != 0 {
        node := queue[0]
        vals = append(vals, node.Val)
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
        queue = queue[1:]
    }
    return
}
```
