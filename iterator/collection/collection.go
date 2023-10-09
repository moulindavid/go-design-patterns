package collection

type Collection interface {
	createIterator() Iterator
}
