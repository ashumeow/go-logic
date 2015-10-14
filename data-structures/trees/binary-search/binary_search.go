// binary search tree

package main

// importing packages
import "ashumeow/meow_tree"

// declaring types
type binary_search_tree struct {
	parent, left, right *binary_search_tree
	meowNode interface{}
	compare CompareFunc
}

// new binary search tree
func new_binary_search_tree(value interface{}, comp CompareFunc) *binary_search_tree {
	return &binary_search_tree{
		meowNode: value,
		compare: comp,
	}
}

// comparing 2 interfaces