package main

import (
	"../quadtree"
)

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
	qt := &quadtree.Quadtree{
		Bounds: quadtree.Bounds{
			X:      0,
			Y:      0,
			Width:  100,
			Height: 100,
		},
		MaxObjects: 10,
		MaxLevels:  4,
		Level:      0,
		Objects:    make([]quadtree.Bounds, 0),
		Nodes:      make([]quadtree.Quadtree, 0),
	}

	// Insert some boxes
	qt.Insert(quadtree.Bounds{
		X:      1,
		Y:      1,
		Width:  10,
		Height: 10,
	})
	qt.Insert(quadtree.Bounds{
		X:      5,
		Y:      5,
		Width:  10,
		Height: 10,
	})
	qt.Insert(quadtree.Bounds{
		X:      10,
		Y:      10,
		Width:  10,
		Height: 10,
	})
	qt.Insert(quadtree.Bounds{
		X:      15,
		Y:      15,
		Width:  10,
		Height: 10,
	})

	//Get all the intersections
	intersections := qt.RetrieveIntersections(quadtree.Bounds{
		X:      5,
		Y:      5,
		Width:  2.5,
		Height: 2.5,
	})

	fmt.Println(intersections)
	// Clear the Quadtree
	qt.Clear()
}
