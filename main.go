package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
)

// ======	FUNKCJA ZWRACAJĄCA MASAKRĘ ======

//go:embed dane.json
var plikJSON []byte

type Dane struct {
	Kol1 []string `json:"kol1"`
	Kol2 []string `json:"kol2"`
	Kol3 []string `json:"kol3"`
	Kol4 []string `json:"kol4"`
	Kol5 []string `json:"kol5"`
	Kol6 []string `json:"kol6"`
}

func masakra() string {
	var dane Dane

	err := json.Unmarshal(plikJSON, &dane)
	if err != nil {
		log.Fatalf("Błąd dekodowania JSON: %v", err)
	}

	wybor1 := dane.Kol1[rand.IntN(len(dane.Kol1))]
	wybor2 := dane.Kol2[rand.IntN(len(dane.Kol2))]
	wybor3 := dane.Kol3[rand.IntN(len(dane.Kol3))]
	wybor4 := dane.Kol4[rand.IntN(len(dane.Kol4))]
	wybor5 := dane.Kol5[rand.IntN(len(dane.Kol5))]
	wybor6 := dane.Kol6[rand.IntN(len(dane.Kol6))]

	var masakra string
	masakra = wybor1 + wybor2 + wybor3 + wybor4 + wybor5 + wybor6

	return masakra
}

// ========================================

// ====== FUNKCJA DRUKUJĄCA MASAKRĘ =======

func drukujMasakre() {
	masakra := masakra()

	fmt.Println(masakra)
}

// ========================================

// ==== FUNKCJA DRUKUJĄCA WIELE MASAKR ====

func drukujKilka(ilosc int) {
	for i := 1; i == ilosc; i++ {
		masakra := masakra()
		fmt.Printf("%d. %s", i, masakra)
	}
}

// ========================================

func main() {

}
