package bst

import (
	"errors"
	"github.com/ag0st/binarytree"
	"log"
)

// BST struct used to store a binary search tree
type BST struct {
	tree    *binarytree.BinaryTree
	crtSize int
}

type Comparable interface {
	// CompareTo returns "< 0" if this > other; ">0" if this > other; "0" if this == other
	CompareTo(other Comparable) int
}

// NewBSTReady creates new Binary Search Tree with values in sorted.
// Method to use when the data are ready: sorted and no duplicates
// PRE: sorted, no duplicate
func NewBSTReady(sorted []Comparable) *BST {
	return &BST{
		tree:    optimalBST(sorted, 0, len(sorted)-1),
		crtSize: len(sorted),
	}
}

// NewBST creates an empty Binary Search Tree
func NewBST() *BST {
	return &BST{
		tree:    &binarytree.BinaryTree{},
		crtSize: 0,
	}
}

// Add a new element in the tree if not present (use CompareTo method)
func (bst *BST) Add(e Comparable) {
	itr := bst.locate(e)
	if !itr.IsBottom() {
		return // Already present in the structure
	}
	itr.Insert(e)
	bst.crtSize++
}

// Remove remove the element if present (use CompareTo method)
func (bst *BST) Remove(e Comparable) {
	itr := bst.locate(e)
	if itr.IsBottom() {
		return // nothing to remove, do not exists
	}
	bst.crtSize--
	// Push the element as leaf with rotation to keep order
	for itr.HasRight() {
		itr.RotateLeft()
		itr = itr.Left()
	}
	// remove the element and paste subtree if not leaf
	if !itr.IsLeaf() {
		l := itr.Left().Cut()
		itr.Cut()
		err := itr.Paste(l)
		if err != nil {
			log.Fatalln("CANNOT CUT AND PASTE WHEN REMOVING")
		}
	} else {
		itr.Cut()
	}
}

// Contains tells if the BST contains the element (use CompareTo)
func (bst *BST) Contains(e Comparable) bool {
	itr := bst.locate(e)
	return !itr.IsBottom()
}

// Size returns the size of the tree
func (bst *BST) Size() int {
	return bst.crtSize
}

// IntervalSearch returns all elements in the interval [min, max] (inclusive)
func (bst *BST) IntervalSearch(min, max Comparable) []Comparable {
	return intervalSearch(bst.tree.Root(), min, max)
}

func (bst *BST) Get(e Comparable) (Comparable, error) {
	itr := bst.locate(e)
	if itr.IsBottom() {
		return nil, errors.New("element not present, please use Contains method before")
	}
	return itr.Consult().(Comparable), nil
}

func (bst *BST) GetPredSucc(e Comparable) (pred, ele, succ Comparable) {
	itr := bst.tree.Root()
	for !itr.IsBottom() {
		current := itr.Consult().(Comparable)
		if e.CompareTo(current) == 0 {
			if itr.HasRight() {
				succ = itr.Right().LeftMost().Up().Consult().(Comparable)
			}
			if itr.HasLeft() {
				pred = itr.Left().RightMost().Up().Consult().(Comparable)
			}
			ele = itr.Consult().(Comparable)
			break
		} else if e.CompareTo(current) < 0 {
			succ = itr.Consult().(Comparable)
			itr = itr.Left()
		} else {
			pred = itr.Consult().(Comparable)
			itr = itr.Right()
		}
	}

	return pred, ele, succ
}

// intervalSearch returns all elements in the interval [min, max] (inclusive) from the iterator position
func intervalSearch(itr *binarytree.Iterator, min, max Comparable) []Comparable {
	var res []Comparable
	if itr.IsBottom() {
		return res
	}
	current := itr.Consult().(Comparable)
	if current.CompareTo(min) >= 0 && current.CompareTo(max) <= 0 {
		res = append(res, current)
	}
	if itr.IsLeaf() { // no more left or right subtree
		return res
	}
	if current.CompareTo(max) < 0 { // if current < max --> check right
		res = append(res, intervalSearch(itr.Right(), min, max)...)
	}
	if current.CompareTo(min) > 0 { // if current > min --> check left
		res = append(res, intervalSearch(itr.Left(), min, max)...)
	}
	return res
}

// locate returns the position where the comparable is or where it must be added if not present
func (bst *BST) locate(e Comparable) *binarytree.Iterator {
	itr := bst.tree.Root()
	for !itr.IsBottom() {
		current := itr.Consult().(Comparable)
		if e.CompareTo(current) == 0 {
			break
		} else if e.CompareTo(current) < 0 {
			itr = itr.Left()
		} else {
			itr = itr.Right()
		}
	}
	return itr
}

// optimalBST creates a balanced BST with the sorted array given in parameter using the slice between left and right
func optimalBST(sorted []Comparable, left, right int) *binarytree.BinaryTree {
	tree := &binarytree.BinaryTree{}
	itr := tree.Root()
	if left > right {
		return tree
	}
	mid := (left + right) / 2
	itr.Insert(sorted[mid])
	err := itr.Left().Paste(optimalBST(sorted, left, mid-1))
	if err != nil {
		log.Fatalln("OPTIMAL BST: Cannot paste left")
	}
	err = itr.Right().Paste(optimalBST(sorted, mid+1, right))
	if err != nil {
		log.Fatalln("OPTIMAL BST: Cannot paste right")
	}

	return tree
}
