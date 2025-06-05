package snake

import (
	"fmt"
	"game_snake/point"
)

type block struct {
	coord point.Point
	next *block
	prev *block
	direction Direction
}

type queue struct {
	head *block
	tail *block
}

func New() *queue {
	c := point.Point{X : -1,Y : -1}
	dummy := &block{c, nil, nil, Up}
	return &queue{dummy, dummy}
}

func (q *queue) PushBack(c point.Point, dir Direction) {
	tmp := &block{c, nil, q.tail, dir}
	q.tail.next = tmp
	q.tail = tmp
}

func (q *queue) PushFront(c point.Point, dir Direction) {
	tmp := &block{c, nil, nil, dir}
	tmp.next = q.head.next
	tmp.prev = q.head
	if (q.head.next == nil) {
		q.tail = tmp
	} else {
		q.head.next.prev = tmp
	}
	q.head.next = tmp
}

func (q *queue) PopBack() (point.Point, bool) {
	if q.tail.prev == nil || q.tail == nil {
		return point.Point{X : -1, Y : -1}, false
	}
	c := q.tail.coord
	q.tail = q.tail.prev
	q.tail.next = nil
	return c, true
}

func (q *queue) PopFront() (point.Point, bool) {
	if q.head.next == nil || q.head == nil {
		return point.Point{X : -1,Y : -1}, false
	}
	c := q.tail.coord
	q.head.next.prev = nil
	q.head.next = q.head.next.next
	if q.head.next != nil {
		q.head.next.prev = nil
	}
	return c, true
}

func (q *queue) PrintFromTail() {
	dummy := q.tail
	for dummy != q.head {
		fmt.Printf("%v %v\n", dummy.coord, dummy.direction)
		dummy = dummy.prev
	}
}

func (q *queue) PrintFromHead() {
	dummy := q.head.next
	for dummy != nil {
		fmt.Printf("%v %v\n", dummy.coord, dummy.direction)
		dummy = dummy.next
	}
}

func (q *queue) Clear() {
	q.head.next = nil
	q.tail = q.head
}