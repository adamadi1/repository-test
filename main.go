package main

import (
	"math"
	"strings"
)

var znaki = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
	"n", "o", "p", "r", "s", "t", "u", "w", "x", "y", "z", "ą", "ę", "ś", "ć", "ź", "ż"}

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
	"c": 0.02782,
	"d": 0.04253,
	"e": 0.12702,
	"f": 0.02228,
	"g": 0.02015,
	"h": 0.06094,
	"i": 0.06966,
	"j": 0.00153,
	"k": 0.00772,
	"l": 0.04025,
	"m": 0.02406,
	"n": 0.06749,
	"o": 0.07507,
	"p": 0.01929,
	"q": 0.00095,
	"r": 0.05987,
	"s": 0.06327,
	"t": 0.09056,
	"u": 0.02758,
	"v": 0.00978,
	"w": 0.02360,
	"x": 0.00150,
	"y": 0.01974,
	"z": 0.00074,
	"ą": 0.00000,
	"ę": 0.00000,
	"ł": 0.00000,
	"ś": 0.00000,
	"ć": 0.00000,
	"ź": 0.00000,
	"ż": 0.00000,
}

var PolskiHistogram = map[string]float32{
	"a": 0.08910,
	"b": 0.01470,
	"c": 0.01242,
	"d": 0.05933,
	"e": 0.18910,
	"f": 0.00300,
	"g": 0.01420,
	"h": 0.01080,
	"i": 0.08210,
	"j": 0.02280,
	"k": 0.03510,
	"l": 0.02100,
	"m": 0.02800,
	"n": 0.05520,
	"o": 0.07750,
	"p": 0.03130,
	"q": 0.00140,
	"r": 0.04690,
	"s": 0.04320,
	"t": 0.03980,
	"u": 0.02500,
	"v": 0.00040,
	"w": 0.04650,
	"x": 0.00020,
	"y": 0.03760,
	"z": 0.05640,
	"ą": 0.00990,
	"ę": 0.01110,
	"ś": 0.00660,
	"ł": 0.01820,
	"ć": 0.00400,
	"ź": 0.00060,
	"ż": 0.00830,
}

var WszystkieDostepneHistogramy = map[string]map[string]float32{
	"Polski":    PolskiHistogram,
	"Angielski": AngielskiHistogram,
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
	sumaRoznic := float32(0)
	for znak1, ilosc1 := range histogram1 {
		ilosc2, istnieje := histogram2[znak1]
		if !istnieje {
			sumaRoznic += 1
			continue
		}

		sumaRoznic += float32(math.Abs(float64(ilosc1 - ilosc2)))
	}

	return sumaRoznic / float32(len(histogram1))
