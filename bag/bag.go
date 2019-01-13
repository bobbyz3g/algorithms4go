package bag

type node struct {
	item interface{}
	next *node
}

type Bag struct {
	first *node
	size  int
}

func New() *Bag {
	return &Bag{
		nil,
		0,
	}
}

// Add adds a new element to bag.
func (b *Bag) Add(item interface{}) {
	b.first = &node{item, b.first}
	b.size++
}

// Empty returns true if bag is empty, or false if bag is not empty.
func (b *Bag) Empty() bool {
	return b.size == 0
}

// Size returns size of bag.
func (b *Bag) Size() int {
	return b.size
}
