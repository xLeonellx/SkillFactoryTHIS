package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func main() {

	t := &TreeNode{val: 8}

	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)

	t.Find(11)

	t.Delete(5)
	t.Delete(7)

	t.PrintInorder()
	fmt.Println("")

	fmt.Println("min is", t.FindMin())
	fmt.Println("max is", t.FindMax())
}

func (t *TreeNode) PrintInorder() {

	if t == nil {
		return
	}

	t.left.PrintInorder()
	fmt.Print(t.val)
	t.right.PrintInorder()
}

func (t *TreeNode) Insert(value int) error {

	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.val == value {
		return errors.New("This node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}

	return nil
}

func (t *TreeNode) Find(value int) (TreeNode, bool) {

	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case value == t.val:
		return *t, true
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (t *TreeNode) remove(value int) *TreeNode {

	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}
	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

func (t *TreeNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}