package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

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
// func Same(t1, t2 *tree.Tree) bool

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	t1 := tree.New(1)
	fmt.Println("Tree 1:")
	PrintBTreePretty(t1, "", false)
	go Walk(t1, ch1)
	go Walk(tree.New(1), ch2)

	for i := 0; i < 10; i++ {
		x := <-ch1
		y := <-ch2
		// fmt.Println(x, y)
		if x != y {
			println("Not the same")
			return
		}
	}

	println("The same")
	// Close the channels
	close(ch1)
	close(ch2)
}

func PrintBTree(t *tree.Tree) {
	if t == nil {
		return
	}

	PrintBTree(t.Left)

	fmt.Print(t.Value, " ")

	PrintBTree(t.Right)
}

func PrintBTreePretty(t *tree.Tree, prefix string, isLeft bool) {
	if t == nil {
		return
	}

	if t.Right != nil {
		PrintBTreePretty(t.Right, prefix+"│   ", false)
	}

	// Print the current node
	if isLeft {
		fmt.Printf("%s└── %d\n", prefix, t.Value)
	} else {
		fmt.Printf("%s┌── %d\n", prefix, t.Value)
	}

	if t.Left != nil {
		PrintBTreePretty(t.Left, prefix+"│   ", true)
	}
}
