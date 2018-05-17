package Lunar

import (
	. "github.com/warrially/BaziGo/Common"
)


// 获取64位时间戳
func Get64TimeStamp(nYear, nMonth, nDay, nHour, nMinute, nSecond)
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

    return dt
}


// 从64位时间戳反推年
func GetYearFrom64TimeStamp(nTimeStamp int64) int {
    // 准备进行二分法
    var nLow int = 0
    var nHigh int = 3001



}
