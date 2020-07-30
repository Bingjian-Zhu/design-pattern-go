package flyweight

func ExampleFlyweight() {
	extrinsicstate := 22

	f := NewFlyweightFactory()

	fx := f.GetFlyweight("X")
	extrinsicstate--
	fx.Operation(extrinsicstate)

	fy := f.GetFlyweight("Y")
	extrinsicstate--
	fy.Operation(extrinsicstate)

	fz := f.GetFlyweight("Z")
	extrinsicstate--
	fz.Operation(extrinsicstate)

	uf := &UnsharedConcreteFlyweight{}
	extrinsicstate--
	uf.Operation(extrinsicstate)

	// OutPut:
	// 具体Flyweight:21
	// 具体Flyweight:20
	// 具体Flyweight:19
	// 不共享的具体Flyweight:18

}
