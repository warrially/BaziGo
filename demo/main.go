package main

import (
	"github.com/warrially/BaziGo"
	"log"
)

func main() {
	// 年月日时分秒 性别(1男0女)
	bazi := BaziGo.GetBazi(1986, 12, 22, 21, 30, 0, 1)
	BaziGo.PrintBazi(bazi)
	log.Println(bazi)
}
