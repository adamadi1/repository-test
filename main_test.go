package main

import (
	"testing"
)

func TestSortuj(t *testing.T) {
	liczby := generujLiczby(ileLiczb)
	posortowane := sortujRosnaco(liczby)

	if len(liczby) == 0 {
		t.Fatalf("funkcja 'generujLiczby' zwróciła zero wyników :( ")
	}

	if len(posortowane) == 0 {
		t.Fatalf("funkcja 'posortowane' zwróciła zero wyników :( ")
	}

	if len(liczby) != len(posortowane) {
		t.Fatalf("Liczb jest %d, ale posortowanych jest %d, a powinno ich być tyle samo", len(liczby), len(posortowane))
	}

	var poprzednia int
	ustawionoPoprzednia := false

	for _, liczba := range posortowane {
		if !ustawionoPoprzednia {
			poprzednia = liczba
			ustawionoPoprzednia = true
			continue
		}

		if liczba < poprzednia {
			t.Fatalf("Liczba %d jest mniejsza od liczby %d", liczba, poprzednia)
			return
		}

		poprzednia = liczba
	}
}

func BenchmarkSortuj(b *testing.B) {
	liczby := generujLiczby(ileLiczb)
	for i := 0; i < b.N; i++ {
		sortujRosnaco(liczby)
	}
}
