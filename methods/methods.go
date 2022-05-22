package methods

import (
	"fmt"
	"mypkg/methods/geometry"
)

func Main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println("Geometry distance", geometry.Distance(p, q)) // "5", function call
	fmt.Println("Path distance",p.Distance(q))           // "5", method call
	Perim();

	r := geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(r) // "{2, 4}"
}

func Perim() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println("Perim:", perim.Distance())
}
