package main

import (
	"github.com/warrially/BaziGo"
	"log"
)

func main() {
	bazi := BaziGo.GetBazi(2018, 4, 25, 19, 7, 0)
	log.Println(bazi)
}
