package bazi

// 起运时间

// NewQiYun 起运时间
func NewQiYun(isShunNi bool, dtPreviousJie *TSolarDate, dtNextJie *TSolarDate, dtSolarDate *TSolarDate) *TSolarDate {
	var nDiffSeconds int64 // 差异的秒数
	if isShunNi {
		// 顺推找下一个节
		nDiffSeconds = dtSolarDate.GetDiffSeconds(dtNextJie)
	} else {
		// 逆推找上一个节
		nDiffSeconds = dtPreviousJie.GetDiffSeconds(dtSolarDate)
	}

	// 大运起运时间的计算方法，
	// 是以出生之日所在月令，按男女顺逆方法推算到下一个节或者上一个节，记下日数。然后按三天为一年，一天为四个月，一个时辰为十天来折算，加上出生时间就是起运的时间。
	nDiffSeconds *= 120

	return NewSolarDateFrom64TimeStamp(dtSolarDate.Get64TimeStamp() + nDiffSeconds)
}
