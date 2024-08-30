package pool

import (
	s "github.com/WatchJani/stack"
)

type Pool[T any] struct {
	store   []T
	pointer int
	stack   s.Stack[int]
}

func New[T any](capacity int) Pool[T] {
	return Pool[T]{
		store: make([]T, capacity),
		stack: s.New[int](20),
	}
}

func (m *Pool[T]) Insert() *T {
	//if something has been deleted, allocate that memory first
	if !m.stack.IsEmpty() {
		index, _ := m.stack.Pop()
		return &m.store[index]
	}

	//if the pool reaches its maximum, allocate additional memory
	if m.pointer == len(m.store) {
		newStore := make([]T, 2*m.pointer)
		copy(newStore, m.store)
		m.store = newStore
	}

	m.pointer++
	return &m.store[m.pointer-1]
}

// this part of memory can reused for later
func (m *Pool[T]) Delete(index int) {
	m.stack.Push(index)
}
