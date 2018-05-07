// 自定义64位时间戳
// 公元1年1月1日到当前时间的秒数

package Days

import (
	. "github.com/warrially/BaziGo/Common"
	"log"
)

// 获取64位时间戳
func Get64TimeStamp(nYear, nMonth, nDay, nHour, nMinute, nSecond int) int64 {
	var nAllDays = GetAllDays(nYear, nMonth, nDay) // 先获取公元原点的日数

	var Result int64 = 0
	Result = int64(nAllDays)
	Result *= 24 * 60 * 60 // 天数换成秒

	//再计算出秒数
	Result += int64(nHour) * 60 * 60
	Result += int64(nMinute) * 60
	Result += int64(nSecond)

	return Result
}

func Get64TimeStampFromDate(dt TDate) int64 {
	return Get64TimeStamp(dt.Year, dt.Month, dt.Day, dt.Hour, dt.Minute, dt.Second)
}

// 从64位时间戳反推日期
func GetDateFrom64TimeStamp(nTimeStamp int64) TDate {
	var dt TDate
	// 计算出年份
	dt.Year = GetYearFrom64TimeStamp(nTimeStamp)
	dt.Month, nTimeStamp = GetMonthFrom64TimeStamp(nTimeStamp, dt.Year)

	dt.Day = int(nTimeStamp / (24 * 60 * 60))
	nTimeStamp -= int64(dt.Day) * 24 * 60 * 60

	dt.Hour = int(nTimeStamp / (60 * 60))
	nTimeStamp -= int64(dt.Hour) * 60 * 60

	dt.Minute = int(nTimeStamp / 60)
	nTimeStamp -= int64(dt.Minute) * 60

	dt.Second = int(nTimeStamp)

	return dt
}

// 从64位时间戳反推年
func GetYearFrom64TimeStamp(nTimeStamp int64) int {
	// 准备进行二分法
	var nLow int = 0
	var nHigh int = 3001

	for {
		var nMid int = (nLow + nHigh) / 2

		// 拿到中间年的数据
		var v int64 = Get64TimeStamp(nMid, 1, 1, 0, 0, 0)

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

// 从64位时间戳反推月, 返回月份和剩余时间戳
func GetMonthFrom64TimeStamp(nTimeStamp int64, nYear int) (int, int64) {
	var v int64 = Get64TimeStamp(nYear, 1, 1, 0, 0, 0)
	log.Println("年份时间戳是 ", v, nYear)
	// 拿到当年的时间戳
	nTimeStamp -= v

	log.Println("减去以后 ", nTimeStamp)

	// 这里开始特殊处理
	for i := 1; i <= 9; i++ {
		if int(nTimeStamp) < GetMonthDays(nYear, i)*24*60*60 {
			return i, nTimeStamp
		}
		nTimeStamp -= int64(GetMonthDays(nYear, i) * 24 * 60 * 60)
	}

	// 10月以后就要扣10天
	for i := 10; i <= 12; i++ {
		if int(nTimeStamp) < (GetMonthDays(nYear, i)-10)*24*60*60 {
			return i, nTimeStamp
		}
		nTimeStamp -= int64(GetMonthDays(nYear, i) * 24 * 60 * 60)
	}

	return 0, nTimeStamp
}
