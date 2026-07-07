package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
}
