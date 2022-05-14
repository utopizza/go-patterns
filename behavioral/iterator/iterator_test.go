package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	itemList := &ItemList{
		Items: []Item{
			{"item1"},
			{"item2"},
			{"item3"},
		},
	}

	iterator := itemList.NewItemListIterator()
	for iterator.HasNext() {
		item := iterator.GetNext().(Item)
		fmt.Printf("item:%s\n", item.Name)
	}
}
