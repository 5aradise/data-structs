package ds

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"strings"
)

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrListBounds      = errors.New("list bounds out of range")
)

type cslList[E comparable] struct {
	data []E
}

func NewCSLList[E comparable](vs ...E) *cslList[E] {
	return &cslList[E]{vs}
}

func (l *cslList[E]) String() string {
	var b strings.Builder
	b.Grow(10 + l.Length()*2 - 1)
	b.WriteString("cslList{ ")
	switch data := any(l.data).(type) {
	case []rune:
		if len(data) > 0 {
			for i := range len(data) - 1 {
				b.WriteString(fmt.Sprintf("%c ", data[i]))
			}
			b.WriteString(fmt.Sprintf("%c", data[len(data)-1]))
		}
	default:
		if len(l.data) > 0 {
			for i := range len(l.data) - 1 {
				b.WriteString(fmt.Sprintf("%v ", l.data[i]))
			}
			b.WriteString(fmt.Sprintf("%v", l.data[len(l.data)-1]))
		}
	}

	b.WriteString(" }")
	return b.String()
}

func (l *cslList[E]) Length() int {
	return len(l.data)
}

func (l *cslList[E]) Clone() *cslList[E] {
	return &cslList[E]{slices.Clone(l.data)}
}

func (l *cslList[E]) Get(i int) (E, error) {
	if i < 0 {
		var zero E
		return zero, fmt.Errorf("%w [%d]", ErrIndexOutOfRange, i)
	}
	if i >= len(l.data) {
		var zero E
		return zero, fmt.Errorf("%w [%d] with length %d", ErrIndexOutOfRange, i, len(l.data))
	}
	return l.data[i], nil
}

func (l *cslList[E]) FindFirst(v E) int {
	return slices.Index(l.data, v)
}

func (l *cslList[E]) FindLast(v E) int {
	for i := len(l.data) - 1; i >= 0; i-- {
		if l.data[i] == v {
			return i
		}
	}
	return -1
}

func (l *cslList[E]) Append(v E) {
	l.data = append(l.data, v)
}

func (l *cslList[E]) Insert(v E, i int) error {
	if i < 0 {
		return fmt.Errorf("%w [%d:]", ErrListBounds, i)
	}
	if i > len(l.data) {
		return fmt.Errorf("%w [%d:%d]", ErrListBounds, i, len(l.data))
	}

	l.data = slices.Insert(l.data, i, v)
	return nil
}

func (l *cslList[E]) Extend(es *cslList[E]) {
	if es == nil {
		return
	}
	l.data = append(l.data, es.data...)
}

func (l *cslList[E]) Reverse() {
	slices.Reverse(l.data)
}

func (l *cslList[E]) Delete(i int) (E, error) {
	v, err := l.Get(i)
	if err != nil {
		return v, err
	}

	l.data = slices.Delete(l.data, i, i+1)
	return v, nil
}

func (l *cslList[E]) DeleteAll(v E) {
	l.data = slices.DeleteFunc(l.data, func(curr E) bool {
		return curr == v
	})
}

func (l *cslList[E]) Clear() {
	l.data = nil
}

func (l *cslList[E]) Iter() iter.Seq2[int, E] {
	return slices.All(l.data)
}
