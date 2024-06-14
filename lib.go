package protected

import (
	"sync"
)

// Protected provides thread-safe access to a value of any type.
type Protected[T any] struct {
	mu sync.RWMutex
	val T
}

// New creates a new Protected instance.
func New[T any](initial T) *Protected[T] {
	return &Protected[T]{
		val: initial,
	}
}

// Get safely retrieves the value.
func (p *Protected[T]) Get() T {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.val
}

// Set safely updates the value.
func (p *Protected[T]) Set(val T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.val = val
}

// Update safely updates the value using a provided function.
func (p *Protected[T]) Update(updateFunc func(val T) T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.val = updateFunc(p.val)
}

func (p *Protected[T]) DoWithLock(action func(val *T)) {
	p.mu.Lock()
	defer p.mu.Unlock()
	action(&p.val)
}

func (p *Protected[T]) DoWithRLock(action func(val *T)) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	action(&p.val)
}
