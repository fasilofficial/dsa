// bst.go

package main

import "fmt"

type BSTNode struct {
    data  int
    left  *BSTNode
    right *BSTNode
}

type BST struct {
    root *BSTNode
}

func NewBST() *BST {
    return &BST{}
}

func (bst *BST) Insert(data int) {
    bst.root = insertNode(bst.root, data)
}

func insertNode(root *BSTNode, data int) *BSTNode {
    if root == nil {
        return &BSTNode{data: data}
    }

    if data < root.data {
        root.left = insertNode(root.left, data)
    } else if data > root.data {
        root.right = insertNode(root.right, data)
    }

    return root
}

func (bst *BST) Display() {
    fmt.Print("DFS Inorder Traversal: ")
    inorderTraversal(bst.root)
    fmt.Println()

    fmt.Print("DFS Preorder Traversal: ")
    preorderTraversal(bst.root)
    fmt.Println()

    fmt.Print("DFS Postorder Traversal: ")
    postorderTraversal(bst.root)
    fmt.Println()

    fmt.Print("BFS Level Order Traversal: ")
    levelOrderTraversal(bst.root)
    fmt.Println()
}

func inorderTraversal(root *BSTNode) {
    if root != nil {
        inorderTraversal(root.left)
        fmt.Printf("%d ", root.data)
        inorderTraversal(root.right)
    }
}

func preorderTraversal(root *BSTNode) {
    if root != nil {
        fmt.Printf("%d ", root.data)
        preorderTraversal(root.left)
        preorderTraversal(root.right)
    }
}

func postorderTraversal(root *BSTNode) {
    if root != nil {
        postorderTraversal(root.left)
        postorderTraversal(root.right)
        fmt.Printf("%d ", root.data)
    }
}

func levelOrderTraversal(root *BSTNode) {
    if root == nil {
        return
    }

    queue := []*BSTNode{root}
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
    bst := NewBST()
    bst.Insert(4)
    bst.Insert(2)
    bst.Insert(6)
    bst.Insert(1)
    bst.Insert(3)
    bst.Insert(5)
    bst.Insert(7)

    fmt.Println("BST:")
    bst.Display()
}
