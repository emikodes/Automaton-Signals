//41461A INGENITO EMIDDIO

package main

import "fmt"

type ostacolo struct {
	verticeBassoSinistra, verticeAltoDestra coordinate
}

func (o *ostacolo) stampa() {
	fmt.Printf("(%d,%d)(%d,%d)\n",
		o.verticeBassoSinistra.X, o.verticeBassoSinistra.Y,
		o.verticeAltoDestra.X, o.verticeAltoDestra.Y)
}

func (p *piano) ostacolo(x0, y0, x1, y1 int) {
	bassoSinistro := nuoveCoordinate(x0, y0)
	altoDestro := nuoveCoordinate(x1, y1)

	// Se ci sono automi nell'area dell'ostacolo, non esegue alcun'operazione.
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			if p.isAutoma(nuoveCoordinate(x, y)) {
				return
			}
		}
	}

	//Aggiungo ostacolo alla lista degli ostacoli del piano
	nuovoOstacolo := ostacolo{bassoSinistro, altoDestro}
	*p.Ostacoli = append(*p.Ostacoli, nuovoOstacolo)

	// Marca l'area dell'ostacolo nella mappa.
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			pos := nuoveCoordinate(x, y)
			p.Mappa[pos] = append(p.Mappa[pos], &nuovoOstacolo)
		}
	}
}
