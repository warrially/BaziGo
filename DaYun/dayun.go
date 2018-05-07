// dayun.go
package DaYun

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/Days"
	"github.com/warrially/BaziGo/SiZhu"
)

// 2、大运起排的顺序是：
// 阳年出生的男性、阴年出生的女性，大运是顺行排列；
// 而阴年出生的男性、阳年出生的女性，则是逆行排列。
// 阳年者，是指生辰八字中年柱天干为甲、丙、戊、庚、壬也；
// 而阴年是指生辰八字中年柱天干为乙、丁、己、辛、癸也。
// 大运顺行排列者，即将六十花甲按甲子、乙丑、丙寅、丁卯……庚申、辛酉、壬戌、癸亥之顺序排列也；
// 而大运逆行排列者，即将六十花甲按癸亥、壬戌、辛酉、庚申……丁卯、丙寅、乙丑、甲子之顺序排列也。
// 大运()
func CalcDaYun(pSiZhu *TSiZhu, nSex int) TDaYun {
	var dayun TDaYun
	// 第一判断年柱的阴阳
	var isYang = (GetYinYangFromZhu(&pSiZhu.YearZhu) == 1)
	// 第二判断性别的男女
	var isMale = (nSex == 1)

	// 月柱干支
	var nMonthGanZhi = pSiZhu.MonthZhu.GanZhi.Value
	// 取出日干十神作为比较
	var nDayGan = pSiZhu.DayZhu.Gan.Value

	for i := 0; i < 12; i++ {
		if isYang == isMale {
			//阳年出生的男性、阴年出生的女性
			dayun.Zhu[i].GanZhi.Value = (nMonthGanZhi + 1 + i) % 60
			dayun.ShunNi = true
		} else {
			// 阴年出生的男性、阳年出生的女性
			dayun.Zhu[i].GanZhi.Value = (nMonthGanZhi - 1 - i) % 60
			dayun.ShunNi = false
		}
		// 获取干支名称
		dayun.Zhu[i].GanZhi.Str = GetGanZhiFromNumber(dayun.Zhu[i].GanZhi.Value)
		SiZhu.ExtractGanZhi2(&dayun.Zhu[i].GanZhi, &dayun.Zhu[i].Gan, &dayun.Zhu[i].Zhi)

		// 年干
		SiZhu.GetShiShenFromGan3(&dayun.Zhu[i].Gan, nDayGan)
		// 藏干
		SiZhu.GetCangGanFromZhi2(&dayun.Zhu[i].Zhi)

		// 地支藏转十神
		for j := 0; j < 3; j++ {
			// 年
			SiZhu.CalcCangGan(nDayGan, &dayun.Zhu[i].Zhi.CangGan[j])
		}
	}

	return dayun
}

// 从柱里获取阴阳 (阴 == 0,  阳 == 1)
func GetYinYangFromZhu(pZhu *TZhu) int {
	// 甲丙戊庚壬 0, 2, 4, 6, 8 阳 (1)
	// 乙丁己辛癸 1, 3, 5, 7, 9 阴 (0)
	return (pZhu.Gan.Value + 1) % 2
}

// 计算起运时间
func CalcQiYun(pDaYun *TDaYun, dtPreviousJie TDate, dtNextJie TDate, dtSolarDate TDate) {
	var nDiffSeconds int64 // 差异的秒数
	if pDaYun.ShunNi {
		// 顺推找下一个节
		nDiffSeconds = Days.GetDiffSeconds(dtSolarDate.Year, dtSolarDate.Month, dtSolarDate.Day, dtSolarDate.Hour, dtSolarDate.Minute, dtSolarDate.Second,
			dtNextJie.Year, dtNextJie.Month, dtNextJie.Day, dtNextJie.Hour, dtNextJie.Minute, dtNextJie.Second)
	} else {
		// 逆推找上一个节
		nDiffSeconds = Days.GetDiffSeconds(dtPreviousJie.Year, dtPreviousJie.Month, dtPreviousJie.Day, dtPreviousJie.Hour, dtPreviousJie.Minute, dtPreviousJie.Second,
			dtSolarDate.Year, dtSolarDate.Month, dtSolarDate.Day, dtSolarDate.Hour, dtSolarDate.Minute, dtSolarDate.Second)
	}

	//  大运起运时间的计算方法，是以出生之日所在月令，按男女顺逆方法推算到下一个节或者上一个节，记下日数。然后按三天为一年，一天为四个月，一个时辰为十天来折算，加上出生时间就是起运的时间。
	nDiffSeconds *= 120

	// 通过计算时间来得到
	pDaYun.QiYun = Days.GetDiffDate2(dtSolarDate, nDiffSeconds)
}
