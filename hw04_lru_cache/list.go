package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

func NewList() List {
	return new(list)
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFront := &ListItem{
		Value: v,
	}
	if l.front != nil {
		newFront.Next = l.front
		l.front.Prev = newFront
	}
	l.front = newFront
	if l.back == nil {
		l.back = l.front
	}
	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBack := &ListItem{
		Value: v,
		Prev:  l.back,
	}
	if l.back != nil {
		newBack.Prev = l.back
		l.back.Next = newBack
	}
	l.back = newBack
	if l.front == nil {
		l.front = l.back
	}
	l.len++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if l.back == i {
		l.back = i.Prev
	}

	if l.front == i {
		l.front = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || l.front == i {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i == l.back {
		l.back = i.Prev
	}
	i.Next = l.front
	i.Prev = nil
	l.front.Prev = i
	l.front = i
}
