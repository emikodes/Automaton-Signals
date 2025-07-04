//41461A INGENITO EMIDDIO

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func esegui(p piano, s string) {
	args := strings.Fields(s)
	if len(args) == 0 {
		return
	}

	cmd := args[0]
	switch cmd {
	case "c": //clear
		for a := range p.Automi {
			delete(p.Automi, a)
		}
		for c := range p.Mappa {
			delete(p.Mappa, c)
		}
		*p.Ostacoli = []ostacolo{}
	case "S": //Stampa
		p.stampa()
	case "s":
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		fmt.Println(p.stato(x, y))
	case "a":
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		nome := args[3]
		p.automa(x, y, nome)
	case "o":
		x0, _ := strconv.Atoi(args[1])
		y0, _ := strconv.Atoi(args[2])
		x1, _ := strconv.Atoi(args[3])
		y1, _ := strconv.Atoi(args[4])
		p.ostacolo(x0, y0, x1, y1)
	case "r":
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		nome := args[3]
		p.richiamo(x, y, nome)
	case "p":
		prefisso := args[1]
		p.posizioni(prefisso)
	case "e":
		x, _ := strconv.Atoi(args[1])
		y, _ := strconv.Atoi(args[2])
		nome := args[3]
		automa, exists := p.Automi[nome]
		if !exists {
			fmt.Println("L'Automa indicato non esiste.")
			return
		}
		if p.esistePercorso(automa.Posizione, nuoveCoordinate(x, y)) {
			fmt.Println("SI")
		} else {
			fmt.Println("NO")
		}

	case "f":
		os.Exit(0)
	}
}
