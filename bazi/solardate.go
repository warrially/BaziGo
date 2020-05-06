package bazi

import (
	"log"

	"github.com/warrially/BaziGo/Days"
)

// NewSolarDate 创建一个新历时间
func NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond int) *TSolarDate {
	if !Days.GetDateIsValid(nYear, nMonth, nDay) {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return nil
	}

	// 把具体时间实例化出来
	pDate := &TDate{
		Year:   nYear,   // 年
		Month:  nMonth,  // 月
		Day:    nDay,    // 日
		Hour:   nHour,   // 时
		Minute: nMinute, // 分
		Second: nSecond, // 秒
	}

	return pDate
}

// TSolarDate 日期
type TSolarDate struct {
	Year   int // 年
	Month  int // 月
	Day    int // 日
	Hour   int // 时
	Minute int // 分
	Second int // 秒
}
