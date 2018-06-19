package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/Days"
)

// 补充五行
func CalcWuXing(pZhu *TZhu) TZhu {
	// 五行
	Get5XingFromGan2(&pZhu.Gan.WuXing, pZhu.Gan.Value)
	Get5XingFromZhi2(&pZhu.Zhi.WuXing, pZhu.Zhi.Value)
	return *pZhu
}

// 从八字年获得年柱
func GetZhuFromYear(nYear int) TZhu {
	var zhu TZhu
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	GetGanZhiFromYear2(&zhu.GanZhi, nYear)
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	ExtractGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 从八字月 和 年干 获得月柱
func GetZhuFromMonth(nMonth int, nGan int) TZhu {
	nMonth = nMonth % 12
	nGan = nGan % 10
	if nMonth <= 0 {
		nMonth += 12
	}

	if nGan < 0 {
		nGan += 10
	}

	var zhu TZhu
	// 根据口诀从本年干数计算本年首月的干数
	switch nGan {
	case 0, 5:
		// 甲己 丙佐首
		nGan = 2
	case 1, 6:
		// 乙庚 戊为头
		nGan = 4
	case 2, 7:
		// 丙辛 寻庚起
		nGan = 6
	case 3, 8:
		// 丁壬 壬位流
		nGan = 8
	case 4, 9:
		// 戊癸 甲好求
		nGan = 0
	}

	// 计算本月干数
	nGan += ((nMonth - 1) % 10)

	zhu.Gan.Value = nGan % 10
	zhu.Zhi.Value = (nMonth - 1 + 2) % 12
	zhu.Gan.Str = GetTianGanFromNumber(zhu.Gan.Value)
	zhu.Zhi.Str = GetDiZhiFromNumber(zhu.Zhi.Value)

	CombineGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 从公历天 获得日柱
func GetZhuFromDay(nYear int, nMonth int, nDay int) TZhu {
	var zhu TZhu
	// 通过总天数 计算当前天的一个值
	zhu.GanZhi.Value = Days.GetGanZhiFromDay(Days.GetAllDays(nYear, nMonth, nDay))
	zhu.GanZhi.Str = GetGanZhiFromNumber(zhu.GanZhi.Value)
	// 获得八字日的干0-9 对应 甲到癸
	// 获得八字日的支0-11 对应 子到亥
	ExtractGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}

// 从公历小时,  获得日柱天干获取时柱
func GetZhuFromHour(nHour int, nGan int) TZhu {
	var zhu TZhu

	zhu.Gan.Value, zhu.Zhi.Value = Days.GetGanZhiFromHour(nHour, nGan)
	zhu.Gan.Str = GetTianGanFromNumber(zhu.Gan.Value)
	zhu.Zhi.Str = GetDiZhiFromNumber(zhu.Zhi.Value)
	CombineGanZhi2(&zhu.GanZhi, &zhu.Gan, &zhu.Zhi)

	return CalcWuXing(&zhu)
}
