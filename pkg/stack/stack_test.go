package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name      string
		stack     *Stack[int]
		elem      int
		wantStack *Stack[int]
	}{
		{
			"Push to empty stack",
			&Stack[int]{},
			0,
			&Stack[int]{data: []int{0}},
		},
		{
			"Push to one elem stack",
			&Stack[int]{[]int{0}},
			1,
			&Stack[int]{data: []int{0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push(tt.elem)
			if !assert.Equal(t, tt.wantStack, tt.stack) {
				t.Errorf("Push() = %v, want %v", tt.stack, tt.wantStack)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name      string
		stack     *Stack[int]
		wantStack *Stack[int]
		wantErr   error
		wantElem  int
	}{
		{
			"Pop from empty stack",
			&Stack[int]{},
			&Stack[int]{},
			EmptyStack,
			0,
		},
		{
			"Pop from one elem stack",
			&Stack[int]{data: []int{1}},
			&Stack[int]{data: []int{}},
			nil,
			1,
		},
		{
			"Pop from two elem stack",
			&Stack[int]{data: []int{1, 2}},
			&Stack[int]{data: []int{1}},
			nil,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, elem := tt.stack.Pop()

			if !assert.Equal(t, tt.wantErr, err) {
				t.Errorf("Pop() error = %v, want %v", err, tt.wantErr)
			}

			if !assert.Equal(t, tt.wantStack, tt.stack) {
				t.Errorf("Pop() = %v, want %v", tt.stack, tt.wantStack)
			}

			if !assert.Equal(t, tt.wantElem, elem) {
				t.Errorf("Pop() elem = %v, want %v", elem, tt.wantElem)
			}
		})
	}
}

func TestTop(t *testing.T) {
	tests := []struct {
		name      string
		stack     *Stack[int]
		wantStack *Stack[int]
		wantErr   error
		wantElem  int
	}{
		{
			"Top from empty stack",
			&Stack[int]{},
			&Stack[int]{},
			EmptyStack,
			0,
		},
		{
			"Top from one elem stack",
			&Stack[int]{data: []int{1}},
			&Stack[int]{data: []int{1}},
			nil,
			1,
		},
		{
			"Top from two elem stack",
			&Stack[int]{data: []int{1, 2}},
			&Stack[int]{data: []int{1, 2}},
			nil,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, elem := tt.stack.Top()

			if !assert.Equal(t, tt.wantErr, err) {
				t.Errorf("Top() error = %v, want %v", err, tt.wantErr)
			}

			if !assert.Equal(t, tt.wantStack, tt.stack) {
				t.Errorf("Top() = %v, want %v", tt.stack, tt.wantStack)
			}

			if !assert.Equal(t, tt.wantElem, elem) {
				t.Errorf("Top() elem = %v, want %v", elem, tt.wantElem)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack[int]
		want  bool
	}{
		{
			"IsEmpty empty stack",
			&Stack[int]{},
			true,
		},
		{
			"IsEmpty non empty stack",
			&Stack[int]{data: []int{1}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stack.IsEmpty()

			if !assert.Equal(t, tt.want, got) {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack[int]
		want  int
	}{
		{
			"Size of empty stack",
			&Stack[int]{},
			0,
		},
		{
			"Size of one elem stack",
			&Stack[int]{data: []int{1}},
			1,
		},
		{
			"Size of two elem stack",
			&Stack[int]{data: []int{1, 2}},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stack.Size()

			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		name      string
		stack     *Stack[int]
		wantStack *Stack[int]
	}{
		{
			"Clear empty stack",
			&Stack[int]{},
			&Stack[int]{data: []int{}},
		},
		{
			"Clear one elem stack",
			&Stack[int]{data: []int{1}},
			&Stack[int]{data: []int{}},
		},
		{
			"Clear two elem stack",
			&Stack[int]{data: []int{1, 2}},
			&Stack[int]{data: []int{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Clear()
			if !assert.Equal(t, tt.wantStack, tt.stack) {
				t.Errorf("Clear() = %v, want %v", tt.stack, tt.wantStack)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name  string
		stack *Stack[int]
		elem  int
		want  bool
	}{
		{
			"Contains empty stack",
			&Stack[int]{},
			0,
			false,
		},
		{
			"Contains one elem stack",
			&Stack[int]{data: []int{1}},
			1,
			true,
		},
		{
			"Not contains one elem stack",
			&Stack[int]{data: []int{1}},
			0,
			false,
		},
		{
			"Contains two elem stack",
			&Stack[int]{data: []int{1, 2}},
			2,
			true,
		},
		{
			"Not contains two elem stack",
			&Stack[int]{data: []int{1, 2}},
			0,
			false,
		},
		{
			"Contains three elem stack",
			&Stack[int]{data: []int{1, 2, 3}},
			2,
			true,
		},
		{
			"Not contains three elem stack",
			&Stack[int]{data: []int{1, 2, 3}},
			0,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stack.Contains(tt.elem)
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
