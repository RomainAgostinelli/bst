package bst

import (
	"math/rand"
	"testing"
	"time"
)

type Point struct {
	p       int
	payload []int
}

func (point *Point) CompareTo(other Comparable) int {
	switch v := other.(type) {
	case *Point:
		if point.p < v.p {
			return -1
		} else if point.p > v.p {
			return 1
		} else {
			return 0
		}
	default:
		return -1
	}
}

func TestBST(t *testing.T) {
	// Creates data
	var points1 []Comparable
	// Create all pair points
	for i := 0; i < 500; i += 2 {
		rand.Seed(time.Now().UnixNano())
		points1 = append(points1, &Point{
			p:       i,
			payload: []int{rand.Intn(2000)},
		})
	}
	tree1 := NewBSTReady(points1[:])
	for i := 0; i < 500; i++ {
		contains := tree1.Contains(&Point{i, nil})
		if i%2 == 0 && !contains {
			t.Fatalf("DO NOT CONTAINS PAIR ELEMENT: %d", i)
		} else if i%2 == 1 && contains {
			t.Fatalf("CONTAINS UN PAIR ELEMENT: %d", i)
		}
	}

	tree2 := NewBST()
	for i := 0; i < 500; i++ {
		rand.Seed(time.Now().UnixNano())
		tree2.Add(&Point{
			p:       rand.Intn(10) + 1,
			payload: []int{rand.Intn(2000)},
		})
	}
	size := tree2.Size()
	if size > 10 { // must not be more than 10
		t.Fatalf("TREE CONTAINS %d ELEMENTS BUT MUST CONTAINS MAX 10", size)
	}

	tree2.Add(&Point{3, nil})
	size = tree2.Size()
	tree2.Remove(&Point{3, nil})
	if size <= tree2.Size() {
		t.Fatalf("REMOVING AN ELEMENT MUST DECREASE THE SIZE")
	}
}

func TestBST_IntervalSearch(t *testing.T) {
	// Creates data
	var points1 []Comparable
	// Create all point from 0 to 499
	for i := 0; i < 500; i ++ {
		rand.Seed(time.Now().UnixNano())
		points1 = append(points1, &Point{
			p:       i,
			payload: []int{rand.Intn(2000)},
		})
	}
	iteration := 200
	tree1 := NewBSTReady(points1[:])
	for i := 0; i < iteration; i++ {
		rand.Seed(time.Now().UnixNano())
		start := rand.Intn(250) // from 0 to 249
		rand.Seed(time.Now().UnixNano())
		end := rand.Intn(250) + 250  // from 250 to 499
		res:= tree1.IntervalSearch(&Point{start, nil}, &Point{end, nil})
		if len(res) != (end - start + 1) {
			t.Fatalf("EXPECTING %d ELEMENTS, RECEIVED %d", end - start + 1, len(res))
		}
	}
}