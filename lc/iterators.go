package main

import "fmt"


type MyCollection struct {
	data []int 
}

type Iterator struct {
	collection *MyCollection 
	index int 
}


func NewMyCollection(data []int) *MyCollection {
	return &MyCollection{data}
}

func (c *MyCollection) NewIterator() *Iterator {
	return &Iterator{collection: c, index: 0}
}

func (it *Iterator) HasNext() bool {
	if it.index < len(it.collection.data) {
		return true 
	}

	return false
}

func (it *Iterator) Next() (int, bool) { 
	fmt.Println("INDEX: ", it.index)
	if it.index < len(it.collection.data) {
		element := it.collection.data[it.index]
		it.index ++
		return element, true
	}
	return 0, false
}



func iteratorMain() {
	data := []int{1,2}

	collection := NewMyCollection(data) 
	iterator := collection.NewIterator()

	elem, ok := iterator.Next()
	if ok {
		fmt.Println(elem)
	} 

	elem, ok = iterator.Next()
	if ok {
		fmt.Println(elem)
	} 

	hasNext := iterator.HasNext()
	fmt.Println(hasNext)
}
