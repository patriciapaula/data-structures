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

func add(accounts *Account, typeAccount int, number float32, amount float64, bonus float64) *Account {
	var newAccount *Account
	newAccount = new(Account)
	newAccount.accountType = typeAccount
	newAccount.number = number
	newAccount.next = accounts
	switch typeAccount {
	case 1:
		addCheckingAccount(number, amount)
	case 2:
		addSavingsAccount(number, amount)
	case 3:
		addLoyaltyAccount(number, amount, bonus)
	}
	return newAccount
}
func addCheckingAccount(number float32, amount float64) {
	var checking *CheckingAccount
	checking = new(CheckingAccount)
	checking.amount = amount
	checking.number = number
}
func addSavingsAccount(number float32, amount float64) {
	var savings *SavingsAccount
	savings = new(SavingsAccount)
	savings.amount = amount
	savings.number = number
}
func addLoyaltyAccount(number float32, amount float64, bonus float64) {
	var loyalty *LoyaltyAccount
	loyalty = new(LoyaltyAccount)
	loyalty.amount = amount
	loyalty.number = number
	loyalty.bonus = bonus
}

func remove(account *Account, number int) *Account {
	var previous *Account = nil /* ponteiro para elemento anterior */
	var p *Account = account    /* ponteiro para percorrer a lista */

	/* procura elem na lista, guardando anterior  - while com for */
	for {
		if p == nil || int(p.number) == number {
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

func recursive_remove(account *Account, number float32) *Account {
	if !intToBool(empty(account)) {
		if account.number == number {
			account = account.next
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

func search(account *Account, number float32) *Account {
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
			fmt.Println(p.number, p.accountType)
		}
	}
}

func reverse_print(account *Account) {
	if !intToBool(empty(account)) {
		reverse_print(account.next)
		fmt.Println(account.number, account.accountType)
	}
}

func recursive_print(account *Account) {
	if !intToBool(empty(account)) {
		fmt.Println(account.number, account.accountType)
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

	account = add(account, 1, 235, 230, 0)
	account = add(account, 3, 563, 230, 50)
	account = add(account, 2, 123, 230, 0)
	account = add(account, 2, 654, 230, 0)
	account = add(account, 1, 205, 150, 0)
	account = add(account, 3, 241, 80, 10)
	account = add(account, 1, 794, 100, 0)

	print(account)
	recursive_print(account)
	reverse_print(account)
	fmt.Println(intToBool(empty(account)))

	var ac = search(account, 654)
	if ac != nil {
		fmt.Println(account.number, account.accountType)
	}

	ac = remove(account, 205)
	print(account)
	ac = recursive_remove(account, 794)
	print(account)

	account = free(account) // O GC faz a liberacao
	fmt.Println(intToBool(empty(account)))
}
