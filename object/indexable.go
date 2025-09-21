package object

type Indexable interface {
	GetElementAtIndex(Object) (Object, error)
	SetElementAtIndex(Object, Object) error
}
