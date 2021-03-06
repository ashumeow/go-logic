// binary search tree

package main

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
type CompareFunc func(x1, x5 interface{}) int

// searching
func(ig *binary_search_tree) meowSearch(x interface{}) interface{} {
	meowDir := ig.compare(ig.meowNode, x)

	switch meowDir {
	case 0:
		return ig.meowNode
	case -1:
		if ig.right==nil {
			return nil
		} else {
			return ig.right.meowSearch(x)
		}
	case 1:
		if ig.left==nil {
			return nil
		} else {
			return ig.left.meowSearch(x)
		}
	default:
		panic("invalid compare function")
	}
}

// checks for key existence
func (ig *binary_search_tree) meowExists(meowKey interface{}) bool {
	return ig.meowSearch(meowKey) != nil
}

// check if empty
func (ig *binary_search_tree) meowEmpty() bool {
	return ig.meowNode == nil
}

// Insertion... adding a key and value pair
func (ig *binary_search_tree) meowInsert(value interface{}) {
	if ig.meowNode==nil {
		ig.meowNode=value
		return
	}
	meowDir := ig.compare(ig.meowNode, value)

	switch meowDir {
	case 0:
		ig.meowNode=value
	case 1:
		if ig.left==nil {
			ig.left=new_binary_search_tree(value,ig.compare)
			ig.left.parent=ig
		} else {
			ig.left.meowInsert(value)
		}
	case -1:
		if ig.right==nil {
			ig.right=new_binary_search_tree(value,ig.compare)
			ig.right.parent=ig
		} else {
			ig.right.meowInsert(value)
		}
	default:
		panic("invalid compare function")
	}
}
func main() {}