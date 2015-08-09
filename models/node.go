package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Node interface {
	GetID() int
	GetChildren() map[int]*Node
	GetParent() *Node
	GetValue() int
}
type BNode struct {
	ID       int
	Children map[int]*Node
	Parent   *Node
	Value    int
	Left     *BNode
	Right    *BNode
}

func (n *BNode) GetID() int {
	return n.ID
}
func (n *BNode) GetChildren() map[int]*Node {
	return n.Children
}
func (n *BNode) GetParent() *Node {
	return n.Parent
}
func (n *BNode) GetValue() int {
	return n.Value
}
func (n *BNode) ToBTree() *BTree {
	return NewBTree(n)
}

func NewBNode(value int) Node {
	r := rand.New(rand.NewSource(int64(time.Now().Second())))
	n := BNode{
		ID:    r.Intn(1000),
		Value: value,
	}
	return &n
}

func (n *BNode) String() string {
	return fmt.Sprintf("id: %d, value: %d", n.ID, n.Value)
}
