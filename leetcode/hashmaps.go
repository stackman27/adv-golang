package main

import "fmt"

// https://www.youtube.com/watch?v=zLnJcAt1aKs
// HASHTABLES insert, search, delete O(1)
// - Storing and searching through array is expensive O(n)
// - instead we store the index of the array we store
// - how do we determine the index? using hash function
// - input "arryVal" into a hashFunction output "hashCode". (HASHCODE WILL BE THE LOCATION OF RANDY)
//
// - When we are searching for "arrayVal" we put it in hashFunction and it outputs HASHCODE!! YAYY
// - Data Structure that stores data this way is call HashTable
//
// COLLISONS
// - two array value can have same hashCode
// - How to handle collisions?
// - Open Addressing (store value in next index, while searching check the hashCode index if not found check the next index)
// - Seperate Chaining (storing multiple names in 1 index by using linked list)
//
 
const ArraySize = 7 
// HashTable Structure 
type HashTable struct {
	array [ArraySize]*Bucket
}

// Insert 
func (h *HashTable) Insert(key string) {
	index := hash(key) 
	h.array[index].insert(key)
	
}

// Search 
func (h *HashTable) Search(key string) bool{
	index := hash(key) 
	return h.array[index].search(key) 
}

// Delete 
func (h *HashTable) Delete(key string) {
	index := hash(key) 
	h.array[index].delete(key)
}

func hash(key string) int{
	// get asci code for each character 
	// sum it up and divide by the size of array 
	sum := 0 
	for _, c := range key {
		sum = sum + int(c) 
	}	

	return sum % ArraySize 
}

// Bucket Structure (AKA linked list)
type Bucket struct {
	head *bucketNode
}

// BucketNode structure 
type bucketNode struct {
	key string 
	next *bucketNode
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *Bucket) insert(key string) {
	if !b.search(key) {
	newNode := bucketNode{key: key}
	newNode.next = b.head // this may not be necessary
	b.head = &newNode // this is the head node (always point to the new value)
	} else {
		// insert new node 
		fmt.Printf("%s already exist", key)
		return 
	}
}

// search will take in a key and return true if the bucket has the key 
func (b *Bucket) search(key string) bool {
	currentNode := b.head 

	for currentNode != nil {
		if currentNode.key == key {
			return true 
		}
		currentNode = currentNode.next
	}
	return false
} 

// delete
func (b *Bucket) delete(key string) {
	previousNode := b.head 
	
	if b.head.key == key {
		b.head = b.head.next
		return 
	}
	
	for previousNode.next != nil {
		// loop it through and update it 
		if previousNode.next.key == key {
			// delete 
			previousNode.next = previousNode.next.next
		}
		 previousNode = previousNode.next
		
	}
}


// hashFunction 

// InitHashTable
func initHashTable() *HashTable{
	newHashTable := &HashTable{}


	for i := range newHashTable.array {
		newHashTable.array[i] = &Bucket{}
	}

	return newHashTable
}

func init() {
	testHashTable :=initHashTable()
	fmt.Println(testHashTable)


	testBucket := &Bucket{}
	testBucket.insert("SISHIR")
	testBucket.delete("SISHIR")
	exist := testBucket.search("SISHIR")
	fmt.Println(exist)


}
 