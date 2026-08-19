package utils

import (
	"bufio"
	"io"
	"os"
)

func IsTerminal() bool {
	// Zjistí, zda je stdout terminál
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// LineScanner pro čtení po řádcích
func NewLineScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	// Můžeme nastavit buffer pro dlouhé řádky
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	return scanner
}
