package main

import (
	"math/rand"
	"time"
	"fmt"
	"sort"
)

const ileLiczb = 1000

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
	for i:= 0; i < len(liczby); i++ {
		for j:= 0; j < len(liczby)-i-1; j++{
			if (liczby[j] > liczby[j+1]) {
				liczby[j], liczby[j+1] = liczby[j+1], liczby[j]
			}
			}
		}
	return liczby
}
func Sort(data Interface) {
        for i := 1; i < data.Len(); i++ {
            for j := i; j > 0 && data.Less(j, j-1); j-- {
                data.Swap(j, j-1)
            }
        }
    }
	type Interface interface {
        Len() int
        Less(i, j int) bool
        Swap(i, j int)
    }
func posortowane() {
	liczby:= []int{ileLiczb};
	fmt.Print(sortujRosnaco(liczby))
}
func ints() {
	data := []int{10000}
	a := sort.IntSlice(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("nie działa")
	}




	//sort.Slice(liczby[], func(i, j int) bool {

//}
}
