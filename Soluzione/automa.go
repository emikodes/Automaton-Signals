//41461A INGENITO EMIDDIO

package main

import "fmt"

type automa struct {
	Nome      string
	Posizione coordinate
}

func (a *automa) stampa() {
	fmt.Printf("%s: %d,%d\n", a.Nome, a.Posizione.X, a.Posizione.Y)
}

func (p *piano) automa(x, y int, nome string) {
	pos := nuoveCoordinate(x, y)

	if !p.isOstacolo(pos) {
		// Rimuovi l'automa dalla vecchia posizione se esiste
		if automa, exists := p.Automi[nome]; exists {
			delete(p.Mappa, automa.Posizione)
		}

		automa := &automa{
			Nome:      nome,
			Posizione: pos,
		}

		p.Automi[nome] = automa
		p.Mappa[pos] = append(p.Mappa[pos], automa)

	}
}
