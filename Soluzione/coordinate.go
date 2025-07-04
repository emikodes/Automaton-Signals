//41461A INGENITO EMIDDIO

package main

import "math"

type coordinate struct {
	X, Y int
}

func nuoveCoordinate(x, y int) coordinate {
	return coordinate{X: x, Y: y}
}

func (current coordinate) Distanza(other coordinate) int {
	return int(math.Abs(float64(current.X-other.X)) + math.Abs(float64(current.Y-other.Y))) //somma di due interi, restituisce un intero (non è necessario arrotondamento).
}
