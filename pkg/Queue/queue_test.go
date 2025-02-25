package Queue

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func fillList(vals ...int) *list.List {
	l := list.New()
	for _, val := range vals {
		l.PushBack(val)
	}

	return l
}

func TestNew(t *testing.T) {
	got := New[int]()
	want := &Queue[int]{}
	if !assert.Equal(t, want, got) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}

func TestEnqueue(t *testing.T) {
	tests := []struct {
		name      string
		queue     *Queue[int]
		val       int
		wantQueue *Queue[int]
	}{
		{
			name:      "Enqueue to empty queue",
			queue:     &Queue[int]{},
			val:       1,
			wantQueue: &Queue[int]{list: fillList(1)},
		},
		{
			name:      "Enqueue to one elem queue",
			queue:     &Queue[int]{list: fillList(1)},
			val:       2,
			wantQueue: &Queue[int]{list: fillList(1, 2)},
		},
		{
			name:      "Enqueue to nil queue",
			queue:     &Queue[int]{list: nil},
			val:       1,
			wantQueue: &Queue[int]{list: fillList(1)},
		},
		{
			name:      "Enqueue to multiple elements queue",
			queue:     &Queue[int]{list: fillList(1, 2, 3)},
			val:       4,
			wantQueue: &Queue[int]{list: fillList(1, 2, 3, 4)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.queue.Enqueue(tt.val)
			if !assert.Equal(t, tt.wantQueue, tt.queue) {
				t.Errorf("Enqueue() = %v, want %v", tt.queue, tt.wantQueue)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		name      string
		queue     *Queue[int]
		wantErr   error
		wantVal   int
		wantQueue *Queue[int]
	}{
		{
			name:      "Dequeue from empty queue",
			queue:     &Queue[int]{},
			wantErr:   EmptyQueue,
			wantVal:   0,
			wantQueue: &Queue[int]{},
		},
		{
			name:      "Dequeue from one elem queue",
			queue:     &Queue[int]{list: fillList(1)},
			wantErr:   nil,
			wantVal:   1,
			wantQueue: &Queue[int]{list: nil},
		},
		{
			name:      "Dequeue from multiple elements queue",
			queue:     &Queue[int]{list: fillList(1, 2, 3)},
			wantErr:   nil,
			wantVal:   1,
			wantQueue: &Queue[int]{list: fillList(2, 3)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, val := tt.queue.Dequeue()
			if !assert.Equal(t, tt.wantErr, err) {
				t.Errorf("Dequeue() error = %v, want %v", err, tt.wantErr)
			}
			if !assert.Equal(t, tt.wantVal, val) {
				t.Errorf("Dequeue() val = %v, want %v", val, tt.wantVal)
			}
		})
	}
}

func TestFront(t *testing.T) {
	tests := []struct {
		name    string
		queue   *Queue[int]
		wantVal int
		wantErr error
	}{
		{
			name:    "Front of empty queue",
			queue:   &Queue[int]{},
			wantVal: 0,
			wantErr: EmptyQueue,
		},
		{
			name:    "Front of one elem queue",
			queue:   &Queue[int]{list: fillList(1)},
			wantVal: 1,
			wantErr: nil,
		},
		{
			name:    "Front of multiple elements queue",
			queue:   &Queue[int]{list: fillList(1, 2, 3)},
			wantVal: 1,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := tt.queue.Front()
			if !assert.Equal(t, tt.wantErr, err) {
				t.Errorf("Front() error = %v, want %v", err, tt.wantErr)
			}

			if !assert.Equal(t, tt.wantVal, got) {
				t.Errorf("Front() = %v, want %v", got, tt.wantVal)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		queue *Queue[int]
		want  bool
	}{
		{
			name:  "Empty queue",
			queue: &Queue[int]{},
			want:  true,
		},
		{
			name:  "One elem queue",
			queue: &Queue[int]{list: fillList(1)},
			want:  false,
		},
		{
			name:  "Multiple elements queue",
			queue: &Queue[int]{list: fillList(1, 2, 3)},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.queue.IsEmpty()
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		name     string
		queue    *Queue[int]
		wantSize int
	}{
		{
			name:     "Empty queue",
			queue:    &Queue[int]{},
			wantSize: 0,
		},
		{
			name:     "One elem queue",
			queue:    &Queue[int]{list: fillList(1)},
			wantSize: 1,
		},
		{
			name:     "Multiple elements queue",
			queue:    &Queue[int]{list: fillList(1, 2, 3)},
			wantSize: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.queue.Size()
			if !assert.Equal(t, tt.wantSize, got) {
				t.Errorf("Size() = %v, want %v", got, tt.wantSize)
			}
		})
	}
}
