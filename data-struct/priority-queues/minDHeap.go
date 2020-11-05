package heap

import (
	"math"
	"fmt"
)
/*******************************************************************************
 * A D-ary implements the Priority Queue ADT, just like the binary heap.
 * What's different is that it has d children.
 *
 * D-ary heaps are better for decreaseKey operations because decreaseKey()
 * means we must swim the node up (min heap). Since d > 2, log_d(n) < log_2(n),
 * meaning we swim up fewer levels.
 *
 * This is preferred for algorithms with heavy decreaseKey() calls like
 * Djikstra's shortest path and Prim's minimum spanning tree.
 *
 *
 * insert(element) - O(log_d n) no comparisons needed for swimming up
 * poll() - O(d log_d n) (d child comparisons when sinking to find min child)
 * decreaseKey() - O(log_d n) bc we swim up
 * peek() - O(1)
 *
 * More info can be found here: https://en.wikipedia.org/wiki/D-ary_heap
 ******************************************************************************/

 // Comparer interface that provides method signature for Compare
type Comparer interface {
	Compare(b Node) int
}

// Node contains Weight and Value
type Node struct {
	Weight, Value int
}

// NewNode creates new node
func NewNode(Value, Weight int) (n *Node){
	n = new(Node)
	n.Value = Value
	n.Weight = Weight
	return
}

// MinDHeap struct
type MinDHeap struct {
	heap []*Node
	size, deg int
}

// NewMinHeap creates new minHeap O(n) using heapify()
func New(degree int, elems ...*Node) *MinDHeap {
	mh := new(MinDHeap)
	mh.size = 0
	mh.deg = int(math.Max(2, float64(degree)))

	if len(elems) > 0 {
		mh.Heapify(elems)
	}

	return mh
}

func (m *MinDHeap) Heapify(arr []*Node) {
	for i:= 0; i < len(arr); i++ {
		m.Insert(arr[i])
	}
}

// Compare method to check equality of two nodes based on weight
func (a *Node) Compare(b *Node) int {
	if a.Weight < b.Weight {
		return -1
	} else if a.Weight == b.Weight {
		return 0
	} else {
		return 1
	}
}

// Size returns heap size
func (m *MinDHeap) Size() int{
	return m.size
}

// IsEmpty ...
func (m *MinDHeap) IsEmpty() bool{
	return m.size == 0
}

// Insert ...
func (m *MinDHeap) Insert(n *Node){
	m.heap = append(m.heap, n)
	m.size++
	idxLastElement := m.size - 1
	m.Swim(idxLastElement)
}

func (m *MinDHeap) Peek() *Node {
	return m.heap[0]
}

func (m *MinDHeap) Contains(elm int) bool {
	_, ok := m.find(elm)
	if ok {
		return true
	}
	return false
}

func (m *MinDHeap) Poll() (*Node, bool) {
	if m.size == 0 {
		return nil, true
	}

	return m.RemoveAt(0), false
}

func (m *MinDHeap) RemoveAt(idxToRemove int) *Node {
	idxLastElement := m.size - 1

	removedElement := m.heap[idxToRemove]

	m.swap(idxToRemove, idxLastElement)
	m.heap = m.heap[:m.size - 1]
	m.size--

	// if last element is to be removed, dont sink/swim
	lastElementIsToRemove := idxToRemove == idxLastElement
	if lastElementIsToRemove {
		return removedElement
	}

	// sink
	idxToBeHeapified := idxToRemove
	elementToBeHeapified := m.heap[idxToBeHeapified]
	m.Sink(idxToBeHeapified)

	// if sinking wont work try swimming; element did not move
	if m.heap[idxToBeHeapified] == elementToBeHeapified {
		m.Swim(idxToBeHeapified)
	}

	return removedElement


}

func (m *MinDHeap) Swim(idx int) {
	parentIdx := m.parent(idx)
	elementLessThanParent := m.heap[idx].Compare(m.heap[parentIdx]) == -1
	for idx > 0 && elementLessThanParent {
		m.swap(idx, parentIdx)
		idx = parentIdx
		parentIdx = m.parent(idx)
	}
}

func (m *MinDHeap) Sink(idx int) {
	for {
		childrenIdxs := m.getChildrenIndices(idx)

		smallestIdx := childrenIdxs[0]

		for _, childIdx := range childrenIdxs {
			childIdxInBounds := childIdx < m.Size()
			childIdxLessThanSmallest := m.heap[childIdx].Compare(m.heap[smallestIdx]) == -1
			if childIdxInBounds && childIdxLessThanSmallest {
				smallestIdx = childIdx
			}
		}

		indexOutOfBounds := smallestIdx >= m.Size()
		elementLessThanChild := m.heap[idx].Compare(m.heap[smallestIdx]) == -1
		if (indexOutOfBounds || elementLessThanChild) {
			break
		}

		m.swap(smallestIdx, idx)
		idx = smallestIdx
	}
}
/*
* Helper functions
*/
func (m *MinDHeap) find(elm int) (int, bool) {
	for i, v := range m.heap {
		if v.Value == elm{
			return i, true
		}
	}
	return -1, false
}
func (m *MinDHeap) getChildrenIndices(idx int) []int {
	var indices []int
	for i:= 1; i <= m.deg; i++ {
		indices = append(indices, idx * m.deg + i)
	}
	return indices
}
// Parent returns parent index
func (m *MinDHeap) parent(idx int) int {
	return (idx - 1) / 2
}
// Swap swaps idxs
func (m *MinDHeap) swap(idx1, idx2 int) {
	m.heap[idx1], m.heap[idx2] = m.heap[idx2], m.heap[idx1]
}
// PrintHeap prints the heap
func (m *MinDHeap) PrintHeap(){
	for _, v := range m.heap{
		fmt.Println(v)
	}
}
