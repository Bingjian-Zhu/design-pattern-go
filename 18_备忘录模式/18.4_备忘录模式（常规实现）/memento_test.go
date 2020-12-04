package memento

func ExampleMemento() {
	o := new(Originator)
	o.State = "On"
	o.Show()

	c := new(Caretaker)
	c.Memento = o.CreateMemento()

	o.State = "Off"
	o.Show()

	o.SetMemento(c.Memento)
	o.Show()

	// OutPut:
	// State=On
	// State=Off
	// State=On
}
