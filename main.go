package main

import (
	"fmt"
	"strings"
)

const Tekst = "aab"

//var znaki = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
//"n", "o", "p", "r", "s", "t", "u", "w", "x", "y", "z"}

func main() {

	txt1 := "testowy tekst oby działało"
	txt2 := "bsdadasdbdadadbdbabbbbadadsbbbada"
	txt3 := ""

	fmt.Println("tekst nr 1: ", txt1)
	fmt.Println("tekst nr 2: ", txt2)
	fmt.Println("tekst nr 3: ", txt3)

	histogram1 := strings.Count(txt1, "a")
	histogram2 := strings.Count(txt2, "b")
	histogram3 := strings.Count(txt3, "" )

	fmt.Println("\nW pierwszym tekscie liter a jest: ", histogram1)
	fmt.Println("W drugim tekscie liter b jest: ", histogram2)
	fmt.Println("Wszytkich znaków jest: ", histogram3)
}

//func Histogram(tekst string) map[string]int {
//wynik := make(map[string]int)
//return wynik
//}
