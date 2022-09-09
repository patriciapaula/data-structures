package main

import (
	"fmt"
)

type List struct {
	value    int
	previous *List
	next     *List
}

func create() *List {
	return nil
}

func equal(lst1 *List, lst2 *List) bool {
	var p1 *List
	var p2 *List
	for p1, p2 = lst1, lst2; p1 != nil && p2 != nil; p1, p2 = p1.next, p2.next {
		if p1.value != p2.value {
			return false
		}
	}
	return p1 == p2
}

// ***** this method has been changed
func add(list *List, val int) *List {
	var node *List
	var previous *List = nil
	var p *List = list
	//search for insertion point
	for {
		if p == nil || p.value >= val {
			break
		}
		previous = p
		p = p.next
	}
	node = new(List)
	node.value = val
	//insert in the beginning
	if previous == nil {
		node.next = list
		node.previous = nil
		list = node
	} else { //in the middle
		node.next = previous.next
		node.previous = previous
		previous.next = node
	}
	return list
}

// ***** changed
func remove(list *List, val int) *List {
	var p *List = search(list, val)
	/* verifica se achou elemento */
	if p == nil {
		return list /* não achou */
	}
	if list == p { /* se é o primeiro elemento */
		list = p.next
	} else { /* retira do meio da lista */
		p.previous.next = p.next
	}
	if p.next != nil {
		p.next.previous = p.previous
	}
	p = nil
	return list
}

func recursive_remove(list *List, val int) *List {
	if !intToBool(empty(list)) {
		if list.value == val {
			list = list.next
		} else {
			list.next = recursive_remove(list.next, val)
		}
	}
	return list
}

func empty(list *List) int {
	//return (list == nil)
	if list == nil {
		return 1
	}
	return 0
}

func search(list *List, val int) *List {
	var p *List
	for p = list; p != nil; p = p.next {
		if p.value == val {
			return p
		}
	}
	return nil
}

func print(list *List) {
	if !intToBool(empty(list)) {
		for p := list; p != nil; p = p.next {
			fmt.Println(p.value) //, p.previous, p.next)
		}
	}
}

func reverse_print(list *List) {
	if !intToBool(empty(list)) {
		reverse_print(list.next)
		fmt.Println(list.value) //, list.previous, list.next)
	}
}

func recursive_print(list *List) {
	if !intToBool(empty(list)) {
		fmt.Println(list.value) //, list.previous, list.next)
		if list.next != nil {
			recursive_print(list.next)
		}
	}
}

func intToBool(b int) bool {
	if b == 1 {
		return true
	}
	return false
}

// O GC do Go eh quem faz a liberação
func free(list *List) *List {
	return nil
}

func main() {
	var list = create()
	fmt.Println(intToBool(empty(list)))
	//ordered
	list = add(list, 9)
	list = add(list, 17)
	list = add(list, 21)
	list = add(list, 31)
	list = add(list, 45)

	print(list)
	recursive_print(list)
	reverse_print(list)
	fmt.Println(intToBool(empty(list)))

	node := search(list, 45)
	if node != nil {
		fmt.Println(node.value, node.next)
	}

	list = remove(list, 17)
	print(list)
	list = recursive_remove(list, 31)
	print(list)

	var list2 = create()
	list2 = add(list2, 9)
	list2 = add(list2, 17)
	list2 = add(list2, 21)
	list2 = add(list2, 31)
	list2 = add(list2, 45)
	fmt.Println(equal(list, list2))

	var list3 = create()
	list3 = add(list3, 9)
	list3 = add(list3, 21)
	list3 = add(list3, 45)
	fmt.Println(equal(list, list3))

	list = free(list) // O GC faz a liberacao
	fmt.Println(intToBool(empty(list)))
}
