package bazi

import (
	"log"

	"github.com/warrially/BaziGo/Days"
)

// NewDate 创建一个新时间
func NewDate(nYear, nMonth, nDay, nHour, nMinute, nSecond int) *TDate {
	if !Days.GetDateIsValid(nYear, nMonth, nDay) {
		log.Println("无效的日期", nYear, nMonth, nDay)
		return nil
	}

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

// TDate 日期
type TDate struct {
	Year   int // 年
	Month  int // 月
	Day    int // 日
	Hour   int // 时
	Minute int // 分
	Second int // 秒
}
