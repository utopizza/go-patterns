package iterator

type Item struct {
	Name string
}

type ItemList struct {
	Items    []Item
	Iterator Iterator
}

func (i *ItemList) NewItemListIterator() *ItemListIterator {
	return &ItemListIterator{
		Items: i.Items,
		Index: 0,
	}
}
