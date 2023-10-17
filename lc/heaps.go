package main

import "fmt"

func testHeapImpl() {
	m := &MaxHeap{}
	
	vals := []int{10,2,3,4,4,1,5,23,23,5,2,12,4}
	for _, val := range vals {
		m.Insert(val)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}

// https://www.youtube.com/watch?v=3DYIgTC4T1o
// HEAP Data structure
// introduced as HEAPSORT
// Particulary good for implementing priorityQueue
// NormalQueue: FIFO
// PriorityQueue: We take out the max number
// Particularly useful when we have to remove the largest key in the DS (for ex: remove largest value in BS)
//
// Minheap: Root is the smallet key
// - Each parent node will have smaller node than its children
//
// Max heap implementation (stored like an array where each node is indexed)
// Therefore heap is actually an array that operates like a tree
//
// Heapify - process of making parent key larger than its children

// TODOs
// Struct for max-heap
type MaxHeap struct {
	array []int
}

// insert heap method 
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key) 

	// rearrange so it can maintain heap properties  
	h.maxHeapifyUp(len(h.array) - 1)
}

// extract method extracts the largest key and removes it from the heap
// AKA removing the largest value from the heap 
func (h *MaxHeap) Extract() int {
	extracted := h.array[0] 

	l := len(h.array) - 1

	// if array is empty return nill 
	if len(h.array) == 0 {
		fmt.Println("EMPTY ARRAY")
		return -1
	}

	// take out the last index and put it in root 
	h.array[0] = h.array[l]
	h.array = h.array[:l] // get everything except the last element 

	h.maxHeapifyDown(0) 


	return extracted	
}

// maxHeapifyDown will heapify top to bottom 
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1 
	l, r := left(index), right(index) 

	childToCompare := 0
 
	// loop while index has at least one child 
	for l < lastIndex {
		if l == lastIndex {
			// when left child if the only child
			childToCompare = l  
		} else if h.array[l] > h.array[r] {
			// when left child is greater
			childToCompare = l
		} else {
			// right side is greater 
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller 
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare) 
			index = childToCompare
			l, r = left(index), right(index) 
		} else {
			return
		}
	}

}

// Max heapify will heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int) {
	// swap when the current index is larger than parent 
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index) , index) 
		index = parent(index)
	}
}

func parent(i int) int {
	// to get the parent index we do (val - 1)/2 
	// left index is always odd and right is even, so we round down
	return (i - 1) / 2
}

// get the left  index
func left(i int) int {
	return (2 * i) + 1
}

// get the right index
func right(i int) int {
	return (2 * i) + 2
}

// swap keys in the array 
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}