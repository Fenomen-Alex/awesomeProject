package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var tree1, tree2 []int

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	tr1, tr2 := 0, 0
	for i := 0; i < 20; i++ {
		select {
		case tr1 = <-ch1:
			tree1 = append(tree1, tr1)
		case tr2 = <-ch2:
			tree2 = append(tree2, tr2)
		}
	}
	return reflect.DeepEqual(tree1, tree2)
}

func main() {
	fmt.Println("Same(tree.New(1), tree.New(1))", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2))", Same(tree.New(1), tree.New(2)))
}
