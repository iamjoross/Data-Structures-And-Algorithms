package main

import (
	"math"
	"fmt"
)

type comparer interface {
	compare(b node) int
}

type node struct {
	weight, value int
}

func newNode(value, weight int) (n *node){
	n = new(node)
	n.value = value
	n.weight = weight
	return
}


type minHeap struct {
	heap []*node
	size int
}

func newMinHeap(elems ...*node) *minHeap {
	mh := new(minHeap)
	mh.size = 0

	if len(elems) > 0 {
		mh.heapify(elems)
	}

	return mh
}

func (a *node) Compare(b *node) int {
	if a.weight < b.weight {
		return -1
	} else if a.weight == b.weight {
		return 0
	} else {
		return 1
	}
}

func (m *minHeap) heapify(arr []*node) {
	m.heap = arr
	i := math.Max(0, math.Floor(float64(len(arr) / 2)) - 1)
	for ; i >= 0; i-- {
		m.sink(int(i))
	}
	m.size = len(arr)
}

func (m* minHeap) sink(idx int) {
	for {
		left, right := m.left(idx), m.right(idx)
		smallestIndex := left

		// break if right is out of bounds
		if right > len(m.heap) {
			break
		}

		// check if right is smaller that left, if it is, swap
		rightLessThanLeft := right < len(m.heap) && m.heap[right].Compare(m.heap[left]) == -1
		if rightLessThanLeft {
			smallestIndex = right
		}

		indexOutOfBounds := left >= len(m.heap)
		elementLessThanSmallest := m.heap[idx].Compare(m.heap[smallestIndex]) == -1


		if (indexOutOfBounds || elementLessThanSmallest) {
			break
		}

		m.swap(smallestIndex, idx)
		idx = smallestIndex
	}
}

func (m *minHeap) insert(n *node){
	m.heap = append(m.heap, n)
	idxLastElement := len(m.heap) - 1
	m.swim(idxLastElement)
	m.size++
}

func (m *minHeap) swim(idx int) {
	parent:=m.parent(idx)
	elementLessThanParent := m.heap[idx].Compare(m.heap[parent]) == -1
	for idx > 0 && elementLessThanParent {
		m.swap(idx, parent)
		idx = parent
		parent = m.parent(idx)
	}
}


/**
* Helper functions
*/
func (m *minHeap) parent(idx int) int {
	return (idx - 1) / 2
}
func (m *minHeap) left(idx int) int {
	return idx * 2 + 1
}
func (m *minHeap) right(idx int) int {
	return idx * 2 + 2
}
func (m *minHeap) swap(idx1, idx2 int) {
	m.heap[idx1], m.heap[idx2] = m.heap[idx2], m.heap[idx1]
}
func (m *minHeap) printHeap(){
	for _, v := range m.heap{
		fmt.Println(v)
	}
}

func main() {

	buildHeap := []*node{
		&node{value: 10, weight: 10},
		&node{value: 20, weight: 20},
		&node{value: 30, weight: 30},
		&node{value: 5, weight: 5},
		&node{value: 7, weight:7},
		&node{value: 9, weight: 9},
		&node{value: 11, weight: 11},
		&node{value: 13, weight: 13},
		&node{value: 15, weight: 15},
		&node{value: 17, weight: 17},
	}

	m := newMinHeap(buildHeap...)

	m.insert(&node{value: 1, weight: 1})
	fmt.Println(m.size)

	m.printHeap()

}


// https://github.com/jeffzh4ng/algorithms-and-data-structures/blob/master/src/data-structures/priority-queues/min-binary-heap.ts
