package object

type Indexable interface {
	GetElementAtIndex(Object) Object
	SetElementAtIndex(Object, Object)
}
