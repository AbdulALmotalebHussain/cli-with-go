// package main

// import (
	
// 	"fmt"

// )

// type node struct {

// 	data int

// 	next *node

// }

// type Link struct {

// 	head   *node

// 	length int

// }

// func (l *Link) prepend(n *node) {

// 	second := l.head

// 	l.head = n

// 	l.head.next = second

// 	l.length++

// }

// func (l Link) printListData() {

// 	toprint := l.head

// 	for l.length != 0 {

// 		fmt.Println(toprint.data)

// 		toprint = toprint.next

// 		l.length--

// 	}

// 	fmt.Printf("%d\n", l.length)

// }

// func (l *Link) deletevalue(value int) {

// 	if l.length==0 {
// 		return 
// 	}
    
//    if l.head.data == value {

// 		l.head = l.head.next

// 		l.length--

// 		return

// 	}

//    	previoustodelete := l.head

// 	for previoustodelete.data != value {

// 		if previoustodelete.next == nil {

// 			return

// 		}

// 		previoustodelete = previoustodelete.next

// 	}

// 	previoustodelete.next = previoustodelete.next.next

// 	l.length--
	
// }

// func main() {

// 	mylist := &Link{}

// 	mylist.prepend(&node{data: 1})

// 	mylist.prepend(&node{data: 2})

// 	mylist.prepend(&node{data: 3})

// 	mylist.prepend(&node{data: 4})

// 	mylist.prepend(&node{data: 5})

// 	mylist.prepend(&node{data: 6})

// 	mylist.prepend(&node{data: 7})

// 	mylist.prepend(&node{data: 8})

// 	mylist.prepend(&node{data: 9})

// 	mylist.prepend(&node{data: 10})

// 	mylist.prepend(&node{data: 11})

// 	mylist.prepend(&node{data: 12})

// 	mylist.prepend(&node{data: 13})

// 	mylist.prepend(&node{data: 14})

// 	mylist.prepend(&node{data: 15})

// 	mylist.prepend(&node{data: 16})

// 	mylist.prepend(&node{data: 17})

// 	mylist.prepend(&node{data: 18})

// 	mylist.prepend(&node{data: 19})

// 	fmt.Println(mylist)

// 	mylist.printListData()

// 	mylist.deletevalue(20)

// 	mylist.deletevalue(14)

// 	mylist.printListData()

// 	emtylist := Link{}

// 	emtylist.deletevalue(50)

// 	fmt.Println(emtylist)

// }