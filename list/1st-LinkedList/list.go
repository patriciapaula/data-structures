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
func free(list *List) {
	/*var p *List = list
	fmt.Println("----------")
	for {
		if p == nil {
			break
		}
		fmt.Println("-----")
		var t *List = p.next
		p = nil
		p = t
		fmt.Println(p.value)
		fmt.Println(p.next)
	}*/
}

func main() {
	var list = create()
	fmt.Println(intToBool(empty(list)))

	list = add(list, 23)
	list = add(list, 17)
	list = add(list, 45)
	list = add(list, 9)
	list = add(list, 31)

	print(list)
	recursive_print(list)
	reverse_print(list)
	fmt.Println(intToBool(empty(list)))

	node := search(list, 9)
	if node != nil {
		fmt.Println(node.value, node.next)
	}

	list = remove(list, 9)
	print(list)
	list = recursive_remove(list, 17)
	print(list)

	free(list)
	fmt.Println(intToBool(empty(list)))
	print(list)
}
