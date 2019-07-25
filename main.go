package iter

type Iterator interface {
	Next() (int, bool)
	Map(fn ...func(p int) int) Iterator
	Reverse() Iterator
}

type iter struct {
	index      int
	items      []int
	decorators []func(p int) int
	getIndex   func(p *iter) (i int, hasNext bool)
	isReversed bool
}

func NewIter(pts []int) Iterator {
	return &iter{items: pts, getIndex: nextIndex}
}

func (p *iter) Next() (pt int, ok bool) {
	if i, hasNext := p.getIndex(p); hasNext {
		pt = p.items[i]
		for _, f := range p.decorators {
			pt = f(pt)
		}
		return pt, true
	}
	return
}

func (p *iter) Map(fns ...func(p int) int) Iterator {
	p.decorators = append(p.decorators, fns...)
	return p
}

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
