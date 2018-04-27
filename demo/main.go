package main

import (
	"github.com/warrially/BaziGo"
	"log"
)

func main() {
	bazi := BaziGo.GetBazi(1995, 6, 16, 19, 7, 0)
	BaziGo.PrintBazi(bazi)
	log.Println(bazi)
}
