package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Node defines the structure of a node in a binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode creates a new node with the given value
func NewNode(value int) *Node {
	return &Node{Value: value, Left: nil, Right: nil}
}

// BinarySearchTree represents the binary search tree
type BinarySearchTree struct {
	Root *Node
}

// NewBinarySearchTree creates a new binary search tree
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{Root: nil}
}

// Insert inserts a new value into the binary search tree
func (bst *BinarySearchTree) Insert(value int) {
	newNode := NewNode(value)
	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	current := bst.Root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

// Find searches for a value in the binary search tree
func (bst *BinarySearchTree) Find(value int) bool {
	current := bst.Root
	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

// Remove removes a value from the binary search tree
func (bst *BinarySearchTree) Remove(value int) {
	bst.Root = removeNode(bst.Root, value)
}

func removeNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = removeNode(node.Left, value)
	} else if value > node.Value {
		node.Right = removeNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = removeNode(node.Right, minNode.Value)
	}
	return node
}

func findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func inOrderTraversal(node *Node, values *[]int) {
	if node != nil {
		inOrderTraversal(node.Left, values)
		*values = append(*values, node.Value)
		inOrderTraversal(node.Right, values)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bst := NewBinarySearchTree()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)

	e.GET("/search/:value", func(c echo.Context) error {
		value := c.Param("value")
		valueInt := 0
		fmt.Sscanf(value, "%d", &valueInt)
		found := bst.Find(valueInt)
		return c.JSON(200, map[string]bool{"found": found})
	})

	e.DELETE("/remove/:value", func(c echo.Context) error {
		value := c.Param("value")
		valueInt := 0
		fmt.Sscanf(value, "%d", &valueInt)
		bst.Remove(valueInt)
		return c.String(200, "Value removed from the tree")
	})

	e.GET("/traversal", func(c echo.Context) error {
		var values []int
		inOrderTraversal(bst.Root, &values)
		return c.JSON(200, values)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

