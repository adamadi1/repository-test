package main

import (
	"testing"
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer bibendum efficitur leo, auctor pulvinar massa aliquet auctor. Suspendisse potenti. Morbi volutpat ligula vel tempus venenatis. Sed fermentum dui quis rhoncus bibendum. Suspendisse volutpat blandit nibh, in tincidunt neque lobortis et. Donec eu risus odio. Nam in tortor dictum, hendrerit nulla ut, maximus ex. Donec eu cursus sapien. Aliquam consectetur dictum elit placerat interdum. Morbi rhoncus sed quam eu commodo. Etiam efficitur ante nulla, sed efficitur eros congue quis. Quisque quis odio in tellus faucibus blandit at nec velit. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas condimentum nibh eget pretium dignissim. Donec mollis, orci a efficitur lacinia, elit sem aliquet metus, molestie sagittis mauris arcu nec leo. Donec vel lobortis dolor, ut pellentesque dui.
In semper risus et condimentum tincidunt. Praesent vel turpis id elit mollis finibus faucibus at quam. Nunc arcu nibh, blandit a tincidunt ut, fermentum a urna. Quisque rhoncus laoreet eros ac lacinia. Nulla diam enim, iaculis at vestibulum sed, tempus eu erat. Phasellus ultricies id elit sit amet euismod. Pellentesque sed placerat metus. Proin velit arcu, eleifend id tortor at, sollicitudin condimentum elit. Vivamus dapibus quis arcu eget rutrum. Fusce lobortis sit amet mauris ac auctor.
Nunc laoreet, orci vitae pharetra gravida, purus ante dictum mauris, a congue purus odio nec lectus. Ut at imperdiet velit. Sed finibus sapien odio, vel tristique nulla pulvinar at. Integer sed mollis nisl. Cras efficitur ante sed est tincidunt dapibus. Donec euismod dictum eros sit amet interdum. Nunc dignissim elit eget malesuada tincidunt. Donec sollicitudin leo libero, eget vehicula neque interdum sagittis. Ut dapibus convallis dapibus. Etiam maximus ligula in massa dictum commodo. Maecenas ultricies maximus feugiat. Nunc eu est vitae magna molestie iaculis quis vitae tellus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nam porttitor metus ac dui vulputate aliquet. Vivamus turpis quam, euismod nec malesuada a, finibus et augue. Ut quam mi, pellentesque ac magna ut, porta feugiat ligula.
Aenean dapibus hendrerit ullamcorper. Proin luctus justo at elit accumsan, sit amet scelerisque elit suscipit. Maecenas lacinia arcu vel tempus vulputate. Fusce eget ex varius, placerat mauris volutpat, pretium dui. Sed placerat vitae lacus at porttitor. Mauris ut scelerisque diam, vel vulputate neque. Aenean egestas varius lacus a sollicitudin. Quisque tincidunt, turpis in congue aliquet, quam leo lobortis leo, id fermentum enim felis in lacus. Nulla nec augue euismod, elementum urna quis, vehicula justo. Nulla in ipsum ac odio sollicitudin semper. Sed nisl sem, tristique non sollicitudin id, molestie suscipit dui. Proin nec lobortis ex. Donec imperdiet sem ante, at eleifend nibh aliquam ac. Quisque sed vulputate lectus.
Duis rutrum rutrum arcu, vestibulum porttitor lorem gravida at. Praesent at sollicitudin est. Fusce non risus pulvinar, tristique orci ut, gravida diam. Nunc nec felis a enim ullamcorper eleifend eget vel nisl. Morbi non elit quis augue finibus ultrices ac non enim. Duis id libero aliquet, eleifend tortor a, euismod risus. Ut interdum, leo nec feugiat posuere, tortor neque rhoncus nunc, eget faucibus elit elit commodo purus. Vivamus vestibulum ipsum quis lorem hendrerit, ac luctus odio sagittis. Nullam interdum nibh ut arcu dignissim pulvinar. Phasellus nec lacinia velit. Sed at finibus urna. Sed fermentum sapien rhoncus lacus blandit venenatis. Phasellus porttitor sem nunc, vitae congue ex dignissim eu. Vestibulum gravida, lacus sit amet auctor viverra, elit odio convallis sem, in condimentum ipsum massa at orci.`

func TestHistogramProcentowy_Prosty(t *testing.T) {
	h := ProcentowyHistogram("aaaaabbbbb")

	if len(h) != 2 {
		t.Fatalf("Histogram ma dlugosc %d, a powinien miec 2", len(h))
	}

	histogramA, ok := h["a"]
	if !ok {
		t.Fatalf("Brak literki A w histogramie")
	}

	if histogramA != float32(0.5) {
		t.Fatalf("Histogram literki A powinien wynosic 50%%, a wynosi: %f", histogramA)
	}
}

func TestHistogram_Prosty(t *testing.T) {
	h := Histogram("aaaabbbccd")

	if len(h) != 4 {
		t.Fatalf("Histogram zawiera tylko %d znaków, a powinien %d", len(h), 4)
	}

	czestotliwoscA, ok := h["a"]

	if !ok {
		t.Fatalf("W histogramie brakuje literki A")
		return
	}

	if czestotliwoscA != 4 {
		t.Fatalf("A wystepuje w tekscie 4 razy, a w histogramie tylko %d", h["a"])
	}
}

func TestHistogram_LoremIpsum(t *testing.T) {
	h := Histogram(loremIpsum)

	if len(h) != 20 {
		t.Fatalf("Histogram zawiera tylko %d znaków, a powinien %d", len(h), len(znaki))
	}
}
