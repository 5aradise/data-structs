package ds

import (
	"errors"
	"fmt"
	"slices"

	"testing"
)

func TestCSLListNewAndString(t *testing.T) {
	tests := []struct {
		name string
		data any
		want string
	}{
		{
			name: "no data",
			data: nil,
			want: "cslList{  }",
		},
		{
			name: "empty",
			data: []rune{},
			want: "cslList{  }",
		},
		{
			name: "ints",
			data: []int{1, 2, 3, 65, 123, 44, 21},
			want: "cslList{ 1 2 3 65 123 44 21 }",
		},
		{
			name: "runes",
			data: []rune{'a', 'b', 'c'},
			want: "cslList{ a b c }",
		},
		{
			name: "strings",
			data: []string{"", "a b", "ab"},
			want: "cslList{  a b ab }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var l fmt.Stringer
			if tt.data == nil {
				l = NewCSLList[any]()
			} else {
				switch data := tt.data.(type) {
				case []rune:
					l = NewCSLList(data...)
				case []int:
					l = NewCSLList(data...)
				case []string:
					l = NewCSLList(data...)
				default:
					t.Errorf("unsupported type %T", tt.data)
				}
			}

			if got := l.String(); got != tt.want {
				t.Errorf("NewCSLList(%v).String() = %v, want %v", tt.data, got, tt.want)
			}
		})
	}
}

func TestCSLListLength(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "no data",
			data: nil,
			want: 0,
		},
		{
			name: "4 elems",
			data: []int{1, 2, 3, 4},
			want: 4,
		},
		{
			name: "6 elems",
			data: []int{1, 2, 3, 4, 5, 6},
			want: 6,
		},
		{
			name: "7 elems",
			data: []int{1, 2, 3, 4, 5, 6, 8},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSLList(tt.data...).Length(); got != tt.want {
				t.Errorf("NewCSLList(%v).Length() = %v, want %v", tt.data, got, tt.want)
			}
		})
	}
}

func TestCSLListAppend(t *testing.T) {
	tests := []struct {
		name     string
		init     []int
		toAppend []int
		want     string
	}{
		{
			name:     "3 elems",
			init:     []int{0, 0, 0},
			toAppend: []int{1, 2, 3},
			want:     "cslList{ 0 0 0 1 2 3 }",
		},
		{
			name:     "2 another elems",
			init:     []int{0, 0, 0},
			toAppend: []int{-1, -2},
			want:     "cslList{ 0 0 0 -1 -2 }",
		},
		{
			name:     "a lot of elems",
			init:     []int{0, 0, 0},
			toAppend: []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3},
			want:     "cslList{ 0 0 0 1 2 3 1 2 3 1 2 3 1 2 3 1 2 3 1 2 3 1 2 3 }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)
			for _, v := range tt.toAppend {
				l.Append(v)
			}
			if got := l.String(); got != tt.want {
				t.Errorf("NewCSLList().Append(...%v).String() = %v, want %v", tt.toAppend, got, tt.want)
			}
		})
	}
}

func TestCSLListGet(t *testing.T) {
	tests := []struct {
		name     string
		init     []int
		index    int
		wantVal  int
		wantList string
		wantErr  error
	}{
		{
			name:     "get first element",
			init:     []int{10, 20, 30},
			index:    0,
			wantVal:  10,
			wantList: "cslList{ 10 20 30 }",
		},
		{
			name:     "get middle element",
			init:     []int{10, 20, 30},
			index:    1,
			wantVal:  20,
			wantList: "cslList{ 10 20 30 }",
		},
		{
			name:     "get last element",
			init:     []int{10, 20, 30},
			index:    2,
			wantVal:  30,
			wantList: "cslList{ 10 20 30 }",
		},
		{
			name:     "get from single-element list",
			init:     []int{42},
			index:    0,
			wantVal:  42,
			wantList: "cslList{ 42 }",
		},

		{
			name:     "empty list",
			init:     []int{},
			index:    0,
			wantVal:  0,
			wantList: "cslList{  }",
			wantErr:  ErrIndexOutOfRange,
		},
		{
			name:     "negative index",
			init:     []int{10, 20, 30},
			index:    -1,
			wantVal:  0,
			wantList: "cslList{ 10 20 30 }",
			wantErr:  ErrIndexOutOfRange,
		},
		{
			name:     "index equals length",
			init:     []int{10, 20, 30},
			index:    3,
			wantVal:  0,
			wantList: "cslList{ 10 20 30 }",
			wantErr:  ErrIndexOutOfRange,
		},
		{
			name:     "index beyond length",
			init:     []int{10, 20, 30},
			index:    100,
			wantVal:  0,
			wantList: "cslList{ 10 20 30 }",
			wantErr:  ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)
			gotVal, err := l.Get(tt.index)

			if gotList := l.String(); gotList != tt.wantList {
				t.Errorf("list after Get() = %v, want %v", gotList, tt.wantList)
			}

			if gotVal != tt.wantVal {
				t.Errorf("Get() returned value = %v, want %v", gotVal, tt.wantVal)
			}

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Get() error = %v, want %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Errorf("unexpected Get() error = %v", err)
			}
		})
	}
}

func TestCSLListInsert(t *testing.T) {
	tests := []struct {
		name     string
		init     []int
		toInsert []struct {
			val int
			idx int
		}
		want    string
		wantErr error
	}{
		{
			name: "insert at beginning",
			init: []int{1, 2, 3},
			toInsert: []struct {
				val int
				idx int
			}{
				{0, 0},
			},
			want: "cslList{ 0 1 2 3 }",
		},
		{
			name: "insert in middle",
			init: []int{1, 2, 3},
			toInsert: []struct {
				val int
				idx int
			}{
				{5, 1},
				{6, 3},
			},
			want: "cslList{ 1 5 2 6 3 }",
		},
		{
			name: "insert at end",
			init: []int{1, 2, 3},
			toInsert: []struct {
				val int
				idx int
			}{
				{4, 3},
			},
			want: "cslList{ 1 2 3 4 }",
		},
		{
			name: "multiple inserts",
			init: []int{1},
			toInsert: []struct {
				val int
				idx int
			}{
				{2, 1},
				{0, 0},
				{3, 3},
				{1, 1},
			},
			want: "cslList{ 0 1 1 2 3 }",
		},
		{
			name: "insert out of bounds (negative)",
			init: []int{1, 2, 3},
			toInsert: []struct {
				val int
				idx int
			}{
				{0, -1},
			},
			want:    "cslList{ 1 2 3 }",
			wantErr: ErrListBounds,
		},
		{
			name: "insert out of bounds (too large)",
			init: []int{1, 2, 3},
			toInsert: []struct {
				val int
				idx int
			}{
				{0, 4},
			},
			want:    "cslList{ 1 2 3 }",
			wantErr: ErrListBounds,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)
			var err error

			for _, ins := range tt.toInsert {
				if err = l.Insert(ins.val, ins.idx); err != nil {
					break
				}
			}

			if got := l.String(); got != tt.want {
				t.Errorf("NewCSLList().Insert(...).String() = %v, want %v", got, tt.want)
			}

			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("NewCSLList().Insert(...) error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSLListDelete(t *testing.T) {
	tests := []struct {
		name      string
		init      []int
		toDelete  []int
		wantList  string
		wantVals  []int
		wantErr   error
		stopOnErr bool
	}{
		{
			name:     "delete from beginning",
			init:     []int{1, 2, 3},
			toDelete: []int{0},
			wantList: "cslList{ 2 3 }",
			wantVals: []int{1},
		},
		{
			name:     "delete from middle",
			init:     []int{1, 2, 3, 4},
			toDelete: []int{1, 1},
			wantList: "cslList{ 1 4 }",
			wantVals: []int{2, 3},
		},
		{
			name:     "delete from end",
			init:     []int{1, 2, 3},
			toDelete: []int{2},
			wantList: "cslList{ 1 2 }",
			wantVals: []int{3},
		},
		{
			name:     "multiple deletions",
			init:     []int{1, 2, 3, 4, 5},
			toDelete: []int{0, 0, 0},
			wantList: "cslList{ 4 5 }",
			wantVals: []int{1, 2, 3},
		},
		{
			name:      "delete until empty",
			init:      []int{1},
			toDelete:  []int{0, 0, 0},
			wantList:  "cslList{  }",
			wantVals:  []int{1},
			wantErr:   ErrIndexOutOfRange,
			stopOnErr: true,
		},
		{
			name:     "delete from empty list",
			init:     []int{},
			toDelete: []int{0},
			wantList: "cslList{  }",
			wantVals: []int{0},
			wantErr:  ErrIndexOutOfRange,
		},
		{
			name:     "negative index",
			init:     []int{1, 2, 3},
			toDelete: []int{-1},
			wantList: "cslList{ 1 2 3 }",
			wantVals: []int{0},
			wantErr:  ErrIndexOutOfRange,
		},
		{
			name:     "index too large",
			init:     []int{1, 2, 3},
			toDelete: []int{3},
			wantList: "cslList{ 1 2 3 }",
			wantVals: []int{0},
			wantErr:  ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)
			var gotVals []int
			var err error

			for _, idx := range tt.toDelete {
				var val int
				val, err = l.Delete(idx)
				if err != nil {
					if tt.stopOnErr {
						break
					}
					gotVals = append(gotVals, 0)
					continue
				}
				gotVals = append(gotVals, val)
			}

			if got := l.String(); got != tt.wantList {
				t.Errorf("list after Delete() = %v, want %v", got, tt.wantList)
			}

			if !slices.Equal(gotVals, tt.wantVals) {
				t.Errorf("returned values from Delete() = %v, want %v", gotVals, tt.wantVals)
			}

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Delete() error = %v, want %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Errorf("unexpected Delete() error = %v", err)
			}
		})
	}
}

func TestCSLListDeleteAll(t *testing.T) {
	tests := []struct {
		name        string
		init        []int
		deleteValue int
		want        string
	}{
		{
			name:        "empty list",
			init:        []int{},
			deleteValue: 1,
			want:        "cslList{  }",
		},
		{
			name:        "value not present",
			init:        []int{1, 2, 3},
			deleteValue: 4,
			want:        "cslList{ 1 2 3 }",
		},
		{
			name:        "single occurrence",
			init:        []int{1, 2, 3},
			deleteValue: 2,
			want:        "cslList{ 1 3 }",
		},
		{
			name:        "multiple occurrences",
			init:        []int{1, 2, 2, 3, 2},
			deleteValue: 2,
			want:        "cslList{ 1 3 }",
		},
		{
			name:        "all elements match",
			init:        []int{5, 5, 5},
			deleteValue: 5,
			want:        "cslList{  }",
		},
		{
			name:        "first and last elements",
			init:        []int{1, 2, 3, 1},
			deleteValue: 1,
			want:        "cslList{ 2 3 }",
		},
		{
			name:        "consecutive elements",
			init:        []int{1, 1, 2, 1, 1, 3, 1, 1},
			deleteValue: 1,
			want:        "cslList{ 2 3 }",
		},
		{
			name:        "zero value",
			init:        []int{0, 1, 0, 2, 0},
			deleteValue: 0,
			want:        "cslList{ 1 2 }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)

			l.DeleteAll(tt.deleteValue)

			if got := l.String(); got != tt.want {
				t.Errorf("DeleteAll() result = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSLListClone(t *testing.T) {
	tests := []struct {
		name       string
		init       []int
		modifyOrig []int
		modifyCopy []int
		wantOrig   string
		wantCopy   string
	}{
		{
			name:       "clone empty list",
			init:       []int{},
			modifyOrig: []int{},
			modifyCopy: []int{},
			wantOrig:   "cslList{  }",
			wantCopy:   "cslList{  }",
		},
		{
			name:       "clone single-element list",
			init:       []int{42},
			modifyOrig: []int{},
			modifyCopy: []int{},
			wantOrig:   "cslList{ 42 }",
			wantCopy:   "cslList{ 42 }",
		},
		{
			name:       "clone multi-element list",
			init:       []int{1, 2, 3},
			modifyOrig: []int{},
			modifyCopy: []int{},
			wantOrig:   "cslList{ 1 2 3 }",
			wantCopy:   "cslList{ 1 2 3 }",
		},
		{
			name:       "modify original after clone",
			init:       []int{1, 2, 3},
			modifyOrig: []int{4, 5},
			modifyCopy: []int{},
			wantOrig:   "cslList{ 1 2 3 4 5 }",
			wantCopy:   "cslList{ 1 2 3 }",
		},
		{
			name:       "modify clone after clone",
			init:       []int{1, 2, 3},
			modifyOrig: []int{},
			modifyCopy: []int{4, 5},
			wantOrig:   "cslList{ 1 2 3 }",
			wantCopy:   "cslList{ 1 2 3 4 5 }",
		},
		{
			name:       "modify both after clone",
			init:       []int{1, 2, 3},
			modifyOrig: []int{9},
			modifyCopy: []int{4, 5},
			wantOrig:   "cslList{ 1 2 3 9 }",
			wantCopy:   "cslList{ 1 2 3 4 5 }",
		},
		{
			name:       "clone then clear original",
			init:       []int{1, 2, 3},
			modifyOrig: []int{-2, -1, -0}, // Assume negative means clear operation
			modifyCopy: []int{},
			wantOrig:   "cslList{  }",
			wantCopy:   "cslList{ 1 2 3 }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := NewCSLList(tt.init...)
			copy := orig.Clone()

			for _, v := range tt.modifyOrig {
				if v > 0 {
					orig.Append(v)
				} else {
					orig.Delete(-v)
				}
			}

			for _, v := range tt.modifyCopy {
				copy.Append(v)
			}

			if got := orig.String(); got != tt.wantOrig {
				t.Errorf("original list after modifications = %v, want %v", got, tt.wantOrig)
			}

			if got := copy.String(); got != tt.wantCopy {
				t.Errorf("copied list after modifications = %v, want %v", got, tt.wantCopy)
			}

			if len(tt.init) > 0 && orig == copy {
				t.Error("original and copy reference the same list, want independent copies")
			}
		})
	}
}

func TestCSLListExtend(t *testing.T) {
	tests := []struct {
		name       string
		init       []int
		extendWith []int
		want       string
	}{
		{
			name:       "extend empty with empty",
			init:       []int{},
			extendWith: []int{},
			want:       "cslList{  }",
		},
		{
			name:       "extend empty with non-empty",
			init:       []int{},
			extendWith: []int{1, 2, 3},
			want:       "cslList{ 1 2 3 }",
		},
		{
			name:       "extend non-empty with empty",
			init:       []int{1, 2, 3},
			extendWith: []int{},
			want:       "cslList{ 1 2 3 }",
		},
		{
			name:       "extend with single element",
			init:       []int{1, 2, 3},
			extendWith: []int{4},
			want:       "cslList{ 1 2 3 4 }",
		},
		{
			name:       "extend with multiple elements",
			init:       []int{1, 2, 3},
			extendWith: []int{4, 5, 6},
			want:       "cslList{ 1 2 3 4 5 6 }",
		},
		{
			name:       "extend large lists",
			init:       []int{1, 2, 3},
			extendWith: []int{4, 5, 6, 7, 8, 9, 10},
			want:       "cslList{ 1 2 3 4 5 6 7 8 9 10 }",
		},
		{
			name:       "extend with nil list",
			init:       []int{1, 2, 3},
			extendWith: nil,
			want:       "cslList{ 1 2 3 }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)

			var extendList *cslList[int]
			if tt.extendWith != nil {
				extendList = NewCSLList(tt.extendWith...)
			}

			l.Extend(extendList)

			if got := l.String(); got != tt.want {
				t.Errorf("after Extend() = %v, want %v", got, tt.want)
			}

			if extendList != nil {
				if got := extendList.String(); got != NewCSLList(tt.extendWith...).String() {
					t.Errorf("extendList was modified, got %v, want %v", got, NewCSLList(tt.extendWith...).String())
				}
			}
		})
	}
}

func TestCSLListFindFirstLast(t *testing.T) {
	tests := []struct {
		name      string
		init      []int
		searchVal int
		wantFirst int
		wantLast  int
	}{
		{
			name:      "empty list",
			init:      []int{},
			searchVal: 1,
			wantFirst: -1,
			wantLast:  -1,
		},
		{
			name:      "single element found",
			init:      []int{5},
			searchVal: 5,
			wantFirst: 0,
			wantLast:  0,
		},
		{
			name:      "single element not found",
			init:      []int{5},
			searchVal: 1,
			wantFirst: -1,
			wantLast:  -1,
		},
		{
			name:      "multiple elements - first occurrence",
			init:      []int{1, 2, 3, 2, 4},
			searchVal: 2,
			wantFirst: 1,
			wantLast:  3,
		},
		{
			name:      "multiple elements - last occurrence",
			init:      []int{1, 2, 3, 2, 4},
			searchVal: 4,
			wantFirst: 4,
			wantLast:  4,
		},
		{
			name:      "all elements match",
			init:      []int{5, 5, 5},
			searchVal: 5,
			wantFirst: 0,
			wantLast:  2,
		},
		{
			name:      "no matches",
			init:      []int{1, 2, 3, 4},
			searchVal: 5,
			wantFirst: -1,
			wantLast:  -1,
		},
		{
			name:      "first and last same",
			init:      []int{1, 2, 3, 4},
			searchVal: 2,
			wantFirst: 1,
			wantLast:  1,
		},
		{
			name:      "multiple occurrences scattered",
			init:      []int{1, 2, 1, 3, 1, 4, 1},
			searchVal: 1,
			wantFirst: 0,
			wantLast:  6,
		},
		{
			name:      "search for zero value",
			init:      []int{0, 1, 0, 2, 0},
			searchVal: 0,
			wantFirst: 0,
			wantLast:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)

			gotFirst := l.FindFirst(tt.searchVal)
			if gotFirst != tt.wantFirst {
				t.Errorf("FindFirst(%v) = %v, want %v", tt.searchVal, gotFirst, tt.wantFirst)
			}

			gotLast := l.FindLast(tt.searchVal)
			if gotLast != tt.wantLast {
				t.Errorf("FindLast(%v) = %v, want %v", tt.searchVal, gotLast, tt.wantLast)
			}

			if gotList := l.String(); gotList != NewCSLList(tt.init...).String() {
				t.Errorf("list was modified, got %v, want %v", gotList, NewCSLList(tt.init...).String())
			}
		})
	}
}

func TestCSLListReverse(t *testing.T) {
	tests := []struct {
		name string
		init []int
		want string
	}{
		{
			name: "empty list",
			init: []int{},
			want: "cslList{  }",
		},
		{
			name: "single element",
			init: []int{1},
			want: "cslList{ 1 }",
		},
		{
			name: "two elements",
			init: []int{1, 2},
			want: "cslList{ 2 1 }",
		},
		{
			name: "odd number of elements",
			init: []int{1, 2, 3, 4, 5},
			want: "cslList{ 5 4 3 2 1 }",
		},
		{
			name: "even number of elements",
			init: []int{1, 2, 3, 4},
			want: "cslList{ 4 3 2 1 }",
		},
		{
			name: "all same elements",
			init: []int{5, 5, 5},
			want: "cslList{ 5 5 5 }",
		},
		{
			name: "palindrome list",
			init: []int{1, 2, 3, 2, 1},
			want: "cslList{ 1 2 3 2 1 }",
		},
		{
			name: "large list",
			init: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: "cslList{ 10 9 8 7 6 5 4 3 2 1 }",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)
			l.Reverse()

			if got := l.String(); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}

			l.Reverse()
			if got := l.String(); got != NewCSLList(tt.init...).String() {
				t.Errorf("Double Reverse() = %v, want original %v", got, NewCSLList(tt.init...).String())
			}
		})
	}
}

func TestCSLListClear(t *testing.T) {
	tests := []struct {
		name string
		init []int
	}{
		{
			name: "empty list",
			init: []int{},
		},
		{
			name: "single element",
			init: []int{42},
		},
		{
			name: "multiple elements",
			init: []int{1, 2, 3, 4, 5},
		},
		{
			name: "large list",
			init: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "duplicate elements",
			init: []int{5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)

			l.Clear()

			if got := l.String(); got != "cslList{  }" {
				t.Errorf("Clear() result = %v, want empty list", got)
			}

			l.Append(1)
			if got := l.String(); got != "cslList{ 1 }" {
				t.Errorf("List not reusable after Clear(), got %v", got)
			}

			l.Clear()
			if got := l.String(); got != "cslList{  }" {
				t.Errorf("Second Clear() failed, got %v", got)
			}
		})
	}
}

func TestCSLListIter(t *testing.T) {
	type pair struct {
		index int
		value int
	}

	tests := []struct {
		name string
		init []int
		want []pair
	}{
		{
			name: "empty list",
			init: []int{},
			want: []pair{},
		},
		{
			name: "single element",
			init: []int{42},
			want: []pair{{0, 42}},
		},
		{
			name: "multiple elements",
			init: []int{1, 2, 3},
			want: []pair{{0, 1}, {1, 2}, {2, 3}},
		},
		{
			name: "with duplicates",
			init: []int{5, 5, 5},
			want: []pair{{0, 5}, {1, 5}, {2, 5}},
		},
		{
			name: "large list",
			init: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []pair{
				{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5},
				{5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewCSLList(tt.init...)

			var got []pair
			for idx, val := range l.Iter() {
				got = append(got, pair{idx, val})
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("Iter() produced %v, want %v", got, tt.want)
			}

			if gotList := l.String(); gotList != NewCSLList(tt.init...).String() {
				t.Errorf("list was modified during iteration, got %v", gotList)
			}

			if len(tt.init) > 0 {
				count := 0
				for range l.Iter() {
					count++
				}
				if count != len(tt.init) {
					t.Errorf("second iteration produced %d elements, want %d", count, len(tt.init))
				}
			}
		})
	}

	t.Run("concurrent modification detection", func(t *testing.T) {
		l := NewCSLList(1, 2, 3)

		var got []pair
		i := 0
		for idx, val := range l.Iter() {
			got = append(got, pair{idx, val})
			if i == 1 {
				l.Append(4)
			}
			i++
		}

		want := []pair{{0, 1}, {1, 2}, {2, 3}}
		if !slices.Equal(got, want) {
			t.Errorf("iteration after modification produced %v, want %v", got, want)
		}
	})
}
