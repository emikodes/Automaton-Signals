// 41461A INGENITO EMIDDIO
package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	p := newPiano()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		esegui(p, input)
	}
}
