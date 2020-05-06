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
	pDate := &TSolarDate{
		Year:   nYear,   // 年
		Month:  nMonth,  // 月
		Day:    nDay,    // 日
		Hour:   nHour,   // 时
		Minute: nMinute, // 分
		Second: nSecond, // 秒
	}

	// 计算64位时间戳值
	pDate.Gen64TimeStamp()

	return pDate
}

// NewSolarDateFrom64TimeStamp 从64位时间戳反推日期
func NewSolarDateFrom64TimeStamp(nTimeStamp int64) *TSolarDate {
	pDate := &TSolarDate{}
	pDate.TimeStamp = nTimeStamp

	// 计算出年份
	pDate.Year = pDate.GetYearFrom64TimeStamp(nTimeStamp)

	// 计算月份, 和剩余的时间零头
	pDate.Month, nTimeStamp = Days.GetMonthFrom64TimeStamp(nTimeStamp, pDate.Year)

	// 计算日
	pDate.Day = int(nTimeStamp / (24 * 60 * 60))
	// 扣掉日
	nTimeStamp -= int64(pDate.Day) * 24 * 60 * 60

	pDate.Day++ // 因为每个月的天数是从1开始的, 所以这里需要补1天
	if pDate.Year == 1582 && pDate.Month == 10 && pDate.Day >= 5 {
		pDate.Day += 10 // 1582 年需要补10天
	}
	pDate.Hour = int(nTimeStamp / (60 * 60))
	nTimeStamp -= int64(pDate.Hour) * 60 * 60
	pDate.Minute = int(nTimeStamp / 60)
	nTimeStamp -= int64(pDate.Minute) * 60
	pDate.Second = int(nTimeStamp)

	return pDate
}

// TSolarDate 日期
type TSolarDate struct {
	Year      int   // 年
	Month     int   // 月
	Day       int   // 日
	Hour      int   // 时
	Minute    int   // 分
	Second    int   // 秒
	TimeStamp int64 // 时间戳
}

// Gen64TimeStamp 生成64位时间戳
func (self *TSolarDate) Gen64TimeStamp() int64 {
	nAllDays := Days.GetAllDays(self.Year, self.Month, self.Day) // 先获取公元原点的日数
	nResult := int64(nAllDays)
	nResult *= 24 * 60 * 60 // 天数换成秒

	//再计算出秒数
	nResult += int64(self.Hour) * 60 * 60
	nResult += int64(self.Minute) * 60
	nResult += int64(self.Second)

	self.TimeStamp = nResult
	return nResult
}

// GetYearFrom64TimeStamp 从64位时间戳反推年
func (self *TSolarDate) GetYearFrom64TimeStamp(nTimeStamp int64) int {
	// 准备进行二分法
	nLow := 0
	nHigh := 3001

	for {
		nMid := (nLow + nHigh) / 2

		// 拿到中间年的数据
		v := NewSolarDate(nMid, 1, 1, 0, 0, 0).Gen64TimeStamp()

		if v <= nTimeStamp {
			nLow = nMid
		} else {
			nHigh = nMid
		}

		if nHigh == nLow+1 {
			break
		}
	}
	return nLow
}

// GetMonthFrom64TimeStamp 从64位时间戳反推月, 返回月份和剩余时间戳
func (self *TSolarDate) GetMonthFrom64TimeStamp(nTimeStamp int64, nYear int) (int, int64) {
	v := NewSolarDate(nYear, 1, 1, 0, 0, 0).Gen64TimeStamp()
	// 拿到当年的时间戳
	nTimeStamp -= v

	nOneDaySecond := 24 * 60 * 60 // 1天多少秒

	// 这里开始特殊处理
	for i := 1; i <= 9; i++ {
		if int(nTimeStamp) < self.GetMonthDays(nYear, i)*nOneDaySecond {
			// 前面9个月秒就用完了. 直接返回
			return i, nTimeStamp
		}
		nTimeStamp -= int64(self.GetMonthDays(nYear, i) * nOneDaySecond)
	}

	//  1582 年10月开始需要特殊处理
	if nYear == 1582 {
		if int(nTimeStamp) < 21*nOneDaySecond {
			return 10, nTimeStamp
		}
		nTimeStamp -= int64(21 * nOneDaySecond)

		if int(nTimeStamp) < 30*nOneDaySecond {
			return 11, nTimeStamp
		}
		nTimeStamp -= int64(30 * nOneDaySecond)

		if int(nTimeStamp) < 31*nOneDaySecond {
			return 12, nTimeStamp
		}
		nTimeStamp -= int64(31 * nOneDaySecond)
	} else {
		// 10月以后就要扣10天
		for i := 10; i <= 12; i++ {
			if int(nTimeStamp) < self.GetMonthDays(nYear, i)*nOneDaySecond {
				return i, nTimeStamp
			}
			nTimeStamp -= int64(self.GetMonthDays(nYear, i) * nOneDaySecond)
		}
	}

	return 0, nTimeStamp
}

// GetMonthDays 取本月天数，不考虑 1582 年 10 月的特殊情况
func (self *TSolarDate) GetMonthDays(nYear, nMonth int) int {
	switch nMonth {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2: // 闰年
		if GetIsLeapYear(nYear) {
			return 29
		} else {
			return 28
		}
	}
	return 0
}
