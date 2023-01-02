package main

// from: https://www.works-hub.com/learn/learn-hashmaps-by-implementing-one-in-golang-3759a

// HashMaps are just key-value storage data structures
// pass in a key and you can get back the value associated with that key
//
// we have:
// - a hash function
// - a linear array that the hash function maps to
// - actual data nodes that hold our key-value pairs
//
// need a statically allocated array of some size n. this will hold the pointer
// of our actual key-value pairs
//
// typical implementation of hash maps relies on a good hash function.
// job of hash function is to convert the value passed in and return an integer
// value representing an index on the array. This integer value
// **has to be the same every time**
//
// we have 3 problems to solve at hand:
// - how to implement a hash function
// - make the hash value fall in the range of our array
// - how to manage hash collisions
//
// there are a couple ways of handling collisions in hashmaps:
// - separate chaining with linked lists
// - open addressing
// - double hashing
// the code below implements separate chaining with linked lists
//
// fetch is identical to insert -- pass the key as input to the hash function
// and get a hash value, which is then mapped to the array using the `getIndex`
// function
// have 3 possible outcomes:
// - check if the key on that index is matching what we are looking for
//      - if yes, we have our match
// - if the key does not match, we check if its next value is not nil -- basically
//   checking for collision and find that element in the separate chained linked
//   list
// - if both of the above fail, the requested key does not exist in our hash map

import (
	"bytes"
	"fmt"
)

// this is assuming that we have a fixed size array and don't account for resizing.
const MAP_SIZE = 50

type Node struct {
	key   string
	value string
	next  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("<Key: %s, Value: %s>\n", n.key, n.value)
}

type HashMap struct {
	Data []*Node
}

func (h *HashMap) Insert(key string, value string) {
	// calls the hash function internally and gives us the index of our array
	// where we store the key-value pair
	index := getIndex(key)

	if h.Data[index] == nil {
		// index is empty, so go ahead and insert
		h.Data[index] = &Node{key: key, value: value, next: nil}
	} else {
		// this is a collision, get into linked-list mode
		startingNode := h.Data[index]
		for ; startingNode.next != nil; startingNode = startingNode.next {
			if startingNode.key == key {
				// the key exists, its a modifying operation
				startingNode.value = value
				return
			}
		}
		startingNode.next = &Node{key: key, value: value, next: nil}
	}
}

func (h *HashMap) Get(key string) (string, bool) {
	index := getIndex(key)
	if h.Data[index] != nil {
		// key is on this index, but might be somewhere in the linked list
		startingNode := h.Data[index]
		for ; ; startingNode = startingNode.next {
			if startingNode.key == key {
				// key matched
				return startingNode.value, true
			}

			if startingNode.next == nil {
				break
			}
		}
	}

	// key does not exist
	return "", false
}

// this overrides the default print output for our defined HashMap type
// this is just a convenience method used for printing the entire HashMap in a
// pretty format
func (h *HashMap) String() string {
	var output bytes.Buffer
	fmt.Fprintln(&output, "{")
	for _, n := range h.Data {
		if n != nil {
			fmt.Fprintf(&output, "\t%s: %s\n", n.key, n.value)
			for node := n.next; node != nil; node = node.next {
				fmt.Fprintf(&output, "\t%s: %s\n", node.key, node.value)
			}
		}
	}

	fmt.Fprintln(&output, "}")

	return output.String()
}

func NewDict() *HashMap {
	return &HashMap{Data: make([]*Node, MAP_SIZE)}
}

func main() {
	a := NewDict()
	a.Insert("name", "jake")
	a.Insert("gender", "male")
	a.Insert("city", "Boston")
	a.Insert("lastname", "correnti")
	if value, ok := a.Get("name"); ok {
		fmt.Println(value)
	} else {
		fmt.Println("value did not match")
	}

	fmt.Println(a)
}

// Jenkins hash function that produces 32 bit hashes
func hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

// getIndex ensures that the index value for a given hashed key is within the
// bounds of the array
func getIndex(key string) int {
	return int(hash(key)) % MAP_SIZE
}
