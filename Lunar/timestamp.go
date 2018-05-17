package Lunar

import (
	. "github.com/warrially/BaziGo/Common"
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

// 从日期计算时间戳
func Get64TimeStampFromDate(dt TDate) int64 {
	return Get64TimeStamp(dt.Year, dt.Month, dt.Day, dt.Hour, dt.Minute, dt.Second)
}

// 从64位时间戳反推日期
func GetDateFrom64TimeStamp(nTimeStamp int64) TDate {
	var dt TDate
	// 计算年份
	dt.Year = GetYearFrom64TimeStamp(nTimeStamp)
	// 计算月份, 和剩余的时间零头
	dt.Month, nTimeStamp = GetMonthFrom64TimeStamp(nTimeStamp, dt.Year)

	// 计算日
	dt.Day = int(nTimeStamp / (24 * 60 * 60))
	// 扣掉日
	nTimeStamp -= int64(dt.Day) * 24 * 60 * 60

	dt.Day++ // 因为每个月的天数是从1开始的, 所以这里需要补1天
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
	var nLow int = 1900
	var nHigh int = 2100

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
	// 拿到当年的时间戳
	nTimeStamp -= v
	var ONE_DAY_SECOND int = 24 * 60 * 60
	// 这里开始特殊处理
	for i := 1; i <= 13; i++ {
		if int(nTimeStamp) < GetMonthDays(nYear, i)*ONE_DAY_SECOND {
			return i, nTimeStamp
		}
		nTimeStamp -= int64(GetMonthDays(nYear, i) * ONE_DAY_SECOND)
	}

	return 0, nTimeStamp
}
