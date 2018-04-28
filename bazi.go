package BaziGo

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/JieQi"
	"github.com/warrially/BaziGo/LiChun"
	"github.com/warrially/BaziGo/SiZhu"
	"log"
)

// 八字
type TBazi struct {
	SolarDate TDate   // 新历日期
	BaziDate  TDate   // 八字日期
	SiZhu     TSiZhu  // 四柱
	XiYong    TXiYong // 喜用神
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
	bazi.SiZhu.YearZhu = SiZhu.GetZhuFromYear(bazi.BaziDate.Year)
	// 通过年干支和八字月
	bazi.SiZhu.MonthZhu = SiZhu.GetZhuFromMonth(bazi.BaziDate.Month, bazi.SiZhu.YearZhu.Gan)
	// 通过公历 年月日计算日柱
	bazi.SiZhu.DayZhu = SiZhu.GetZhuFromDay(nYear, nMonth, nDay)
	//
	bazi.SiZhu.HourZhu = SiZhu.GetZhuFromHour(nHour, bazi.SiZhu.DayZhu.Gan)

	// 计算十神
	SiZhu.CalcShiShen(&bazi.SiZhu)
	// 计算纳音
	SiZhu.CalcNaYin(&bazi.SiZhu)
	// 计算喜用神
	bazi.XiYong = SiZhu.CalcXiYong(&bazi.SiZhu)

	return bazi
}

// 从农历获取八字
func GetBaziFromLunar() {

}

// 按照特殊格式化输出(未完成)
func PrintBazi(bazi TBazi) {
	log.Println("======================================================================")
	log.Println("出生日期新历： ", bazi.SolarDate.Year, "年",
		bazi.SolarDate.Month, "月",
		bazi.SolarDate.Day, "日  ",
		bazi.SolarDate.Hour, ":",
		bazi.SolarDate.Minute, ":",
		bazi.SolarDate.Second,
	)
	log.Println("基本八字：",
		bazi.SiZhu.YearZhu.GanZhi.Str,
		bazi.SiZhu.MonthZhu.GanZhi.Str,
		bazi.SiZhu.DayZhu.GanZhi.Str,
		bazi.SiZhu.HourZhu.GanZhi.Str)

	log.Println("命盘解析：")
	log.Println(
		bazi.SiZhu.YearZhu.GanStr, "(",
		bazi.SiZhu.YearZhu.G5XStr, ")[",
		bazi.SiZhu.YearZhu.GSSStr, "]\t",
		bazi.SiZhu.MonthZhu.GanStr, "(",
		bazi.SiZhu.MonthZhu.G5XStr, ")[",
		bazi.SiZhu.MonthZhu.GSSStr, "]\t",
		bazi.SiZhu.DayZhu.GanStr, "(",
		bazi.SiZhu.DayZhu.G5XStr, ")[",
		bazi.SiZhu.DayZhu.GSSStr, "]\t",
		bazi.SiZhu.HourZhu.GanStr, "(",
		bazi.SiZhu.HourZhu.G5XStr, ")[",
		bazi.SiZhu.HourZhu.GSSStr, "]")
	log.Println(
		bazi.SiZhu.YearZhu.ZhiStr, "(",
		bazi.SiZhu.YearZhu.Z5XStr, ")",
		bazi.SiZhu.YearZhu.CangGanSSStr, "\t",
		bazi.SiZhu.MonthZhu.ZhiStr, "(",
		bazi.SiZhu.MonthZhu.Z5XStr, ")",
		bazi.SiZhu.MonthZhu.CangGanSSStr, "\t",
		bazi.SiZhu.DayZhu.ZhiStr, "(",
		bazi.SiZhu.DayZhu.Z5XStr, ")",
		bazi.SiZhu.DayZhu.CangGanSSStr, "\t",
		bazi.SiZhu.HourZhu.ZhiStr, "(",
		bazi.SiZhu.HourZhu.Z5XStr, ")",
		bazi.SiZhu.HourZhu.CangGanSSStr)

	log.Println(
		bazi.SiZhu.YearZhu.NaYinStr, "\t\t",
		bazi.SiZhu.MonthZhu.NaYinStr, "\t\t",
		bazi.SiZhu.DayZhu.NaYinStr, "\t\t",
		bazi.SiZhu.HourZhu.NaYinStr, "\t\t",
	)

	log.Println("======================================================================")

}
