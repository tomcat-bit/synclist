package synclist

import (
	"container/list"
	"sync"
)

type List struct {
	lock sync.RWMutex
	list *list.List
}

type Element struct {
	*list.Element
}

func New() *List {
	return &List{
		list: list.New(),
	}
}

func (s *List) Back() *list.Element {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.list.Back()
}

func (s *List) Front() *list.Element {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.list.Front()
}

func (l *List) Init() *list.List {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.Init()
}

func (l *List) InsertAfter(v interface{}, mark *list.Element) *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.InsertAfter(v, mark)
}

func (l *List) InsertBefore(v interface{}, mark *list.Element) *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.InsertBefore(v, mark)
}

func (l *List) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.Len()
}

func (l *List) MoveAfter(e, mark *list.Element) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.MoveAfter(e, mark)
}

func (l *List) MoveBefore(e, mark *list.Element) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.MoveBefore(e, mark)
}

func (l *List) MoveToBack(e *list.Element) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.MoveToBack(e)
}

func (l *List) MoveToFront(e *list.Element) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.MoveToFront(e)
}

func (l *List) PushBack(v interface{}) *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.PushBack(v)
}

func (l *List) PushBackList(other *list.List) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.PushBackList(other)
}

func (l *List) PushFront(v interface{}) *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.PushFront(v)
}

func (l *List) PushFrontList(other *list.List) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.PushFrontList(other)
}

func (l *List) Remove(e *list.Element) interface{} {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.Remove(e)
}