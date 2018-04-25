package BaziGo

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/JieQi"
	"github.com/warrially/BaziGo/LiChun"
	"github.com/warrially/BaziGo/SiZhu"
)

// 八字
type TBazi struct {
	SolarDate TDate  // 新历日期
	BaziDate  TDate  // 八字日期
	Sizhu     TSiZhu // 四柱
}

// 从新历获取八字(年, 月, 日, 时, 分, 秒)
func GetBazi(nYear int, nMonth int, nDay int, nHour int, nMinute int, nSecond int) TBazi {
	var bazi TBazi

	// 新历年
	bazi.SolarDate.Year = nYear
	bazi.SolarDate.Month = nMonth
	bazi.SolarDate.Day = nDay
	bazi.SolarDate.Hour = nHour
	bazi.SolarDate.Minute = nMinute
	bazi.SolarDate.Second = nSecond

	// 通过立春获取当年的年份
	bazi.BaziDate.Year = LiChun.GetLiChun(bazi.SolarDate)
	// 通过节气获取当年的月份
	bazi.BaziDate.JieQi = JieQi.GetJieQi(bazi.SolarDate)
	// 节气0 是立春 是1月
	bazi.BaziDate.Month = bazi.BaziDate.JieQi/2 + 1

	// 通过八字年来获取年柱
	bazi.Sizhu.YearZhu = SiZhu.GetZhuFromYear(bazi.BaziDate.Year)
	// 通过年干支和八字月
	bazi.Sizhu.MonthZhu = SiZhu.GetZhuFromMonth(bazi.BaziDate.Month, bazi.Sizhu.YearZhu.Gan)
	return bazi
}

// 从农历获取八字
func GetBaziFromLunar() {

}
