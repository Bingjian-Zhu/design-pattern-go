package composite

func ExampleComposite() {
	root := &Composite{name: "root"}
	root.Add(&Leaf{name: "Leaf A"})
	root.Add(&Leaf{name: "Leaf B"})

	comp := &Composite{name: "Composite X"}
	comp.Add(&Leaf{name: "Leaf XA"})
	comp.Add(&Leaf{name: "Leaf XB"})

	root.Add(comp)

	comp2 := &Composite{name: "Composite XY"}
	comp2.Add(&Leaf{name: "Leaf XYA"})
	comp2.Add(&Leaf{name: "Leaf XYB"})

	comp.Add(comp2)

	root.Add(&Leaf{name: "Leaf C"})

	leaf := &Leaf{name: "Leaf D"}
	root.Add(leaf)
	root.Remove(leaf)

	root.Display(1)

	// OutPut:
	// -root
	// ---Leaf A
	// ---Leaf B
	// ---Composite X
	// -----Leaf XA
	// -----Leaf XB
	// -----Composite XY
	// -------Leaf XYA
	// -------Leaf XYB
	// ---Leaf C
}
