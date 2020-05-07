package bazi

// NewSiZhu 新四柱
func NewSiZhu(pSolarDate *TSolarDate, pBaziDate *TBaziDate) *TSiZhu {
	pSiZhuDate := &TSiZhu{}
	// pSiZhuDate.pYearZhu =

	// 通过八字年来获取年柱
	// pSiZhuDate.YearZhu = SiZhu.GetZhuFromYear(bazi.BaziDate.Year)
	// 通过年干支和八字月
	// pSiZhuDate.MonthZhu = SiZhu.GetZhuFromMonth(bazi.BaziDate.Month, pSiZhuDate.YearZhu.Gan.Value)
	// 通过公历 年月日计算日柱
	// pSiZhuDate.DayZhu = SiZhu.GetZhuFromDay(bazi.SolarDate.Year, bazi.SolarDate.Month, bazi.SolarDate.Day)
	//
	// pSiZhuDate.HourZhu = SiZhu.GetZhuFromHour(bazi.SolarDate.Hour, pSiZhuDate.DayZhu.Gan.Value)
	return pSiZhuDate
}

// TSiZhu 四柱
type TSiZhu struct {
	pYearZhu  *TZhu // 年柱
	pMonthZhu *TZhu // 月柱
	pDayZhu   *TZhu // 日柱
	pHourZhu  *TZhu // 时柱
}
