package main

import (
	"fmt"
	"math"
)

// Node represents the component of the binary tree
type Node struct {
	key int 
	left *Node 
	right *Node  
}

 
// Binary Search 
//    	  10  - root
//    	 / \
//   	9   11 - (parent of 5,15 and child of 10)
//     /     \         
//    5      15 - (leaves)
// 
// Runtime 
// Search = (log n) 
// Q. Does it handle duplicates? - easy answer no, find out why?
 
// Insert = log(n) because of the recursive function
func (n *Node) insert(key int) {
	if key == n.key {
		fmt.Println("Sorry we donot allow duplicates")
		return 
	}

 
	if key > n.key  {
		// move to the right 
		// if the child is empty just add the node 
		if n.right == nil {
			n.right = &Node{key:key}
		} else {
			// recursively traverse through the right branch 
			n.right.insert(key)
		}
	}
	
	if key < n.key {
		// we want the node to go to left 
		// if the child is empty just add the node 
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			// recursively travese through the left branch 
			n.left.insert(key);
		}
	} 

	return 
}

// Search in log(n)  
func (n *Node) Search(key int) bool{
	// there are no values (end of the tree )
	if n == nil {
		return false;
	}

	if key < n.key {
		// go left of the tree 
		return n.left.Search(key);
	}

	if key > n.key {
		// go right of the tree 
		return n.right.Search(key);
	}

	return true 
}
 
// check if a given tree is BST or not 
// function that can be used to check if a binary tree is a BST or not. The logic behind this is to keep track of the minimum and maximum values a node can take. 
// And at each node we will check if its value is between the min and max values it’s allowed to take. The root can take any value between negative infinity and positive infinity. 
// Here we are taking INT_MIN and INT_MAX for simplicity. 
func isValidBST(node *Node) bool{
	return validate(node, math.MinInt64, math.MaxInt64)
}

func validate(node *Node, minValue int, maxValue int) bool{
	if node == nil {
		return true 
	}
	
	if node.key < minValue || node.key > maxValue {
		return false 
	}

	return validate(node.left, minValue, node.key) && validate(node.right, node.key, maxValue)
}


// ******** BREADTH FIRST SEARCH ********** 
// TODO
func (n *Node) LevelOrderTraversal() {

}


// ******** DEPTH FIRST SEARCH **************
// In Order 

// For Inorder, you traverse from the left subtree to the root then to the right subtree.  
func (n *Node) InOrderTraversal() {
	if n != nil {
		n.left.InOrderTraversal()
		fmt.Printf("%d ", n.key)
		n.right.InOrderTraversal()
	}
}


// Pre Order 

// For Preorder, you traverse from the root to the left subtree then to the right subtree.  
func (n *Node) PreOrderTraversal() {
	if n != nil {
		fmt.Printf("%d ", n.key)
		n.left.PreOrderTraversal()
		n.right.PreOrderTraversal()
	}
}


// Post Order 

// For Post order, you traverse from the left subtree to the right subtree then to the root.
func (n *Node) PostOrderTraversal() {
	if n!= nil {
		n.left.PostOrderTraversal()
		n.right.PostOrderTraversal()
		fmt.Printf("%d", n.key)
	}
}
 
func binarySearch() {
	newNode := Node {key: 10} // root node 
	newNode.insert(12)
	newNode.insert(2)
	newNode.insert(5)


	fmt.Println(newNode)


	nodeExist := newNode.Search(5);

	fmt.Println(nodeExist)

	newNode.InOrderTraversal()

	fmt.Println("")
	value:= isValidBST(&newNode)

	fmt.Println(value)

}


 