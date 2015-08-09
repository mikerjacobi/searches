package models

import (
	"errors"
	"fmt"
)

type Tree interface {
	GetRoot() Node
	AddNode(*Node) int
}

type BTree struct {
	Root *Node
}

func BSearch(t *BTree, value int) (bool, *BNode, error) {
	root := t.GetRoot()
	r, ok := root.(*BNode)
	if !ok {
		return false, nil, errors.New("can't btree bfs on non btree")
	}

	if r.Value == value {
		return true, r, nil
	} else if value < r.Value {
		if r.Left == nil {
			return false, r, nil
		}
		return BSearch(r.Left.ToBTree(), value)
	} else {
		if r.Right == nil {
			return false, r, nil
		}
		return BSearch(r.Right.ToBTree(), value)
	}
}

func (bt *BTree) GetRoot() Node {
	return *bt.Root
}
func (bt *BTree) AddNode(newNode *BNode) error {
	found, retNode, err := BSearch(bt, newNode.GetValue())
	if err != nil {
		return err
	}
	if found == true {
		return fmt.Errorf("value %d already added", newNode.GetValue())
	}

	if newNode.GetValue() < retNode.GetValue() {
		retNode.Left = newNode
	} else {
		retNode.Right = newNode
	}
	return nil
}

func NewBTree(root Node) *BTree {
	return &BTree{
		Root: &root,
	}
}
