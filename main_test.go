package main

import (
	"testing"
)

func TestSortuj(t *testing.T) {
	liczby := generujLiczby(ileLiczb)
	posortowane := sortujRosnaco(liczby)

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
