package heap

import (
	"fmt"
)

// create, add, remove, print, print_indent

type Ordered interface {
	~float64 | ~float32 | ~int
}

type Heap[T Ordered] struct {
	items []T
}

func (heap *Heap[T]) swap(index1, index2 int) {
	heap.items[index1], heap.items[index2] = heap.items[index2], heap.items[index1]
}

func (heap *Heap[T]) buildHeap(index int) {
	var parent int
	if index > 0 {
		parent = (index - 1) / 2
		if heap.items[index] > heap.items[parent] {
			heap.swap(index, parent)
		}
		heap.buildHeap(parent)
	}
}

func (heap *Heap[T]) rebuildHeap(index int) {
	length := len(heap.items)

	if (2*index + 1) < length {
		left := 2*index + 1
		right := 2*index + 2
		larg := index

		if left < length && right < length && heap.items[left] >= heap.items[right] &&
			heap.items[index] < heap.items[left] {
			larg = left
		} else if right < length && heap.items[right] >= heap.items[left] &&
			heap.items[index] < heap.items[right] {
			larg = right
		} else if left < length && right >= length && heap.items[index] < heap.items[left] {
			larg = left
		}

		if index != larg {
			heap.swap(index, larg)
			heap.rebuildHeap(larg)
		}
	}
}

func (heap *Heap[T]) largest() T {
	return heap.items[0]
}

func create[T Ordered](input []T) *Heap[T] {
	heap := &Heap[T]{}
	for i := 0; i < len(input); i++ {
		heap.add(input[i])
	}
	return heap
}

func (heap *Heap[T]) add(value T) {
	heap.items = append(heap.items, value)
	heap.buildHeap(len(heap.items) - 1)
}

func (heap *Heap[T]) remove() {
	//remove the greatest value - in items[0]
	heap.items[0] = heap.items[len(heap.items)-1]
	heap.items = heap.items[:(len(heap.items) - 1)]
	heap.rebuildHeap(0)
}

//func reverse_print(heap *Heap) {
//if !intToBool(empty(heap)) {
//	reverse_print(heap.next)
//	fmt.Println(heap.value)
//}
//}

func (heap *Heap[T]) print() {
	fmt.Println(heap.items)
}

/* func (heap *Heap[T]) recursive_print(max int) {
	if len(heap.items) > 0 {
		fmt.Println(heap.value)
		if heap.next != nil {
			recursive_print(heap.next)
		}
	}
} */

func main() {
	values := []int{20, 51, 63, 42, 19}
	heap := create(values)

	heap.add(57)
	heap.print()

	heap.remove()
	heap.print()
}
