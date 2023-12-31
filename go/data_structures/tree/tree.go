// tree.go

package main

import "fmt"

type TreeNode struct {
    data  int
    left  *TreeNode
    right *TreeNode
}

type Tree struct {
    root *TreeNode
}

func NewTree() *Tree {
    return &Tree{}
}

func (t *Tree) Insert(data int) {
    newNode := &TreeNode{data: data}
    if t.root == nil {
        t.root = newNode
        return
    }

    queue := []*TreeNode{t.root}
    for len(queue) > 0 {
        currentNode := queue[0]
        queue = queue[1:]

        if currentNode.left == nil {
            currentNode.left = newNode
            return
        } else {
            queue = append(queue, currentNode.left)
        }

        if currentNode.right == nil {
            currentNode.right = newNode
            return
        } else {
            queue = append(queue, currentNode.right)
        }
    }
}

func (t *Tree) Display() {
    if t.root == nil {
        fmt.Println("Tree is empty.")
        return
    }

    fmt.Print("DFS Inorder Traversal: ")
    t.inorderTraversal(t.root)
    fmt.Println()

    fmt.Print("DFS Preorder Traversal: ")
    t.preorderTraversal(t.root)
    fmt.Println()

    fmt.Print("DFS Postorder Traversal: ")
    t.postorderTraversal(t.root)
    fmt.Println()

    fmt.Print("BFS Level Order Traversal: ")
    t.levelOrderTraversal()
    fmt.Println()
}

func (t *Tree) inorderTraversal(root *TreeNode) {
    if root != nil {
        t.inorderTraversal(root.left)
        fmt.Printf("%d ", root.data)
        t.inorderTraversal(root.right)
    }
}

func (t *Tree) preorderTraversal(root *TreeNode) {
    if root != nil {
        fmt.Printf("%d ", root.data)
        t.preorderTraversal(root.left)
        t.preorderTraversal(root.right)
    }
}

func (t *Tree) postorderTraversal(root *TreeNode) {
    if root != nil {
        t.postorderTraversal(root.left)
        t.postorderTraversal(root.right)
        fmt.Printf("%d ", root.data)
    }
}

func (t *Tree) levelOrderTraversal() {
    queue := []*TreeNode{t.root}
    for len(queue) > 0 {
        currentNode := queue[0]
        queue = queue[1:]

        fmt.Printf("%d ", currentNode.data)

        if currentNode.left != nil {
            queue = append(queue, currentNode.left)
        }

        if currentNode.right != nil {
            queue = append(queue, currentNode.right)
        }
    }
}

func main() {
    tree := NewTree()
    tree.Insert(1)
    tree.Insert(2)
    tree.Insert(3)
    tree.Insert(4)

    fmt.Println("Binary Tree:")
    tree.Display()
}
