package main

import (
	"fmt"
	"math"
)

type List struct {
	value string
	next  *List
}
type HashTable struct {
	key   int
	value *List
	next  *HashTable
}

func create() *HashTable {
	return nil
}

func searchList(list *List, val string) *List {
	var p *List
	for p = list; p != nil; p = p.next {
		if p.value == val {
			return p
		}
	}
	return nil
}

// returns a hash between 0 and m-1
func hash(key int) int {
	return (key & 0x7fffffff) % m
}

func add(hash *HashTable, val string) *HashTable {
	var list *List
	list = new(List)
	list.value = val
	list.next = nil

	var item *HashTable
	item = new(HashTable)
	item.key = 0 // *** falta
	item.value = list
	item.next = hash

	return item
}

func remove(list *HashTable, val string) *HashTable {
	var previous *HashTable = nil /* ponteiro para elemento anterior */
	var p *HashTable = list       /* ponteiro para percorrer a lista */

	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p == nil || searchList(p.value, val).value == val {
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

func empty(list *HashTable) bool {
	return (list == nil)
}

func search(hash *HashTable, val string) *HashTable {
	var p *HashTable
	for p = hash; p != nil; p = p.next {
		for i := p.value; i != nil; i = i.next {
			if i.value == val {
				return p
			}
		}
	}
	return nil
}

func print(list *HashTable) {
	if !empty(list) {
		for p := list; p != nil; p = p.next {
			fmt.Println(p.value)
		}
	}
}

// O GC do Go eh quem faz a liberação
func free(list *HashTable) *HashTable {
	return nil
}

var n float64
var m int

func main() {
	//fmt.Print("Informe um num inteiro:")
	//fmt.Scanln("%d\n", &n)

	n = 12
	m = int(math.Round(n / 2)) //returns the nearest integer

	var hash = create()
	fmt.Println(empty(hash))

	hash = add(hash, "23")
	hash = add(hash, "oi")
	hash = add(hash, "-")
	hash = add(hash, "9")
	hash = add(hash, "abc")

	print(hash)
	fmt.Println(empty(hash))

	node := search(hash, "9")
	if node != nil {
		fmt.Println(node.value, node.next)
	}

	//list = remove(list, "9")
	//print(list)

	hash = free(hash) // O GC faz a liberacao
	fmt.Println(empty(hash))
}
