// doubly_linked_list.go

package main

import "fmt"
type Node struct {
    data int
    prev *Node
    next *Node
}
type DoublyLinkedList struct {
    head *Node
    tail *Node
}

func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
    newNode := &Node{data: data}
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
        return
    }

    newNode.next = dll.head
    dll.head.prev = newNode
    dll.head = newNode
}

func (dll *DoublyLinkedList) InsertAtEnd(data int) {
    newNode := &Node{data: data}
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
        return
    }

    newNode.prev = dll.tail
    dll.tail.next = newNode
    dll.tail = newNode
}

func (dll *DoublyLinkedList) InsertAtPosition(data, position int) {
    newNode := &Node{data: data}
    if dll.head == nil || position == 0 {
        dll.InsertAtBeginning(data)
        return
    }

    currentNode := dll.head
    for i := 0; i < position-1 && currentNode != nil; i++ {
        currentNode = currentNode.next
    }

    if currentNode == nil {
        return
    }

    newNode.prev = currentNode
    newNode.next = currentNode.next
    currentNode.next.prev = newNode
    currentNode.next = newNode
}

func (dll *DoublyLinkedList) DeleteFromBeginning() {
    if dll.head == nil {
        return
    }

    if dll.head == dll.tail {
        dll.head = nil
        dll.tail = nil
        return
    }

    dll.head = dll.head.next
    dll.head.prev = nil
}

func (dll *DoublyLinkedList) DeleteFromEnd() {
    if dll.head == nil {
        return
    }

    if dll.head == dll.tail {
        dll.head = nil
        dll.tail = nil
        return
    }

    dll.tail = dll.tail.prev
    dll.tail.next = nil
}

func (dll *DoublyLinkedList) DeleteFromPosition(position int) {
    if dll.head == nil {
        return
    }

    if position == 0 {
        dll.DeleteFromBeginning()
        return
    }

    currentNode := dll.head
    for i := 0; i < position && currentNode != nil; i++ {
        currentNode = currentNode.next
    }

    if currentNode == nil {
        return
    }

    if currentNode == dll.tail {
        dll.tail = currentNode.prev
        dll.tail.next = nil
    } else {
        currentNode.prev.next = currentNode.next
        currentNode.next.prev = currentNode.prev
    }
}

func (dll *DoublyLinkedList) DisplayForward() {
    currentNode := dll.head
    for currentNode != nil {
        fmt.Printf("%d <-> ", currentNode.data)
        currentNode = currentNode.next
    }
    fmt.Println("nil")
}

func (dll *DoublyLinkedList) DisplayBackward() {
    currentNode := dll.tail
    for currentNode != nil {
        fmt.Printf("%d <-> ", currentNode.data)
        currentNode = currentNode.prev
    }
    fmt.Println("nil")
}

func main() {
    doublyLinkedList := &DoublyLinkedList{}

    doublyLinkedList.InsertAtEnd(10)
    doublyLinkedList.InsertAtBeginning(5)
    doublyLinkedList.InsertAtEnd(20)
    doublyLinkedList.InsertAtPosition(15, 2)

    fmt.Println("Original Doubly Linked List (Forward Order):")
    doublyLinkedList.DisplayForward()

    fmt.Println("Original Doubly Linked List (Backward Order):")
    doublyLinkedList.DisplayBackward()

    doublyLinkedList.DeleteFromBeginning()
    doublyLinkedList.DeleteFromEnd()
    doublyLinkedList.DeleteFromPosition(1)

    fmt.Println("Updated Doubly Linked List (Forward Order):")
    doublyLinkedList.DisplayForward()

    fmt.Println("Updated Doubly Linked List (Backward Order):")
    doublyLinkedList.DisplayBackward()
}
