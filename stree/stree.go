package stree

type Node struct {
	Key int
	Value interface{}
	Left *Node
	Right *Node
	Parent *Node
}

type Tree struct {
	Size int
	Root *Node
}

func (tree *Tree) Search(key int) *Node {
	return search(tree.Root, key)
}

func (tree *Tree) Min() *Node{
	return min(tree.Root)
}

func (tree *Tree) Max() *Node{
	return max(tree.Root)
}

func (tree *Tree) Sort() []*Node{
	result := make([]*Node, 0)
	sort(tree.Root, &result)
	return result
}

func (tree *Tree) Add(key int, value interface{}) {
	node := &Node{
		Key: key,
		Value: value,
		Left: nil,
		Right: nil,
	}
	if tree.Root == nil {
		tree.Root = node
		tree.Size++
		return
	}
	add(tree, tree.Root, node)
}

func (tree *Tree) Delete(key int) {
	node := tree.Search(key)
	if node == nil {
		return
	}
	// no child
	if node.Right == nil && node.Left == nil {
		if node.Parent.Key < node.Key {
			node.Parent.Right = nil
			return
		}
		node.Parent.Left = nil
		return
	}
	// two child
	if node.Right != nil && node.Left != nil {
		predecessor := predecessor(node)
		//detach predecessor
		if predecessor.Key > predecessor.Parent.Key {
			if predecessor.Left != nil {
				predecessor.Parent.Right = predecessor.Left
				predecessor.Left.Parent = predecessor.Parent
			} else {
				predecessor.Parent.Right = nil
			}
		} else {
			if predecessor.Left != nil {
				predecessor.Parent.Left = predecessor.Left
				predecessor.Left.Parent = predecessor.Parent
			} else {
				predecessor.Parent.Left = nil
			}
		}
		//swap current node and predecessor
		//TODO swap pointer you lazy fuck!
		node.Key = predecessor.Key
		node.Value = predecessor.Value
		tree.Size--
		return
	}
	// only one child
	if node.Parent == nil {
		if node.Right != nil {
			node.Right.Parent = nil
			tree.Root = node.Right
		}
		if node.Left != nil {
			node.Left.Parent = nil
			tree.Root = node.Left
		}
		tree.Size--
		return
	}
	if node.Key > node.Parent.Key {
		if node.Right != nil {
			node.Parent.Right = node.Right
			node.Right.Parent = node.Parent
		}
		if node.Left != nil {
			node.Parent.Right = node.Left
			node.Left.Parent = node.Parent
		}
		tree.Size--
		return
	}
	if node.Right != nil {
		node.Parent.Left = node.Right
		node.Right.Parent = node.Parent
	}
	if node.Left != nil {
		node.Parent.Left = node.Left
		node.Left.Parent = node.Parent
	}
	tree.Size--
	return
}

func search(node *Node, key int) *Node  {
	if key == node.Key {
		return node
	}
	if key < node.Key {
		if node.Left == nil {
			return nil
		}
		return search(node.Left, key)
	}
	if node.Right == nil {
		return nil
	}
	return search(node.Right, key)
}

func min(node *Node) *Node {
	if node.Left != nil {
		return min(node.Left)
	}
	return node
}

func max(node *Node) *Node {
	if node.Right != nil {
		return max(node.Right)
	}
	return node
}

func predecessor(node *Node) *Node {
	if node.Left != nil {
		return max(node.Left)
	}
	if node.Parent != nil {
		if node.Parent.Key < node.Key {
			return node.Parent
		}
		if node.Parent.Parent != nil {
			if node.Parent.Parent.Key < node.Key {
				return node.Parent.Parent
			}
		}
		return nil
	}
	return nil
}

func sort(node *Node, result *[]*Node) {
	if node.Left != nil {
		sort(node.Left, result)
	}
	*result = append(*result, node)
	if node.Right != nil {
		sort(node.Right, result)
	}
}

func add(tree *Tree,node *Node, newNode *Node) {
	if newNode.Key == node.Key {
		return
	}
	if newNode.Key < node.Key {
		if node.Left == nil {
			newNode.Parent = node
			node.Left = newNode
			tree.Size++
			return
		}
		add(tree, node.Left, newNode)
		return
	}
	if node.Right == nil {
		newNode.Parent = node
		node.Right = newNode
		tree.Size++
		return
	}
	add(tree, node.Right, newNode)
	return
}

