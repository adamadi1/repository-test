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
	if len(tekst) == 0 {
		return wynik
	}
	histogram := Histogram(tekst)
	dlugoscTekstu := len(tekst)

	for znak, ilosc := range histogram {
		wynik[znak] = float32(ilosc) / float32(dlugoscTekstu)
	}

	return wynik
}

// wziete z wikipedii
var AngielskiHistogram = map[string]float32{
	"a": 0.08167,
	"b": 0.01492,
	"c": 0.02782, // TUTAJ DOPISZ Z WIKIPEDII
}

var PolskiHistogram = map[string]float32{
	"a": 0.08910,
	"b": 0.01535,
	"c": 0.01486,
}



var WszystkieDostepneHistogramy = map[string]map[string]float32{
	"Polski":    PolskiHistogram,
	"Angielski": AngielskiHistogram, // trzeba dodac wiecej histogramow
}

func PorownajZHistogramami(tekst string) string {
	histogram := ProcentowyHistogram(tekst)

	var najmniejszaRoznica float32 = 999999
	var wybranyHistogram string

	for nazwa, dostepnyHistogram := range WszystkieDostepneHistogramy {
		roznica := RoznicaHistogramow(histogram, dostepnyHistogram)
		if roznica < najmniejszaRoznica {
			wybranyHistogram = nazwa
		}
	}

	return wybranyHistogram
}

func RoznicaHistogramow(histogram1, histogram2 map[string]float32) float32 {
	// Tutaj sumujemy różnice między wszystkimi znakami w histogramach i dzielimy przez ilosc znakow
	// TUTAJ ZADANIE DLA ADAMA
	return 0
}
