package main
''
import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}
type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value}
	} else {
		t.Root.insertNode(value)
	}
}

func (t *BinaryTree) Remove(value int) {
	t.Root = removeNode(t.Root, value)
}

func removeNode(n *Node, value int) *Node {
	if n == nil {
		return nil
	}
	if value < n.Value {
		n.Left = removeNode(n.Left, value)
	} else if value > n.Value {
		n.Right = removeNode(n.Right, value)
	} else {
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}
		minRight := findMin(n.Right)
		n.Value = minRight.Value
		n.Right = removeNode(n.Right, minRight.Value)
	}
	return n
}

func findMin(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (n *Node) insertNode(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.insertNode(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.insertNode(value)
		}
	}
}

func inOrder(n *Node) {
	if n != nil {
		inOrder(n.Left)
		fmt.Printf("%d ", n.Value)
		inOrder(n.Right)
	}
}

func (t *BinaryTree) inOrder() {
	inOrder(t.Root)
	fmt.Println()
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(16)
	fmt.Println("Arvore apos a inserção")
	tree.inOrder()
	tree.Remove(10)
	fmt.Println("Arvore apos a Remoção")
	tree.inOrder()
	tree.Remove(5)
	fmt.Println("Arvore apos a Remoção")
	tree.inOrder()
}
