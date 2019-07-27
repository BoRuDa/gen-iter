package main

// TODO: add warning
const iteratorTemplate = `package {{ .PkgName }}

type {{.IteratorName}} interface {
	Next() ({{.Type}}, bool)
	ApplyForEach(fn ...func(p {{.Type}}) {{.Type}}) {{.IteratorName}}
	Reverse() {{.IteratorName}}
	Slice() []{{.Type}}
}

type iter_{{.Type}} struct {
	index      int
	items      []{{.Type}}
	modifyFns  []func(p {{.Type}}) {{.Type}}
	getIndex   func(p *iter_{{.Type}}) (i {{.Type}}, hasNext bool)
	isReversed bool
}

func NewIter(items []int) {{.IteratorName}} {
	return &iter_{{.Type}}{items: items, getIndex: nextIndex}
}

// Next returns current element and return "true" on success
// returns "false" when reaches the end
func (p *iter_{{.Type}}) Next() (pt {{.Type}}, ok bool) {
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
func (p *iter_{{.Type}}) ApplyForEach(fns ...func(p {{.Type}}) {{.Type}}) {{.IteratorName}} {
	p.modifyFns = append(p.modifyFns, fns...)
	return p
}

// Reverse changes direction of iterations
// 0 -> N | N -> 0
func (p *iter_{{.Type}}) Reverse() {{.IteratorName}} {
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
func (p *iter_{{.Type}}) Slice() []{{.Type}} {
	return p.items
}

func nextIndex(p *iter_{{.Type}}) (i {{.Type}}, hasNext bool) {
	if p.index < len(p.items) {
		i = p.index
		p.index++
		hasNext = true
	}
	return
}

func previousIndex(p *iter_{{.Type}}) (i {{.Type}}, hasNext bool) {
	if p.index >= 0 {
		i = p.index
		p.index--
		hasNext = true
	}
	return
}
`
