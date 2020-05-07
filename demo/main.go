package main

import (
	"flag"
	"fmt"

	"github.com/warrially/BaziGo/bazi"
)

func main() {
	var nYear int
	var nMonth int
	var nDay int
	var nHour int
	var nMinute int
	var nSecond int
	var nSex int

	flag.IntVar(&nYear, "y", 1986, "-y=1986 ")
	flag.IntVar(&nMonth, "m", 9, "-m=9 ")
	flag.IntVar(&nDay, "d", 22, "-d=22 ")
	flag.IntVar(&nHour, "h", 12, "-h=12 ")
	flag.IntVar(&nMinute, "n", 0, "-n=0 ")
	flag.IntVar(&nSecond, "s", 0, "-s=0 ")
	flag.IntVar(&nSex, "x", 1, "-x=1  1是男0是女 ")

	flag.Parse() //解析输入的参数

	pBazi := bazi.GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex)
	fmt.Println(pBazi)
}
