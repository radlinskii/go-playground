package ctci

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head *node
}

func makeList() *list {
	return &list{}
}

func (l *list) appendHead(d int) {
	n := &node{data: d}
	n.next = l.head
	l.head = n
}

func (l *list) appendTail(d int) {
	n := l.head
	if n == nil {
		l.head = &node{data: d}
		return
	}
	for n.next != nil {
		n = n.next
	}
	n.next = &node{data: d}
}

func (l *list) print() {
	n := l.head
	fmt.Printf("->")
	for n != nil {
		fmt.Printf("%d->", n.data)
		n = n.next
	}
}
