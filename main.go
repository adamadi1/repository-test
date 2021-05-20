package main

import (
	"strings"
)

var znaki = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "r", "s", "t", "u", "w", "x", "y", "z"}

func Histogram(tekst string) map[string]int {
	wynik := make(map[string]int)

	for _, znak := range znaki {
		iloscWystapien := strings.Count(tekst, znak)
		if iloscWystapien > 0 {
			wynik[znak] = iloscWystapien
		}
	}

	return wynik
}

func ProcentowyHistogram(tekst string) map[string]float32 {
	wynik := make(map[string]float32)
	// oryginalnyHistogram := Histogram(tekst)
	// W TYM MIEJSCU
	// ZADNEGO MAIN

	return wynik
}

// Majac histogram procentowy wykombinowac jak porównać go z histogramami z wikipedii
// dla konkretnego języka
// Porownac każdą literkę w histogramie między tym co mamy, a tym co jest w wikipedii
// i wynik przedstawić jako procentowe podobienstwo

// np. w wikipedii lda jezyka X A=50%, B=30%
// a u nas A=51%, B=29% to jest to całkiem podobne, ale jak bardzo?
