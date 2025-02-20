package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		list     *Node[int]
		val      int
		wantErr  error
		wantList *Node[int]
	}{
		{
			name:     "Append to nil list",
			list:     nil,
			val:      1,
			wantErr:  EmptyList,
			wantList: nil,
		},
		{
			name:     "Append to single-node list",
			list:     &Node[int]{Val: 1},
			val:      2,
			wantErr:  nil,
			wantList: &Node[int]{Val: 1, Next: &Node[int]{Val: 2}},
		},
		{
			name:     "Append to multi-node list",
			list:     &Node[int]{Val: 1, Next: &Node[int]{Val: 2}},
			val:      3,
			wantErr:  nil,
			wantList: &Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.list.Append(tt.val)

			if !assert.ErrorIs(t, err, tt.wantErr) {
				t.Errorf("Append() error = %v, want %v", err, tt.wantErr)
			}

			if tt.wantList != nil {
				currGot, currWant := tt.list, tt.wantList
				for currGot != nil && currWant != nil {
					if !assert.Equal(t, currWant.Val, currGot.Val) {
						t.Errorf("List mismatch: got %v, want %v", currGot.Val, currWant.Val)
					}
					currGot, currWant = currGot.Next, currWant.Next
				}
				if currGot != nil || currWant != nil {
					t.Errorf("List lengths do not match")
				}
			}
		})
	}
}

func TestAppendAfter(t *testing.T) {
	tests := []struct {
		name     string
		list     *Node[int]
		val      int
		afterVal int
		wantErr  error
		wantList *Node[int]
	}{
		{
			"Empty list",
			nil,
			1,
			0,
			EmptyList,
			nil,
		},
		{
			"Success",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1}},
			2,
			1,
			nil,
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2}}},
		},
		{
			"Not in list",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2}}},
			4,
			3,
			NotInList,
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.AppendAfter(tt.val, tt.afterVal)
			if !assert.Equal(t, tt.wantErr, got) {
				t.Errorf("AppendAfter() error = %v, want %v", got, tt.wantErr)
			}

			if tt.wantList != nil {
				currGot, currWant := tt.list, tt.wantList
				for currGot != nil && currWant != nil {
					if !assert.Equal(t, currWant.Val, currGot.Val) {
						t.Errorf("List mismatch: got %v, want %v", currGot.Val, currWant.Val)
					}
					currGot, currWant = currGot.Next, currWant.Next
				}
				if currGot != nil || currWant != nil {
					t.Errorf("List lengths do not match")
				}
			}
		})
	}
}

func TestAppendBefore(t *testing.T) {
	tests := []struct {
		name      string
		list      *Node[int]
		val       int
		beforeVal int
		wantErr   error
		wantList  *Node[int]
	}{
		{
			"Empty list",
			nil,
			1,
			0,
			EmptyList,
			nil,
		},
		{
			"Before head val",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1}},
			-1,
			0,
			nil,
			&Node[int]{Val: -1, Next: &Node[int]{Val: 0, Next: &Node[int]{Val: 1}}},
		},
		{
			"Success",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 2}},
			1,
			2,
			nil,
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2}}},
		},
		{
			"Not in list",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1}},
			1,
			2,
			NotInList,
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := tt.list.AppendBefore(tt.val, tt.beforeVal)
			if !assert.Equal(t, tt.wantErr, err) {
				t.Errorf("AppendBefore() error = %v, want %v", err, tt.wantErr)
			}

			if tt.wantList != nil {
				currGot, currWant := got, tt.wantList
				for currGot != nil && currWant != nil {
					if !assert.Equal(t, currWant.Val, currGot.Val) {
						t.Errorf("List mismatch: got %v, want %v", currGot.Val, currWant.Val)
					}
					currGot, currWant = currGot.Next, currWant.Next
				}
				if currGot != nil || currWant != nil {
					t.Errorf("List lengths do not match")
				}
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name     string
		list     *Node[int]
		val      int
		wantHead *Node[int]
	}{
		{
			"Empty list",
			nil,
			1,
			&Node[int]{Val: 1},
		},
		{
			"Non-empty list",
			&Node[int]{Val: 2, Next: &Node[int]{Val: 3}},
			1,
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Prepend(tt.val)
			if !assert.Equal(t, tt.wantHead.Val, got.Val) {
				t.Errorf("Prepend() = %v, want %v", got.Val, tt.wantHead.Val)
			}
			if got.Next != nil && !assert.Equal(t, tt.wantHead.Next.Val, got.Next.Val) {
				t.Errorf("Prepend() Next = %v, want %v", got.Next.Val, tt.wantHead.Next.Val)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		list      *Node[int]
		separator string
		want      string
	}{
		{
			"Empty list",
			nil,
			" ",
			"",
		},
		{
			"With separator",
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
			" ",
			"1 2 3\n",
		},
		{
			"With long separator",
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
			"onpv83209nv8wbav93j",
			"1onpv83209nv8wbav93j2onpv83209nv8wbav93j3\n",
		},
		{
			"Without separator",
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
			"",
			"123\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.String(tt.separator)
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name string
		list *Node[int]
		want int
	}{
		{
			"Empty list",
			nil,
			0,
		},
		{
			"One elem",
			&Node[int]{Val: 1},
			1,
		},
		{
			"Some elems",
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Len()
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name    string
		list    *Node[int]
		val     int
		idx     int
		wantErr error
	}{
		{
			"Empty list",
			nil,
			0,
			-1,
			EmptyList,
		},
		{
			"Not in list one item list",
			&Node[int]{Val: 1},
			0,
			-1,
			NotInList,
		},
		{
			"In list one item list",
			&Node[int]{Val: 1},
			1,
			0,
			nil,
		},
		{
			"Not in list some items list",
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
			0,
			-1,
			NotInList,
		},
		{
			"In list some item list",
			&Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}},
			2,
			1,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, idx := tt.list.Index(tt.val)
			if !assert.Equal(t, tt.wantErr, err) {
				t.Errorf("Len() error = %v, want %v", err, tt.wantErr)
			}

			if !assert.Equal(t, tt.idx, idx) {
				t.Errorf("Len() %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		list     *Node[int]
		wantList *Node[int]
	}{
		{
			"Empty list",
			nil,
			nil,
		},
		{
			"One node list",
			&Node[int]{Val: 0},
			&Node[int]{Val: 0},
		},
		{
			"Two node list",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1}},
			&Node[int]{Val: 1, Next: &Node[int]{Val: 0}},
		},
		{
			"Three node list",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2}}},
			&Node[int]{Val: 2, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 0}}},
		},
		{
			"Four node list",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3}}}},
			&Node[int]{Val: 3, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 0}}}},
		},
		{
			"Five node list",
			&Node[int]{Val: 0, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 3, Next: &Node[int]{Val: 4}}}}},
			&Node[int]{Val: 4, Next: &Node[int]{Val: 3, Next: &Node[int]{Val: 2, Next: &Node[int]{Val: 1, Next: &Node[int]{Val: 0}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Reverse()
			if tt.wantList != nil {
				currGot, currWant := got, tt.wantList
				for currGot != nil && currWant != nil {
					if !assert.Equal(t, currWant.Val, currGot.Val) {
						t.Errorf("List mismatch: got %v, want %v", currGot.Val, currWant.Val)
					}
					currGot, currWant = currGot.Next, currWant.Next
				}
				if currGot != nil || currWant != nil {
					t.Errorf("List lengths do not match")
				}
			}
		})
	}
}
