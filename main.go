package main

import (
	"math/rand"
	"time"
)

const ileLiczb = 500000

func main() {
	liczby := generujLiczby(ileLiczb)
	sortujRosnaco(liczby)
}

func generujLiczby(ilosc uint) []int {
	wynik := make([]int, 0)
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := uint(0); i<=ilosc; i++ {
		wynik = append(wynik, gen.Int())
	}

	return wynik
}

func sortujRosnaco(liczby []int) []int {
	posortowane := make([]int, 0)
	for j := 1; j< i; j++{
		if liczby[j-1] > liczby[j]{
			asd := liczby[j]
			liczby[j] = liczby[j-1]
			liczby[j-1] = asd
		}
	
	}

	}
	

 	// @todo zadania dla Adama
	// na potrzeby testu zwracam liczby, ale trzeba zwracac posortowane
	return posortowane
}

