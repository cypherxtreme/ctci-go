// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm FIND_STARTING_POINT(list):
 1.Get the input list
 2.if list==nil return nil
 3.Make two pointers fast and slow
 4.move slow once and fast twice on every iteration
 5.if slow == fast then break
 6.if fast == nil || fast.next == nil then return nil
 7.slow = head
 8.Iterate through the 
 3.add elements and maintain carry from lists until both the lists has elements
 4.Add carry to the continuing list and add value to result list if the lists are of different size
 5.return the result list
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	e1 := l.PushFront(4)
	l.PushFront(5)
	l.PushFront(7)
	l.PushFront(9)


	m := list.New()
	m.PushFront(3)
	m.PushFront(2)
	m.PushFront(8)
	//m.PushFront(6)
	res := findLoopsInList(l, m)
	for e := res.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println(" ")


}

func findLoopsInList(l *list.List) *list.List {

	if l == nil {
		return nil
	}

	var slow *list.Element
	var fast *list.Element
	for slow, fast = l.Front(), l.Front(); slow != nil && fast != nil; slow, fast = slow.Next(), fast.Next().Next() {
		if slow == fast {
			break
		}
		value = carry + e.Value.(int) + f.Value.(int)
		carry = 0
		carry = value / 10
		value = value % 10
		resList.PushFront(value)
	}

	var p *list.Element
	if lLength > mLength {
		p = e
	} else {
		p = f
	}

	for ; p != nil; p = p.Next() {
		value = carry + p.Value.(int)
		carry = 0
		carry = value / 10
		value = value % 10
		resList.PushFront(value)
	}

	if carry != 0 {
		resList.PushFront(carry)
	}

	return resList
}
