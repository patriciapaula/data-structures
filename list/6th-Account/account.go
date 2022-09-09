package main

import (
	"fmt"
)
// incompleta
type Account struct {
	accountType int // 1 corrente, 2 poupanca, 3 fidelidade
	number      float32
	next        *Account
}
type CheckingAccount struct { //corrente
	number float32
	amount float64
}
type SavingsAccount struct { //poupanca
	number float32
	amount float64
}
type LoyaltyAccount struct { //fidelidade
	number float32
	amount float64
	bonus  float64
}

func create() *Account {
	return nil
}

func add(account *Account, typeAccount int, number float32, amount float64, bonus float64) *Account {
	var newAccount *Account
	newAccount = new(Account)
	newAccount.number = number
	switch typeAccount {
    case 1:
        
    case 2:
        
    case 3:
        
    }
	newAccount.amount = amount
	newAccount.next = account
	return newAccount
}
func addCheckingAccount {
	
}

func (l LoyaltyAccount) add(account *LoyaltyAccount, number float32, amount float64, bonus float64) *LoyaltyAccount {
	var newAccount *LoyaltyAccount
	newAccount = new(LoyaltyAccount)
	newAccount.number = number
	newAccount.amount = amount
	newAccount.next = account
	return newAccount
}

func remove(account *Account, val int) *Account {
	var previous *Account = nil /* ponteiro para elemento anterior */
	var p *Account = account    /* ponteiro para percorrer a lista */

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
		return account /* não achou: ret lista original */
	}
	/* achou: retira */
	if previous == nil { /* retira elemento do inicio */
		account = p.next
	} else { /* retira elemento do meio da lista */
		previous.next = p.next
	}
	p = nil /* libera espaço ocupado pelo elemento */
	return account
}

func recursive_remove(account *Account, number int) *Account {
	if !intToBool(empty(account)) {
		if account.number == number {
			account = account.next
			//free(t)
		} else {
			account.next = recursive_remove(account.next, number)
		}
	}
	return account
}

func empty(account *Account) int {
	//return (account == nil)
	if account == nil {
		return 1
	}
	return 0
}

func search(account *Account, number int) *Account {
	var p *Account
	for p = account; p != nil; p = p.next {
		if p.number == number {
			return p
		}
	}
	return nil
}

func print(account *Account) {
	if !intToBool(empty(account)) {
		for p := account; p != nil; p = p.next {
			fmt.Println(p.number, p.amount)
		}
	}
}

func reverse_print(account *Account) {
	if !intToBool(empty(account)) {
		reverse_print(account.next)
		fmt.Println(account.number, account.amount)
	}
}

func recursive_print(account *Account) {
	if !intToBool(empty(account)) {
		fmt.Println(account.number, account.amount)
		if account.next != nil {
			recursive_print(account.next)
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
func free(account *Account) *Account {
	return nil
}

func main() {
	var account = create()
	fmt.Println(intToBool(empty(account)))

	account = add(account, 23)
	account = add(account, 17)
	account = add(account, 45)
	account = add(account, 9)
	account = add(account, 31)

	print(account)
	recursive_print(account)
	reverse_print(account)
	fmt.Println(intToBool(empty(account)))

	account := search(account, 9)
	if account != nil {
		fmt.Println(account.value, account.next)
	}

	account = remove(account, 9)
	print(account)
	account = recursive_remove(account, 17)
	print(account)

	account = free(account) // O GC faz a liberacao
	fmt.Println(intToBool(empty(account)))
}
