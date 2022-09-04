package main

import "fmt"

type Lista struct {
	info int
	prox *Lista
}

func lst_cria() *Lista {
	return nil
}

func lst_insere(lst *Lista, val int) *Lista {
	var novo *Lista
	novo = new(Lista)
	novo.info = val
	novo.prox = lst
	return novo
}

func lst_vazia(lst *Lista) bool {
	return (lst == nil)
}

func lst_busca(lst *Lista, val int) *Lista {
	var p *Lista
	for p := lst; p != nil; p = p.prox {
		if p.info == val {
			return p
		}
	}
	return nil
}

func lst_libera(lst *Lista) {
	var p *Lista = lst
	for {
		//work()
		if p == nil {
			break
		}
		var t *Lista = p.prox
		//free(p)
		p = t
	}
}

func lst_imprime (lst *Lista)
{
	var p *Lista
	for (p = lst; p != NULL; p = p.prox)
		fmt.Println(p.info)
	fmt.Println("fim")
}


func main() {

}
