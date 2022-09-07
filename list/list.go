package main

import "fmt"

type List struct {
	value int
	next  *List
}

func create() *List {
	return nil
}

func add(list *List, val int) *List {
	var node *List
	node = new(List)
	node.value = val
	node.next = list
	return node
}

func remove(list *List, val int) *List {
	var previous *List = nil /* ponteiro para elemento anterior */
	var p *List = list       /* ponteiro para percorrer a lista */

	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p == nil || p.value == val {
			break
		}
		previous = p
		p = p.next
	}
	/* verifica se achou elemento */
	if p == nil {
		return list /* não achou: ret lista original */
	}
	/* achou: retira */
	if previous == nil { /* retira elemento do inicio */
		list = p.next
	} else { /* retira elemento do meio da lista */
		previous.next = p.next
	}
	p = nil /* libera espaço ocupado pelo elemento */
	return list
}

func empty(list *List) bool {
	return (list == nil)
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

func free(list *List) {
	var p *List = list
	for {
		if p == nil {
			break
		}
		var t *List = p.next
		p = t
	}
}

func print(list *List) {
	if !empty(list) {
		for p := list; p != nil; p = p.next {
			fmt.Println(p.value)
		}
	} else {
		fmt.Println("-")
	}
}

func reverse_print(list *List) {
	if !empty(list) {
		reverse_print(list.next)
		fmt.Println(list.value)
	} else {
		fmt.Println("-")
	}
}

func recursive_print(list *List) {
	if !empty(list) {
		fmt.Println(list.value)
		if list.next != nil {
			recursive_print(list.next)
		}
	} else {
		fmt.Println("-")
	}
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

func main() {
	var list *List
	list = create()
	print(list)
	fmt.Println(empty(list))
	list = add(list, 23)
	list = add(list, 45)
	list = add(list, 9)
	print(list)
	recursive_print(list)
	reverse_print(list)
	//fmt.Println(empty(list))
}
