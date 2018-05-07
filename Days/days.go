package Days

import (
	. "github.com/warrially/BaziGo/Common"
)

const (
	ctInvalid   = iota //非法，
	ctJulian           //儒略，
	ctGregorian        //格利高里
)

// 返回公历日期是否合法
func GetDateIsValid(nYear, nMonth, nDay int) bool {
	// 没有公元0年
	if nYear == 0 {
		return false
	}

	// 1月开始
	if nMonth < 1 {
		return false
	}

	// 12月结束
	if nMonth > 12 {
		return false
	}

	// 1号开始
	if nDay < 1 {
		return false
	}

	// 获取每个月有多少天
	if nDay > GetMonthDays(nYear, nMonth) {
		return false
	}

	// 1582 年的特殊情况
	if nYear != 1582 {
		return true
	}
	if nMonth != 10 {
		return true
	}
	if nDay < 5 {
		return true
	}
	if nDay > 14 {
		return true
	}
	return false
}

// 获得某公历时的天干地支，0-59 对应 甲子到癸亥
func GetGanZhiFromHour(nHour int, nGan int) (int, int) {
	nHour %= 24

	// Zhi是时辰数(0-11)也就是支数
	var nZhi int
	if nHour == 23 {
		// 次日子时
		nGan = (nGan + 1) % 10
		nZhi = 0
	} else {
		nZhi = (nHour + 1) / 2
	}

	// Gan 此时是本日干数，根据规则换算成本日首时辰干数
	if nGan >= 5 {
		nGan -= 5
	}

	// 计算此时辰干数
	nGan = (2*nGan + nZhi) % 10
	return nGan, nZhi
}

// 获得某公历日的天干地支，0-59 对应 甲子到癸亥
func GetGanZhiFromDay(nAllDays int) int {
	return (nAllDays + 12) % 60
}

// 获得距公元原点的日数 这里是公历的年月日
func GetAllDays(nYear int, nMonth int, nDay int) int {
	return GetBasicDays(nYear, nMonth, nDay) + GetLeapDays(nYear, nMonth, nDay)
}

// 获取基本数据
func GetBasicDays(nYear int, nMonth int, nDay int) int {
	var Result int

	// 去掉公元0年
	if nYear > 0 {
		Result = (nYear - 1) * 365
	} else {
		Result = nYear * 365
	}

	// 加上月天数
	for i := 1; i < nMonth; i++ {
		Result += GetMonthDays(nYear, i)
	}

	// 加上日天数
	Result += nDay
	// 返回基础天数
	return Result
}

// 获取闰年天数
func GetLeapDays(nYear int, nMonth int, nDay int) int {
	var Result int

	if nYear >= 0 {
		// 公元后
		if GetCalendarType(nYear, nMonth, nDay) < ctGregorian {
			Result = 0
		} else {
			// 1582.10.5/15 前的 Julian 历只有四年一闰，历法此日后调整为 Gregorian 历
			Result = 10 // 被 Gregory 删去的 10 天

			// 修正算法简化版，从 1701 年的 11 起
			if nYear > 1700 {
				// 每一世纪累加一
				Result += (1 + ((nYear - 1701) / 100))
				// 但 400 整除的世纪不加
				Result -= ((nYear - 1601) / 400)
			}
		}
		Result = ((nYear - 1) / 4) - Result // 4 年一闰数
	} else {
		// 公元前
		Result = -((-nYear + 3) / 4)
	}
	return Result
}

// 根据公历日期判断当时历法
func GetCalendarType(nYear int, nMonth int, nDay int) int {
	if nYear > 1582 {
		return ctGregorian
	} else if nYear < 1582 {
		return ctJulian
	} else if nMonth < 10 {
		return ctJulian
	} else if (nMonth == 10) && (nDay <= 4) {
		return ctJulian
	} else if (nMonth == 10) && (nDay <= 14) {
		return ctInvalid
	} else {
		return ctGregorian
	}
	// 在现在通行的历法记载上，全世界居然有十天没有任何人出生过，也没有任何人死亡过，也没有发生过大大小小值得纪念的人或事。这就是1582年10月5日至10月14日。格里奥，提出了公历历法。这个历法被罗马教皇格里高利十三世采纳了。那么误差的十天怎么办？罗马教皇格里高利十三世下令，把1582年10月4日的后一天改为10月15日，这样误差的十天没有了，历史上也就无影无踪地消失了十天，当然史书上也就没有这十天的记载了。“格里高利公历”一直沿用到今天。
}

// 取本月天数，不考虑 1582 年 10 月的特殊情况
func GetMonthDays(nYear int, nMonth int) int {
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

// 返回某公历是否闰年
func GetIsLeapYear(nYear int) bool {
	if GetCalendarType(nYear, 1, 1) == ctGregorian {
		return (nYear%4 == 0) && ((nYear%100 != 0) || (nYear%400 == 0))
	} else if nYear >= 0 {
		return nYear%4 == 0
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear-3)%4 == 0
	}

}

// 得到两个天数之间的差值
func GetDiffDays(nYear1, nMonth1, nDay1, nYear2, nMonth2, nDay2 int) int {
	var nAllDay1 = GetAllDays(nYear1, nMonth1, nDay1)
	var nAllDay2 = GetAllDays(nYear2, nMonth2, nDay2)

	return nAllDay2 - nAllDay1
}

// 得到两个天数之间的差值(通过时间)
func GetDiffDays2(dt1 TDate, dt2 TDate) int {
	return GetDiffDays(dt1.Year, dt1.Month, dt1.Day, dt2.Year, dt2.Month, dt2.Day)
}

// 得到两个天数之间的秒数差距
func GetDiffSeconds(nYear1, nMonth1, nDay1, nHour1, nMinute1, nSecond1 int,
	nYear2, nMonth2, nDay2, nHour2, nMinute2, nSecond2 int) int64 {

	// 这个数据非常大, 得用64位计算
	var Result int64 = 0

	// // 先计算出天数的差异
	Result = int64(GetDiffDays(nYear1, nMonth1, nDay1, nYear2, nMonth2, nDay2))
	Result *= 24 * 60 * 60 // 天数换成秒

	// // 再计算出秒数的差异
	Result += int64(nHour2-nHour1) * 60 * 60
	Result += int64(nMinute2-nMinute1) * 60
	Result += int64(nSecond2 - nSecond1)

	return Result
}

// 得到两个天数之间的秒数差距(通过日期)
func GetDiffSeconds2(dt1 TDate, dt2 TDate) int64 {
	return GetDiffSeconds(dt1.Year, dt1.Month, dt1.Day, dt1.Hour, dt1.Minute, dt1.Second,
		dt2.Year, dt2.Month, dt2.Day, dt2.Hour, dt2.Minute, dt2.Second)
}

// 根据秒数, 得到新的差异的天数
func GetDiffDate(nYear1, nMonth1, nDay1, nHour1, nMinute1, nSecond1 int, nDiffSecond int64) (int, int, int, int, int, int) {
	//
	var nTime1 int64 = Get64TimeStamp(nYear1, nMonth1, nDay1, nHour1, nMinute1, nSecond1)
	var nTime2 int64 = nTime1 + nDiffSecond

	var dt = GetDateFrom64TimeStamp(nTime2)

	return dt.Year, dt.Month, dt.Day, dt.Hour, dt.Minute, dt.Second
}

// 根据秒数, 得到新的差异天数
func GetDiffDate2(dt1 TDate, nDiffSecond int64) TDate {
	var nTime1 int64 = Get64TimeStampFromDate(dt1)
	var nTime2 int64 = nTime1 + nDiffSecond
	// 从时间戳反推时间
	return GetDateFrom64TimeStamp(nTime2)
}
