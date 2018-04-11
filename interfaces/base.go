package interfaces

type A interface {
	Haha(int) int
}

type B struct {}

func (b *B) Haha(i int) int {
	return i
}

