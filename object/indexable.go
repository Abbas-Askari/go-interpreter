package object

type Indexable interface {
	GetElementAtIndex(int) Object
	SetElementAtIndex(int, Object)
}
