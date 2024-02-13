package structure

type LinkedList struct {
	Head *Node
	Last *Node
}

func (l *LinkedList) DeleteBefore(n *Node) {
	if n.Prev == nil {
		return
	}
	if n.Prev.Prev != nil {
		n.Prev = n.Prev.Prev
		n.Prev.Next = n
	} else {
		l.Head = n
		n.Prev = nil
	}
}

func (l *LinkedList) InsertAfter(curN, newN *Node) {
	if curN == nil {
		l.Head = newN
		l.Last = newN
		return
	}
	newN.Next = curN.Next
	curN.Next = newN
	if newN.Next != nil {
		newN.Next.Prev = newN
	} else {
		l.Last = newN
	}
	newN.Prev = curN
}
