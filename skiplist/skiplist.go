package skiplist

import (
	"iter"
	"math/rand"
)

const maxLevel = 32
const probability = 0.25

// Node represents a node in the skip list
type Node[V any] struct {
	value   V
	forward []*Node[V]
	key     uint
}

// SkipList represents the skip list data structure
type SkipList[V any] struct {
	header *Node[V]
	level  int
	length uint
}

// NewSkipList creates a new skip list
func NewSkipList[V any]() *SkipList[V] {
	return &SkipList[V]{
		header: &Node[V]{
			forward: make([]*Node[V], maxLevel),
		},
		level: 0,
	}
}

// randomLevel generates a random level for a new node
func randomLevel() int {
	level := 1
	for rand.Float64() < probability && level < maxLevel {
		level++
	}
	return level
}

// Insert adds a new key-value pair to the skip list
func (sl *SkipList[V]) Insert(key uint, value V) {
	update := make([]*Node[V], maxLevel)
	current := sl.header

	// Find where to insert the new node
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}
	current = current.forward[0]

	// Create a new node if the key doesn't exist
	if current == nil || current.key != key {
		level := randomLevel()

		// Update the skip list level if necessary
		if level > sl.level {
			for i := sl.level; i < level; i++ {
				update[i] = sl.header
			}
			sl.level = level
		}

		// Create and link the new node
		newNode := &Node[V]{
			value:   value,
			key:     key,
			forward: make([]*Node[V], level),
		}

		// Update the forward pointers
		for i := 0; i < level; i++ {
			newNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = newNode
			sl.length++

		}
	} else {
		// Update value if key already exists
		current.value = value
	}
}

// Search finds a value by its key
func (sl *SkipList[V]) Search(key uint) (V, bool) {
	current := sl.header

	// Traverse the skip list
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}

	current = current.forward[0]

	// Check if we found the key
	if current != nil && current.key == key {
		return current.value, true
	}

	// Return zero value if not found
	return *new(V), false
}

// Delete removes a node with the given key
func (sl *SkipList[V]) Delete(key uint) bool {

	if sl.length < 1 {
		return false
	}
	update := make([]*Node[V], maxLevel)
	current := sl.header

	// Find the node to delete
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}
	current = current.forward[0]

	// Delete the node if found
	if current != nil && current.key == key {
		for i := 0; i < sl.level; i++ {
			if update[i].forward[i] != current {
				break
			}
			update[i].forward[i] = current.forward[i]
		}

		// Update the skip list level if necessary
		for sl.level > 1 && sl.header.forward[sl.level-1] == nil {
			sl.level--
		}
		sl.length--
		return true
	}
	return false
}

// Iterator returns a channel that yields all key-value pairs in order
func (sl *SkipList[V]) Iterator() iter.Seq2[uint, V] {
	return func(yield func(uint, V) bool) {
		current := sl.header.forward[0]
		for current != nil {
			if !yield(current.key, current.value) {
				return
			}
			current = current.forward[0]
		}
	}
}

func (sl *SkipList[V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {

		for _, v := range sl.Iterator() {
			if !yield(v) {
				return
			}
		}
	}
}

// Len returns the number of elements in the skip list
func (sl *SkipList[V]) len() uint {
	var count uint = 0
	current := sl.header.forward[0]
	for current != nil {
		count++
		current = current.forward[0]
	}
	return count
}

// Len returns the number of elements in the skip list
func (sl *SkipList[V]) Length() uint {
	return sl.length
}
