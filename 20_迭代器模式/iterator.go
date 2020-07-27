package iterator

type Iterator interface {
	First() interface{}
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	current   int
}

func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: aggregate,
		current:   0,
	}
}

func (c *ConcreteIterator) First() interface{} {
	return c.aggregate.items[0]
}

func (c *ConcreteIterator) Next() interface{} {
	var ret interface{}
	c.current++

	if c.current < len(c.aggregate.items) {
		ret = c.aggregate.items[c.current]
	}
	return ret
}

func (c *ConcreteIterator) CurrentItem() interface{} {
	return c.aggregate.items[c.current]
}

func (c *ConcreteIterator) IsDone() bool {
	if c.current >= len(c.aggregate.items) {
		return true
	}
	return false
}

type ConcreteIteratorDesc struct {
	aggregate *ConcreteAggregate
	current   int
}

func NewConcreteIteratorDesc(aggregate *ConcreteAggregate) *ConcreteIteratorDesc {
	return &ConcreteIteratorDesc{
		aggregate: aggregate,
		current:   len(aggregate.items) - 1,
	}
}

func (c *ConcreteIteratorDesc) First() interface{} {
	return c.aggregate.items[len(c.aggregate.items)-1]
}

func (c *ConcreteIteratorDesc) Next() interface{} {
	var ret interface{}
	c.current--
	if c.current >= 0 {
		ret = c.aggregate.items[c.current]
	}

	return ret
}

func (c *ConcreteIteratorDesc) CurrentItem() interface{} {
	return c.aggregate.items[c.current]
}

func (c *ConcreteIteratorDesc) IsDone() bool {
	if c.current < 0 {
		return true
	}
	return false
}

type Aggregate interface {
	CreateIterator() Iterator
}

type ConcreteAggregate struct {
	items []interface{}
}

func (c *ConcreteAggregate) CreateIterator() Iterator {
	return &ConcreteIterator{
		aggregate: c,
	}
}

func (c *ConcreteAggregate) Count() int {
	return len(c.items)
}
