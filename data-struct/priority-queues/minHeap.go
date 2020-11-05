package heap

import (
	"math"
	"fmt"
)

/*******************************************************************************
 * A min binary heap implements the Priority Queue ADT. It has constant access
 * to the min element of the heap, with logarithmic insertions and deletions.
 *
 * add(element) - O(log n)
 * poll() - O(log n)
 * peek() - O(1)
 *
 * For more info, refer to https://en.wikipedia.org/wiki/Binary_heap
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

// MinHeap struct
type MinHeap struct {
	Heap []*Node
	Size int
}

// NewMinHeap creates new minHeap O(n) using heapify()
func NewMinHeap(elems ...*Node) *MinHeap {
	mh := new(MinHeap)
	mh.Size = 0

	if len(elems) > 0 {
		mh.Heapify(elems)
	}

	return mh
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

// Heapify builds the heap using nodes O(n)
func (m *MinHeap) Heapify(arr []*Node) {
	m.Heap = arr
	i := math.Max(0, math.Floor(float64(len(arr) / 2)) - 1)
	for ; i >= 0; i-- {
		m.Sink(int(i))
	}
	m.Size = len(arr)
}

// Sink bubbles down item O(log n)
func (m* MinHeap) Sink(idx int) {
	for {
		Left, Right := m.Left(idx), m.Right(idx)
		smallestIndex := Left

		// break if Right is out of bounds
		if Right > len(m.Heap) {
			break
		}

		// check if Right is smaller that Left, if it is, Swap
		rightLessThanLeft := Right < len(m.Heap) && m.Heap[Right].Compare(m.Heap[Left]) == -1
		if rightLessThanLeft {
			smallestIndex = Right
		}

		indexOutOfBounds := Left >= len(m.Heap)
		elementLessThanSmallest := m.Heap[idx].Compare(m.Heap[smallestIndex]) == -1


		if (indexOutOfBounds || elementLessThanSmallest) {
			break
		}

		m.Swap(smallestIndex, idx)
		idx = smallestIndex
	}
}

// Insert element, while maintaining heap invariant O(log n)
func (m *MinHeap) Insert(n *Node){
	m.Heap = append(m.Heap, n)
	idxLastElement := len(m.Heap) - 1
	m.Swim(idxLastElement)
	m.Size++
}

// Swim bubbles up an element until heap invariant is satisfied O(log n)
func (m *MinHeap) Swim(idx int) {
	Parent:=m.Parent(idx)
	elementLessThanParent := m.Heap[idx].Compare(m.Heap[Parent]) == -1
	for idx > 0 && elementLessThanParent {
		m.Swap(idx, Parent)
		idx = Parent
		Parent = m.Parent(idx)
	}
}

// Peek returns top most element in heap O(1)
func (m *MinHeap) Peek() *Node {
	return m.Heap[0]
}

// Poll returns and removes the top most element in heap O(log n)
func (m *MinHeap) Poll() (*Node, bool) {
	if m.Size == 0 {
		return nil, true
	}

	return m.RemoveAt(0), false
}

// RemoveAt removes element at index O(log n)
func (m *MinHeap) RemoveAt(idx int) *Node {
	idxLastElement := m.Size - 1

	removedElement := m.Heap[idx]

	m.Swap(idx, idxLastElement)
	m.Heap = m.Heap[:m.Size - 1]
	m.Size--

	// if last element is to be removed, dont sink/swim
	lastElementIsToRemove := idx == idxLastElement
	if lastElementIsToRemove {
		return removedElement
	}

	// sink
	idxToBeHeapified := idx
	elementToBeHeapified := m.Heap[idxToBeHeapified]
	m.Sink(idxToBeHeapified)

	// if sinking wont work try swimming
	if m.Heap[idxToBeHeapified] == elementToBeHeapified {
		m.Swim(idxToBeHeapified)
	}

	return removedElement


}

// Contains checks if element in heap
func (m *MinHeap) Contains(elm int) bool {
	_, ok := m.find(elm)
	if ok {
		return true
	}
	return false
}

/*
* Helper functions
*/
func (m *MinHeap) find(elm int) (int, bool) {
	for i, v := range m.Heap {
		if v.Value == elm{
			return i, true
		}
	}
	return -1, false
}

// Parent returns parent index
func (m *MinHeap) Parent(idx int) int {
	return (idx - 1) / 2
}
// Left returns left child idx
func (m *MinHeap) Left(idx int) int {
	return idx * 2 + 1
}
// Right returns right child idx
func (m *MinHeap) Right(idx int) int {
	return idx * 2 + 2
}
// Swap swaps idxs
func (m *MinHeap) Swap(idx1, idx2 int) {
	m.Heap[idx1], m.Heap[idx2] = m.Heap[idx2], m.Heap[idx1]
}
// PrintHeap prints the heap
func (m *MinHeap) PrintHeap(){
	for _, v := range m.Heap{
		fmt.Println(v)
	}
}

func main() {

	buildHeap := []*Node{
		&Node{Value: 10, Weight: 10},
		&Node{Value: 20, Weight: 20},
		&Node{Value: 30, Weight: 30},
		&Node{Value: 5, Weight: 5},
		&Node{Value: 7, Weight:7},
		&Node{Value: 9, Weight: 9},
		&Node{Value: 11, Weight: 11},
		&Node{Value: 13, Weight: 13},
		&Node{Value: 15, Weight: 15},
		&Node{Value: 17, Weight: 17},
	}

	m := NewMinHeap(buildHeap...)

	m.Insert(&Node{Value: 1, Weight: 1})
	fmt.Println(m.Size)

	// m.PrintHeap()

	fmt.Println(m.Poll())
	fmt.Println(m.Size)

	fmt.Println(m.Contains(15))
	fmt.Println(m.Contains(151))

}


// https://github.com/jeffzh4ng/algorithms-and-data-structures/blob/master/src/data-structures/priority-queues/min-binary-Heap.ts
// https://github.com/python/cpython/blob/3.9/Lib/heapq.py
