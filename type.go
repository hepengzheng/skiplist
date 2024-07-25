package skiplist

import (
	"math/rand"
	"sync"
)

type elementNode[T TKey] struct {
	next []*Element[T]
}

type TKey interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~string
}

type Element[T TKey] struct {
	elementNode[T]
	key   T
	value interface{}
}

// Key allows retrieval of the key for a given Element
func (element *Element[T]) Key() T {
	return element.key
}

// Value allows retrieval of the value for a given Element
func (element *Element[T]) Value() interface{} {
	return element.value
}

// Next returns the following Element or nil if we're at the end of the list.
// Only operates on the bottom level of the skip list (a fully linked list).
func (element *Element[T]) Next() *Element[T] {
	return element.next[0]
}

type SkipList[T TKey] struct {
	elementNode[T]
	maxLevel       int
	Length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	mutex          sync.RWMutex
	prevNodesCache []*elementNode[T]
}
