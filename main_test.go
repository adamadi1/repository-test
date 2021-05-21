package main

import (
	"testing"
)

const tekstPoAngielsku = `Over fact all son tell this any his. No insisted confined of weddings to returned to debating rendered. Keeps order fully so do party means young. Table nay him jokes quick. In felicity up to graceful mistaken horrible consider. Abode never think to at. So additions necessary concluded it happiness do on certainly propriety. On in green taken do offer witty of. 
Satisfied conveying an dependent contented he gentleman agreeable do be. Warrant private blushes removed an in equally totally if. Delivered dejection necessary objection do mr prevailed. Mr feeling do chiefly cordial in do. Water timed folly right aware if oh truth. Imprudence attachment him his for sympathize. Large above be to means. Dashwood do provided stronger is. But discretion frequently sir the she instrument unaffected admiration everything. 
Death there mirth way the noisy merit. Piqued shy spring nor six though mutual living ask extent. Replying of dashwood advanced ladyship smallest disposal or. Attempt offices own improve now see. Called person are around county talked her esteem. Those fully these way nay thing seems. 
From they fine john he give of rich he. They age and draw mrs like. Improving end distrusts may instantly was household applauded incommode. Why kept very ever home mrs. Considered sympathize ten uncommonly occasional assistance sufficient not. Letter of on become he tended active enable to. Vicinity relation sensible sociable surprise screened no up as. 
Inhabit hearing perhaps on ye do no. It maids decay as there he. Smallest on suitable disposed do although blessing he juvenile in. Society or if excited forbade. Here name off yet she long sold easy whom. Differed oh cheerful procured pleasure securing suitable in. Hold rich on an he oh fine. Chapter ability shyness article welcome be do on service. 
So if on advanced addition absolute received replying throwing he. Delighted consisted newspaper of unfeeling as neglected so. Tell size come hard mrs and four fond are. Of in commanded earnestly resources it. At quitting in strictly up wandered of relation answered felicity. Side need at in what dear ever upon if. Same down want joy neat ask pain help she. Alone three stuff use law walls fat asked. Near do that he help. 
With my them if up many. Lain week nay she them her she. Extremity so attending objection as engrossed gentleman something. Instantly gentleman contained belonging exquisite now direction she ham. West room at sent if year. Numerous indulged distance old law you. Total state as merit court green decay he. Steepest sex bachelor the may delicate its yourself. As he instantly on discovery concluded to. Open draw far pure miss felt say yet few sigh. 
Detract yet delight written farther his general. If in so bred at dare rose lose good. Feel and make two real miss use easy. Celebrated delightful an especially increasing instrument am. Indulgence contrasted sufficient to unpleasant in in insensible favourable. Latter remark hunted enough vulgar say man. Sitting hearted on it without me. 
May indulgence difficulty ham can put especially. Bringing remember for supplied her why was confined. Middleton principle did she procuring extensive believing add. Weather adapted prepare oh is calling. These wrong of he which there smile to my front. He fruit oh enjoy it of whose table. Cultivated occasional old her unpleasing unpleasant. At as do be against pasture covered viewing started. Enjoyed me settled mr respect no spirits civilly. 
Tiled say decay spoil now walls meant house. My mr interest thoughts screened of outweigh removing. Evening society musical besides inhabit ye my. Lose hill well up will he over on. Increasing sufficient everything men him admiration unpleasing sex. Around really his use uneasy longer him man. His our pulled nature elinor talked now for excuse result. Admitted add peculiar get joy doubtful. 
`

// func TestPorownajHistogramy(t *testing.T) {
// 	nazwa := PorownajZHistogramami(tekstPoAngielsku)

// 	// TUTAJ ZADANIE DLA ADAMA
// }

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
	h := Histogram(tekstPoAngielsku)

	if len(h) != 24 {
		t.Fatalf("Histogram zawiera tylko %d znaków, a powinien %d", len(h), len(znaki))
	}
}
