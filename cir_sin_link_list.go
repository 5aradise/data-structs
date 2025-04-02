package ds

import (
	"errors"
	"fmt"
	"iter"
	"strings"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrListBounds      = errors.New("list bounds out of range")
)

type node[E comparable] struct {
	data E
	next *node[E]
}

type cslList[E comparable] struct {
	tail *node[E]
	len  int
}

func NewCSLList[E comparable](vs ...E) *cslList[E] {
	var head, tail *node[E]
	if len(vs) > 0 {
		curr := &node[E]{data: vs[0]}
		head = curr
		for _, v := range vs[1:] {
			curr.next = &node[E]{data: v}
			curr = curr.next
		}
		tail = curr
		tail.next = head
	}
	return &cslList[E]{tail, len(vs)}
}

func (l *cslList[E]) String() string {
	if (l.len == 0) {
		return "cslList{  }"
	}

	var b strings.Builder
	b.Grow(10 + l.Length()*2 - 1)
	b.WriteString("cslList{ ")
	for _, v := range l.Iter() {
		switch v := any(v).(type) {
		case rune:
			b.WriteString(fmt.Sprintf("%c ", v))
		default:
			b.WriteString(fmt.Sprintf("%v ", v))
		}
	}
	b.WriteString("}")
	return b.String()
}

func (l *cslList[E]) Length() int {
	return l.len
}

func (l *cslList[E]) Clone() *cslList[E] {
	tail := &node[E]{data: l.tail.data}
	currCopy := tail
	curr := l.tail
	for range l.len - 1 {
		curr = curr.next
		currCopy.next = &node[E]{data: curr.data}
		currCopy = currCopy.next
	}
	currCopy.next = tail
	return &cslList[E]{tail, l.len}
}

func (l *cslList[E]) Get(i int) (E, error) {
	var zero E
	if i < 0 {
		return zero, fmt.Errorf("%w [%d]", ErrIndexOutOfRange, i)
	}
	if i >= l.len {
		return zero, fmt.Errorf("%w [%d] with length %d", ErrIndexOutOfRange, i, l.len)
	}
	for idx, v := range l.Iter() {
		if idx == i {
			return v, nil
		}
	}
	return zero, fmt.Errorf("%w [%d] with length %d", ErrIndexOutOfRange, i, l.len)
}

func (l *cslList[E]) FindFirst(v E) int {
	for i, val := range l.Iter() {
		if val == v {
			return i
		}
	}
	return -1
}

func (l *cslList[E]) FindLast(v E) int {
	last := -1
	for i, val := range l.Iter() {
		if val == v {
			last = i
		}
	}
	return last
}

func (l *cslList[E]) Append(v E) {
	node := &node[E]{data: v}
	node.next = l.tail.next
	l.tail.next = node
	l.tail = node
	l.len += 1
}

func (l *cslList[E]) Insert(v E, i int) error {
	if i < 0 {
		return fmt.Errorf("%w [%d:]", ErrListBounds, i)
	}
	if i > l.len {
		return fmt.Errorf("%w [%d:%d]", ErrListBounds, i, l.len)
	}

	prev := l.tail
	for range i {
		prev = prev.next
	}
	node := &node[E]{v, prev.next}
	prev.next = node
	if i == l.len {
		l.tail = node
	}
	l.len++
	return nil
}

func (l *cslList[E]) Extend(es *cslList[E]) {
	copy := es.Clone()
	l.tail.next, copy.tail.next, l.tail = copy.tail.next, l.tail.next, copy.tail
	l.len += es.len
}

func (l *cslList[E]) Reverse() {
	head := l.tail.next
	prev := l.tail
	curr := head
	for range l.len {
		next := curr.next
		curr.next = prev
		prev, curr = curr, next
	}
	l.tail = head
}

func (l *cslList[E]) Delete(i int) (E, error) {
	var zero E
	if i < 0 {
		return zero, fmt.Errorf("%w [%d]", ErrIndexOutOfRange, i)
	}
	if i >= l.len {
		return zero, fmt.Errorf("%w [%d] with length %d", ErrIndexOutOfRange, i, l.len)
	}

	prev := l.tail
	for range i {
		prev = prev.next
	}
	v := prev.next.data
	prev.next = prev.next.next
	l.len--
	return v, nil
}

func (l *cslList[E]) DeleteAll(v E) {
	prev := l.tail
	for range l.len {
		curr := prev.next
		if curr.data == v {
			prev.next = curr.next
			l.len--
		}
		prev = prev.next
	}
}

func (l *cslList[E]) Clear() {
	l.tail = nil
	l.len = 0
}

func (l *cslList[E]) Iter() iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		curr := l.tail
		for i := range l.len {
			curr = curr.next
			if !yield(i, curr.data) {
				return
			}
		}
	}
}
