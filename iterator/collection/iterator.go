package collection

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
