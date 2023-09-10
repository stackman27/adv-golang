package main

// Stacks and Queues
// STACK (last in first out)
// QUEUE (First in first out)

// stack represents stack that hold a slice
type Stack struct {
	items []int
}


// push  O(1)
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

// pop the value at the end and remove the removed value O(1)
func (s *Stack) pop() int{
	removedVal := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]

	return removedVal
}

type Queue struct {
	items []int
}


// Enque (add) O(1)
func (q *Queue) pushQueue(i int ) {
	q.items = append(q.items, i)
} 

// Deque (remove) O(1)
func (q *Queue) popQueue() int{
	removedValue := q.items[0]
	q.items = q.items[1:]

	return removedValue
}


// Monotonic Queue 
// DS where order of elements are strictly increasing or strictly decreasing, for ex: [1,3,5,6,7] or [5,4,2,1]
// Increasing Queue:  push a number that is less than the maximum so far
// For ex: we want an increasing queue with {5,3,1,2,4} 
// push(5) -> [5]
// push(3) -> remove 5 since this is increasing queue [3]
// push(1) -> remove 3 since this is increasing queue [1]
// push(2) -> add 2 to the list [1,2]
// push(4) -> add 4 to the list [1,2,4]
// if we want to push(3) -> list becomes [1,2,3]
// if we want to push(1) -> we remove everything and only add 1 [1]
