// hash_table.go

package main

import (
	"fmt"
	"hash/fnv"
)

const arraySize = 10
type Node struct {
	key   string
	value int
	next  *Node
}

type Bucket struct {
	head *Node
}
type HashTable struct {
	array [arraySize]Bucket
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (ht *HashTable) insert(key string, value int) {
	index := hash(key)
	bucket := &ht.array[index]

	newNode := &Node{
		key:   key,
		value: value,
		next:  bucket.head,
	}

	bucket.head = newNode
}

func (ht *HashTable) get(key string) (int, bool) {
	index := hash(key)
	bucket := &ht.array[index]
	current := bucket.head

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}

	return 0, false
}

func hash(key string) int {
	hasher := fnv.New32()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % arraySize
}

func main() {
	hashTable := NewHashTable()

	hashTable.insert("apple", 5)
	hashTable.insert("banana", 10)
	hashTable.insert("cherry", 15)

	if value, found := hashTable.get("apple"); found {
		fmt.Println("Value for key 'apple':", value)
	} else {
		fmt.Println("Key 'apple' not found")
	}

	if value, found := hashTable.get("banana"); found {
		fmt.Println("Value for key 'banana':", value)
	} else {
		fmt.Println("Key 'banana' not found")
	}

	if value, found := hashTable.get("grape"); found {
		fmt.Println("Value for key 'grape':", value)
	} else {
		fmt.Println("Key 'grape' not found")
	}
}
