package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
)

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

func main() {
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

	fmt.Println(masakra)
}
