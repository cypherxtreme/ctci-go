// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm ADD_LISTS(list1,list2):
 1.Get the input list1 and list2
 2.if list1==nil && list2==nil then return nil
 	else  if list == nill || list2 == nil
 		return list1==nil?list2:list1
 3.add elements and maintain carry from lists until both the lists has elements
 4.Add carry to the continuing list and add to result list if the lists are of different size
 5.return the result list
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(7)
	l.PushFront(9)


	m := list.New()
	m.PushFront(3)
	m.PushFront(2)
	m.PushFront(8)
	//m.PushFront(6)
	res := addLists(l, m)
	for e := res.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println(" ")


}

//ERROR

func addLists(l *list.List, m *list.List) *list.List {

	if l == nil && m == nil {
		return nil
	}
	lLength := l.Len()
	mLength := m.Len()

	carry := 0
	value := 0
	resList := list.New()

	var e *list.Element
	var f *list.Element
	for e, f = l.Front(), m.Front(); e != nil && f != nil; e, f = e.Next(), f.Next() {
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
