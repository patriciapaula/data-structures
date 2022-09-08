package main

import "fmt"

type List struct {
	value int
	next  *List
}

func create() *List {
	return nil
}

// ***** new method
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

// ***** add method is different from the 1st question add method
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
		list = node
	} else { //in the middle
		node.next = previous.next
		previous.next = node
	}
	return list
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
	//free(p)
	return list
}

func recursive_remove(list *List, val int) *List {
	if !intToBool(empty(list)) {
		if list.value == val {
			list = list.next
			//free(t)
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
			fmt.Println(p.value)
		}
	}
}

func reverse_print(list *List) {
	if !intToBool(empty(list)) {
		reverse_print(list.next)
		fmt.Println(list.value)
	}
}

func recursive_print(list *List) {
	if !intToBool(empty(list)) {
		fmt.Println(list.value)
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
	print(list)
}
