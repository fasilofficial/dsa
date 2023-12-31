// linked_list.go

package main

import "fmt"
type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

// InsertAtBeginning inserts a new node at the beginning of the linked list.
func (ll *LinkedList) InsertAtBeginning(data int) {
    newNode := &Node{data: data}
    newNode.next = ll.head
    ll.head = newNode
}

// InsertAtEnd inserts a new node at the end of the linked list.
func (ll *LinkedList) InsertAtEnd(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
        return
    }

    lastNode := ll.head
    for lastNode.next != nil {
        lastNode = lastNode.next
    }
    lastNode.next = newNode
}

// InsertAtPosition inserts a new node at a specific position in the linked list.
func (ll *LinkedList) InsertAtPosition(data, position int) {
    newNode := &Node{data: data}
    if position == 0 {
        newNode.next = ll.head
        ll.head = newNode
        return
    }

    currentNode := ll.head
    for i := 0; i < position-1 && currentNode != nil; i++ {
        currentNode = currentNode.next
    }

    if currentNode == nil {
        return
    }

    newNode.next = currentNode.next
    currentNode.next = newNode
}

// DeleteFromBeginning deletes a node from the beginning of the linked list.
func (ll *LinkedList) DeleteFromBeginning() {
    if ll.head == nil {
        return
    }

    ll.head = ll.head.next
}

// DeleteFromEnd deletes a node from the end of the linked list.
func (ll *LinkedList) DeleteFromEnd() {
    if ll.head == nil || ll.head.next == nil {
        ll.head = nil
        return
    }

    currentNode := ll.head
    for currentNode.next.next != nil {
        currentNode = currentNode.next
    }
    currentNode.next = nil
}

// DeleteFromPosition deletes a node from a specific position in the linked list.
func (ll *LinkedList) DeleteFromPosition(position int) {
    if ll.head == nil {
        return
    }

    if position == 0 {
        ll.head = ll.head.next
        return
    }

    currentNode := ll.head
    for i := 0; i < position-1 && currentNode.next != nil; i++ {
        currentNode = currentNode.next
    }

    if currentNode.next == nil {
        return
    }

    currentNode.next = currentNode.next.next
}

// Display prints the elements of the linked list.
func (ll *LinkedList) Display() {
    currentNode := ll.head
    for currentNode != nil {
        fmt.Printf("%d -> ", currentNode.data)
        currentNode = currentNode.next
    }
    fmt.Println("nil")
}

func main() {
    // Example usage
    linkedList := &LinkedList{}

    linkedList.InsertAtEnd(10)
    linkedList.InsertAtBeginning(5)
    linkedList.InsertAtEnd(20)
    linkedList.InsertAtPosition(15, 2)

    fmt.Println("Original Linked List:")
    linkedList.Display()

    linkedList.DeleteFromBeginning()
    linkedList.DeleteFromEnd()
    linkedList.DeleteFromPosition(1)

    fmt.Println("Updated Linked List:")
    linkedList.Display()
}
