package main

import (
	"math/rand"
	"time"
	"fmt"
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
	for i:= 0; i < len(liczby)-i-1; i++ {
		for j:= 0; j < len(liczby)-i-1; j++{
			if (liczby[j] > liczby[j+1]) {
				liczby[j], liczby[j+1] = liczby[j+1], liczby[j]
			}
		}
	}
	return liczby
}
func posortowane() {
	liczby:= []int{ileLiczb};
	fmt.Print(sortujRosnaco(liczby))


}

