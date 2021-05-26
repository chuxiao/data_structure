package tree

import (
	"strconv"
	"strings"
)

type ElementType int

func (e *ElementType) ToString() string {
	return strconv.Itoa(int(*e))
}
type BinaryTreeNode struct {
	e ElementType
	left, right *BinaryTreeNode
}

type BinarySearchTree struct {
	root *BinaryTreeNode
}

func (tree *BinarySearchTree) Insert(e ElementType) {
	n := &BinaryTreeNode{
		e: e,
	}
	if tree.root == nil {
		tree.root = n
		return
	}
	p := tree.root
	for {
		if p.e == e {
			return
		} else if p.e > e {
			if p.left != nil {
				p = p.left
			} else {
				p.left = n
				return
			}
		} else {
			if p.right != nil {
				p = p.right
			} else {
				p.right = n
				return
			}
		}
	}
}

func (tree *BinarySearchTree) Find(e ElementType) bool {
	p := tree.root
	for p != nil {
		if p.e == e {
			return true
		} else if p.e > e {
			p = p.left
		} else {
			p = p.right
		}
	}
	return false
}

func (tree *BinarySearchTree) Delete(e ElementType) {
	// 1. 空树
	if tree.root == nil {
		return
	}
	// 2. 只有一个结点且等于e的情况，删除后变成空树
	r := tree.root
	if r.e == e && r.left == nil && r.right == nil {
		tree.root = nil
		return
	}
	// 3. 在树中查找是否有等于e的节点
	var p, c *BinaryTreeNode
	p = r
	if tree.root.e > e {
		c = p.left
	} else {
		c = p.right
	}
	for c != nil {
		if c.e == e {
			break
		} else if c.e > e {
			p = c
			c = c.left
		} else {
			p = c
			c = c.right
		}
	}
	// 没有找到，直接返回
	if c == nil {
		return
	}
	// 没有右子树，删除即可
	if c.right == nil {
		if p.left == c {
			p.left = c.left
		} else {
			p.right = c.left
		}
		return
	}
	// 找到，用右子树最左结点
	p = c
	cr := c.right
	for cr.left != nil {
		p = cr
		cr = cr.left
	}
	c.e = cr.e
	if p.left == cr {
		p.left = cr.right
	} else {
		p.right = cr.right
	}
}

func (tree *BinarySearchTree) ToString() string {
	var path []*BinaryTreeNode
	ss := []string{}
	p := tree.root
	for p != nil || len(path) > 0 {
		if p != nil {
			if p.left != nil {
				path = append(path, p)
				p = p.left
			} else {
				ss = append(ss, p.e.ToString())
				p = p.right
			}
		} else {
			size := len(path)
			p = path[size - 1]
			path = path[:size - 1]
			ss = append(ss, p.e.ToString())
			p = p.right
		}
	}
	return strings.Join(ss, "    ")
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func BinarySearch(a []ElementType, e ElementType) int {
	i, j := 0, len(a) - 1
	for i <= j {
		c := (i + j) / 2
		t := &a[c]
		if *t == e {
			return c
		} else if *t < e {
			i = c + 1
		} else {
			j = c - 1
		}
	}
	return -1
}