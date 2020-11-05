package main

import (
	MinDHeap "./goheap/MinDHeap"
	// "fmt"
)

func main() {

	buildHeap := []*MinDHeap.Node{
		&MinDHeap.Node{Value: 10, Weight: 10},
		&MinDHeap.Node{Value: 20, Weight: 20},
		&MinDHeap.Node{Value: 30, Weight: 30},
		&MinDHeap.Node{Value: 5, Weight: 5},
		&MinDHeap.Node{Value: 7, Weight:7},
		&MinDHeap.Node{Value: 9, Weight: 9},
		&MinDHeap.Node{Value: 11, Weight: 11},
		&MinDHeap.Node{Value: 13, Weight: 13},
		&MinDHeap.Node{Value: 15, Weight: 15},
		&MinDHeap.Node{Value: 17, Weight: 17},
	}


	md := MinDHeap.New(2, buildHeap...)

	md.PrintHeap()

	// m.Insert(&Node{Value: 1, Weight: 1})
	// fmt.Println(m.Size)

	// // m.PrintHeap()

	// fmt.Println(m.Poll())
	// fmt.Println(m.Size)

	// fmt.Println(m.Contains(15))
	// fmt.Println(m.Contains(151))
}
