package main

import (
	"github.com/warrially/BaziGo"
	"log"
)

func main() {
	bazi := BaziGo.GetBazi(2018, 9, 22, 12, 0, 0)
	log.Println(bazi)
}
