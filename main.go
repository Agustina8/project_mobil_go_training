package main

import (
	"fmt"
)

type Mobil struct {
	Roda       [4]string
	PintuKanan string
	PintuKiri  string
}

func main() {
	mobil := Mobil{
		Roda:       [4]string{"karet", "karet", "karet", "karet"},
		PintuKanan: "Knock",
		PintuKiri:  "beep",
	}

	fmt.Println("Bunyi pintu kanan ketika diketuk:", mobil.PintuKanan)
	fmt.Println("Bunyi pintu kiri ketika diketuk:", mobil.bunyiPintuKiri())

	RodaBaru := [4]string{"kayu", "besi", "karet", "besi"}
	mobil.gantiRoda(RodaBaru)

	fmt.Println("Bunyi pintu kanan ketika dibuka:", mobil.PintuKanan)
	fmt.Println("Bunyi pintu kiri ketika dibuka:", mobil.bunyiPintuKiri())
}

func (m *Mobil) gantiRoda(RodaBaru [4]string) {
	if len(RodaBaru) == 4 {
		m.Roda = RodaBaru
		fmt.Println("Roda mobil telah diganti.")
	} else {
		fmt.Println("Jumlah Roda tidak sesuai. Harap ganti dengan 4 Roda.")
	}
}

func (m Mobil) bunyiPintuKiri() string {
	if m.PintuKanan == "Knock" {
		return "beep"
	}
	return "Knock"
}
