package iter

import (
	"sort"
)

type Iterator interface {
	Next() (int, bool)
	ApplyForEach(fn ...func(p int) int) Iterator
	Reverse() Iterator
	Sort() Iterator
	Slice() []int
}

type iter struct {
	index      int
	items      []int
	decorators []func(p int) int
	getIndex   func(p *iter) (i int, hasNext bool)
	isReversed bool
}

func NewIter(items []int) Iterator {
	return &iter{items: items, getIndex: nextIndex}
}

// Next returns current element and return `true` on success
// returns `false` when reaches the end
func (p *iter) Next() (pt int, ok bool) {
	if i, hasNext := p.getIndex(p); hasNext {
		for _, f := range p.decorators {
			p.items[i] = f(p.items[i])
		}
		return p.items[i], true
	}
	return
}

// ApplyForEach applies specified functions to each element;
// executes all `fns` in order of their definitions;
// calls only on `Next()`
func (p *iter) ApplyForEach(fns ...func(p int) int) Iterator {
	p.decorators = append(p.decorators, fns...)
	return p
}

// Reverse changes direction of iterations
// 0 -> N | N -> 0
func (p *iter) Reverse() Iterator {
	if p.isReversed {
		p.getIndex = nextIndex
		p.index = 0
		p.isReversed = false
		return p
	}

	p.getIndex = previousIndex
	p.index = len(p.items) - 1
	p.isReversed = true
	return p
}

// Sort does sorting in ASC order
func (p *iter) Sort() Iterator {
	sort.Slice(p.items, func(i, j int) bool {
		return p.items[i] < p.items[j]
	})
	return p
}

// Slice returns underling slice with all changes
func (p *iter) Slice() []int {
	return p.items
}

func nextIndex(p *iter) (i int, hasNext bool) {
	if p.index < len(p.items) {
		i = p.index
		p.index++
		hasNext = true
	}
	return
}

func previousIndex(p *iter) (i int, hasNext bool) {
	if p.index >= 0 {
		i = p.index
		p.index--
		hasNext = true
	}
	return
}
