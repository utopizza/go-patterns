package observer

type Subject interface {
	Register(o Observer)
	Deregister(o Observer)
	NotifyAll()
}

type Item struct {
	ObserverMap map[string]Observer
	Name        string
}

func (i *Item) Register(o Observer) {
	if i.ObserverMap == nil {
		i.ObserverMap = make(map[string]Observer)
	}
	i.ObserverMap[o.GetID()] = o
}

func (i *Item) Deregister(o Observer) {
	if i.ObserverMap != nil {
		delete(i.ObserverMap, o.GetID())
	}
}

func (i *Item) NotifyAll() {
	for _, o := range i.ObserverMap {
		o.Update(i.Name)
	}
}
