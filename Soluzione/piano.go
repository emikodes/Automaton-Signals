//41461A INGENITO EMIDDIO

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type piano struct {
	Automi   map[string]*automa
	Ostacoli *[]ostacolo
	Mappa    map[coordinate][]interface{}
}

func newPiano() piano {
	return piano{
		Automi:   make(map[string]*automa),
		Ostacoli: &[]ostacolo{},
		Mappa:    make(map[coordinate][]interface{}),
	}
}

func (p *piano) stampa() {
	fmt.Println("(")
	if len(p.Automi) > 0 {
		for _, automa := range p.Automi {
			automa.stampa()
		}
	}
	fmt.Println(")")
	fmt.Println("[")

	if len(*p.Ostacoli) > 0 {
		for _, ostacolo := range *p.Ostacoli {
			ostacolo.stampa()
		}
	}
	fmt.Println("]")
}

func (p *piano) stato(x, y int) string {
	pos := nuoveCoordinate(x, y)

	if entities, exists := p.Mappa[pos]; exists && len(entities) > 0 {
		switch entities[0].(type) {
		case *automa:
			return "A"
		case *ostacolo:
			return "O"
		}
	}

	return "E"
}

func (p *piano) richiamo(x, y int, nome string) {
	sorgente := nuoveCoordinate(x, y)

	if !p.isOstacolo(sorgente) {
		// Mappa per raggruppare gli automi in base alla distanza dalla sorgente
		automaByDistanza := make(map[int][]*automa)
		distanze := []int{}

		// Consideriamo tutti gli automi nel piano
		for _, automa := range p.Automi {
			if strings.HasPrefix(automa.Nome, nome) {
				distanza := automa.Posizione.Distanza(sorgente)
				automaByDistanza[distanza] = append(automaByDistanza[distanza], automa)
				distanze = append(distanze, distanza)
			}
		}

		sort.Ints(distanze)

		spostati := false
		for _, dist := range distanze {
			if !spostati {
				for _, automa := range automaByDistanza[dist] {
					// Se esiste un percorso libero verso la sorgente, spostiamo l'automa
					if p.esistePercorso(automa.Posizione, sorgente) {
						p.automa(sorgente.X, sorgente.Y, automa.Nome)
						spostati = true
					}
				}
			}

		}
	}
}

// Variante di BFS
func (p *piano) esistePercorso(start, end coordinate) bool {
	// Mappa delle coordinate già visitate
	visited := make(map[coordinate]bool)

	// Coda delle coordinate da visitare
	queue := []coordinate{start}

	// Marco start come visitato.
	visited[start] = true

	// Finchè la coda non è vuota.
	for len(queue) > 0 {
		// Nodo da visitare
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true
		}

		movimenti := []coordinate{
			{X: current.X, Y: current.Y + 1}, // Giù
			{X: current.X, Y: current.Y - 1}, // Su
			{X: current.X + 1, Y: current.Y}, // Destra
			{X: current.X - 1, Y: current.Y}, // Sinistra
		}

		//Espando il nodo estratto, nei 4 movimenti (su, giù, destra,sinistra.)
		for _, next := range movimenti {
			//Se lo spostamento mi porta in una "casella" non visitata, che non sia un ostacolo, e che rimanga "nell'area" tra start ed end
			if !visited[next] && !p.isOstacolo(next) &&
				next.X >= int(math.Min(float64(start.X), float64(end.X))) && next.X <= int(math.Max(float64(start.X), float64(end.X))) &&
				next.Y >= int(math.Min(float64(start.Y), float64(end.Y))) && next.Y <= int(math.Max(float64(start.Y), float64(end.Y))) {

				//Marco il nodo come visitato
				visited[next] = true
				//Lo aggiungo alla coda dei nodi da espandere
				queue = append(queue, next)
			}
		}
	}

	return false
}

func (p *piano) posizioni(prefisso string) {
	fmt.Println("(")
	for _, automa := range p.Automi {
		if strings.HasPrefix(automa.Nome, prefisso) {
			automa.stampa()
		}
	}
	fmt.Println(")")
}

func (p *piano) isOstacolo(pos coordinate) bool {
	return p.stato(pos.X, pos.Y) == "O"
}

func (p *piano) isAutoma(pos coordinate) bool {
	return p.stato(pos.X, pos.Y) == "A"
}
