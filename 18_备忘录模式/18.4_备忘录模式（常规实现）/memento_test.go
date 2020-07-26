package memento

func ExampleMemento() {
	o := &Originator{}
	o.State = "On"
	o.Show()

	c := &Caretaker{}
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
