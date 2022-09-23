package main

import (
	"fmt"
	"math"
	//"crypto/sha256"
)

type List struct {
	value string
	next  *List
}
type HashItem struct {
	key   int
	value *List
	next  *HashItem
}

func create() *HashItem {
	return nil
}

// lists in hashtable
func addList(list *List, val string) *List {
	var node *List
	node = new(List)
	node.value = val
	node.next = list
	return node
}
func removeList(list *List, val string) *List {
	var previous *List = nil
	var p *List = list
	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p == nil || p.value == val {
			break
		}
		previous = p
		p = p.next
	}
	if p == nil {
		return list /* não achou: ret lista original */
	}
	if previous == nil { /* retira elemento do inicio */
		list = p.next
	} else { /* retira elemento do meio da lista */
		previous.next = p.next
	}
	p = nil /* libera espaço ocupado pelo elemento */
	return list
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

func search(hash *HashItem, val string) *List {
	var p *HashItem
	var list *List
	for p = hash; p != nil; p = p.next {
		for list = p.value; list != nil; list = list.next {
			if list.value == val {
				return list
			}
		}
	}
	return nil
}

// returns a hash between 0 and m-1
func hashF(key int) int {
	return (key & 0x7fffffff) % m
}

func add(hash *HashItem, val string) *HashItem {
	var list *List
	var p *HashItem = hash
	var previous *HashItem
	var newPos = 0

	count++

	countPos := count - 1
	for {
		newPos = hashF(countPos) // search new position
		if countPos >= newPos && p == nil {
			break
		}
		countPos++
		previous = p
		p = p.next
	}

	if countPos <= 1 { //first
		list = new(List)
		list.value = val
		list.next = nil

		p = new(HashItem)
		p.key = newPos
		p.value = list
	} else {
		list = addList(p.value, val) // erro ***
		p.value = list
		//if previous != nil { //there is a previous
		previous.next = p
		//}
	}
	return p
}

func remove(list *HashItem, val string) *HashItem {
	var previous *HashItem = nil
	var p *HashItem = list

	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p == nil || searchList(p.value, val).value == val {
			break
		}
		previous = p
		p = p.next
	}
	if p == nil {
		return list /* não achou: ret lista original */
	}
	count--
	if previous == nil { /* retira elemento do inicio */
		list = p.next
	} else { /* retira elemento do meio da lista */
		previous.next = p.next
	}
	p = nil /* libera espaço ocupado pelo elemento */
	return list
}

func empty(list *HashItem) bool {
	return (list == nil)
}

func print(hash *HashItem) {
	if !empty(hash) {
		for p := hash; p != nil; p = p.next {
			for list := p.value; list != nil; list = list.next {
				fmt.Println(p.value)
			}
		}
	}
}

// O GC do Go eh quem faz a liberação
func free(list *HashItem) *HashItem {
	return nil
}

var n float64
var m int
var count int = 0

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
	print(hash)

	hash = free(hash) // O GC faz a liberacao
	fmt.Println(empty(hash))
}
