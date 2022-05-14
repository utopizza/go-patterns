package iterator

type Iterator interface {
	HasNext() bool
	GetNext() interface{}
}

type ItemListIterator struct {
	Index int
	Items []Item
}

func (i *ItemListIterator) HasNext() bool {
	return i.Index < len(i.Items)
}

func (i *ItemListIterator) GetNext() interface{} {
	if i.HasNext() {
		item := i.Items[i.Index]
		i.Index++
		return item
	}
	return nil
}
