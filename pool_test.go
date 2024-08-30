package pool

import (
	"testing"
	"time"
)

const TestNumber = 1000

// 9.400 ns/op
func BenchmarkPointer(b *testing.B) {
	b.StopTimer()

	memory := New[int](TestNumber)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		*memory.Insert() = i
	}
}

type Node struct {
	time  time.Time
	value int
	key   int
	next  *Node
}

// 55 ns
func BenchmarkSpecificTest(b *testing.B) {
	b.StopTimer()

	memory := New[Node](b.N)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		*memory.Insert() = Node{
			time:  time.Now(),
			key:   i,
			value: i,
		}
	}
}
