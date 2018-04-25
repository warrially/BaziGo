package Days

const (
	ctInvalid   = iota //非法，
	ctJulian           //儒略，
	ctGregorian        //格利高里
)

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
	} else { // 公元前
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
