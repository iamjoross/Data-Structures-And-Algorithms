package heap

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
	"fmt"
)

func TestHeapify(t *testing.T){
	m := &MinHeap{}
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
	m.Heapify(buildHeap)
	assert.Equal(t, m.Size, 10)

	m = nil
	buildHeap = nil
}

func BenchmarkHeapify(b *testing.B){
	m := &MinHeap{}

	maxOps := 5

	for op:= 1; op <= maxOps; op++ {
		b.Run("operation "+ strconv.Itoa(op), func(b *testing.B){
			bh := []*Node{}
			ops := b.N
			fmt.Println("-", ops)
			for i := 1; i <= ops; i++ {
				bh = append(bh, &Node{Value: i, Weight: i})
			}
			m.Heapify(bh)
		})
	}

	// buildHeap := []*Node{
	// 	&Node{Value: 10, Weight: 10},
	// 	&Node{Value: 20, Weight: 20},
	// 	&Node{Value: 30, Weight: 30},
	// 	&Node{Value: 5, Weight: 5},
	// 	&Node{Value: 7, Weight:7},
	// 	&Node{Value: 9, Weight: 9},
	// 	&Node{Value: 11, Weight: 11},
	// 	&Node{Value: 13, Weight: 13},
	// 	&Node{Value: 15, Weight: 15},
	// 	&Node{Value: 17, Weight: 17},
	// }
	// m.Heapify(buildHeap)

}

func TestPoll(t *testing.T){
	m := &MinHeap{}
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
	m.Heapify(buildHeap)
	v, _ := m.Poll()
	assert.Equal(t, v.Value, 5)
	assert.Equal(t, m.Size, 9)

	m = nil
	buildHeap = nil
}

func BenchmarkPoll(b *testing.B){
	m := &MinHeap{}
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
	m.Heapify(buildHeap)
	m.Poll()

	m = nil
	buildHeap = nil
}


// go test -run=^$ -bench=.

// go test -run=^$ -bench=. | tee -a change-benchmark
