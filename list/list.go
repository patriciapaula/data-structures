package list

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

func lst_imprime (lst *Lista) {
	var p *Lista
	for (p = lst; p != NULL; p = p.prox)
		fmt.Println(p.info)
	fmt.Println("fim")
}

int lst_igual (Lista* lst1, Lista* lst2) {
	var p1 *Lista, p2 *Lista //para percorrer l1 e l2
	for (p1 = lst1;)
/* for (p1=lst1, p2=lst2;
p1 != NULL && p2 != NULL;
p1 = p1->prox, p2 = p2->prox)
{
if (p1->info != p2->info) return 0;
}
return p1==p2;
} */
}

func list() {

}