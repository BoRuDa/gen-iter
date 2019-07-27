package iter

type IntIterator interface {
	Next() (int, bool)
	ApplyForEach(fn ...func(p int) int) IntIterator
	Reverse() IntIterator
	Slice() []int
}

type iter_int struct {
	index      int
	items      []int
	modifyFns  []func(p int) int
	getIndex   func(p *iter_int) (i int, hasNext bool)
	isReversed bool
}

func NewIter(items []int) IntIterator {
	return &iter_int{items: items, getIndex: nextIndex}
}

// Next returns current element and return "true" on success
// returns "false" when reaches the end
func (p *iter_int) Next() (pt int, ok bool) {
	if i, hasNext := p.getIndex(p); hasNext {
		for _, f := range p.modifyFns {
			p.items[i] = f(p.items[i])
		}
		return p.items[i], true
	}
	return
}

// ApplyForEach applies specified functions to each element;
// executes all "fns" in order of their definitions;
// calls only on "Next()"
func (p *iter_int) ApplyForEach(fns ...func(p int) int) IntIterator {
	p.modifyFns = append(p.modifyFns, fns...)
	return p
}

// Reverse changes direction of iterations
// 0 -> N | N -> 0
func (p *iter_int) Reverse() IntIterator {
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

// Slice returns underling slice with all changes
func (p *iter_int) Slice() []int {
	return p.items
}

func nextIndex(p *iter_int) (i int, hasNext bool) {
	if p.index < len(p.items) {
		i = p.index
		p.index++
		hasNext = true
	}
	return
}

func previousIndex(p *iter_int) (i int, hasNext bool) {
	if p.index >= 0 {
		i = p.index
		p.index--
		hasNext = true
	}
	return
}
