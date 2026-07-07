package piscine

func ListSize(l *List) int {
	if l == nil || l.Head == nil {
		return 0
	}
	count := 0
	node := l.Head
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
