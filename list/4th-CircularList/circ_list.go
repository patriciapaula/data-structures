package main

import "fmt"

type List struct {
	value int
	next  *List
}

var head *List = nil
var tail *List = nil

func create() *List {
	return nil
}

// --- muda onde verifica nil e no add/remove ---

func add(list *List, val int) *List {
	var node *List
	node = new(List)
	node.value = val
	node.next = list
	if tail == nil {
		tail = node
	}
	head = node
	tail.next = node
	return node
}

func remove(list *List, val int) *List {
	var previous *List = nil /* ponteiro para elemento anterior */
	var p *List = list       /* ponteiro para percorrer a lista */
	var head int = list.value

	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p.next.value == head || p.value == val {
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
	//if previous == nil { /* retira elemento do inicio */ // -----
	//	list = p.next
	//} else { /* retira elemento do meio da lista */
	previous.next = p.next
	//}
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
	var p *List = list
	for {
		if p.value == val {
			return p
		}
		if p.next.value == head.value {
			break
		}
		p = p.next
	}
	return nil
}

func print(list *List) {
	if !intToBool(empty(list)) {
		p := list
		for {
			fmt.Println(p.value, p.next.value)
			if p.next == head {
				break
			}
			p = p.next
		}
	}
}

func reverse_print(list *List) {
	if !intToBool(empty(list)) {
		if list.next != head {
			reverse_print(list.next)
		}
		fmt.Println(list.value, list.next.value)
	}
}

func recursive_print(list *List) {
	if !intToBool(empty(list)) {
		fmt.Println(list.value, list.next.value)
		if list.next != head {
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

	list = add(list, 9)
	list = add(list, 17)
	list = add(list, 23)
	list = add(list, 31)
	list = add(list, 45)

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

	list = free(list) // O GC faz a liberacao
	fmt.Println(intToBool(empty(list)))
	print(list)
}
