package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
)

// 十神
var SHI_SHEN_LIST = [...][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, // 甲
	{1, 0, 3, 2, 5, 4, 7, 6, 9, 8}, // 乙
	{8, 9, 0, 1, 2, 3, 4, 5, 6, 7}, // 丙
	{9, 8, 1, 0, 3, 2, 5, 4, 7, 6}, // 丁
	{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, // 戊
	{7, 6, 9, 8, 1, 0, 3, 2, 5, 4}, // 己
	{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, // 庚
	{5, 4, 7, 6, 9, 8, 1, 0, 3, 2}, // 辛
	{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}, // 壬
	{3, 2, 5, 4, 7, 6, 9, 8, 1, 0}} // 癸

// 从日干和目标干获取十神
// nGan1 日干 nGan2 目标干
func GetShiShenFromGan(nGan1 int, nGan2 int) int {
	if nGan1 < 0 || nGan1 >= 10 || nGan2 < 0 || nGan2 >= 10 {
		return -1
	}
	return SHI_SHEN_LIST[nGan1][nGan2]
}

// 地支藏干表
var DI_ZHI_CANG_GAN_LIST = [12][3]int{
	{9, -1, -1}, // 子水 藏干 癸水。
	{5, 9, 7},   // 丑土 藏干 己土、癸水、辛金。
	{0, 2, 4},   // 寅木 藏干 甲木、丙火、戊土。
	{1, -1, -1}, // 卯木 藏干 乙木。
	{4, 1, 9},   // 辰土 藏干 戊土、乙木、癸水。
	{2, 4, 6},   // 巳火 藏干 丙火、戊土、庚金。
	{3, 5, -1},  // 午火 藏干 丁火、己土。
	{5, 1, 3},   // 未土 藏干 己土、乙木、丁火。
	{6, 4, 8},   // 申金 藏干 庚金、戊土、壬水。
	{7, -1, -1}, // 酉金 藏干 辛金。
	{4, 7, 3},   // 戌土 藏干 戊土、辛金、丁火。
	{8, 0, -1}}  // 亥水 藏干 壬水、甲木。

func GetCangGanFromZhi(nZhi int) [3]int {
	if nZhi < 0 || nZhi >= 12 {
		return [3]int{-1, -1, -1}
	}

	return DI_ZHI_CANG_GAN_LIST[nZhi]
}

// 计算十神
func CalcShiShen(pSiZhu *TSiZhu) {
	// 取出日干十神作为比较
	var nGan = pSiZhu.DayZhu.Gan.Value
	// 天干
	// 年干
	pSiZhu.YearZhu.Gan.ShiShen.Value = GetShiShenFromGan(nGan, pSiZhu.YearZhu.Gan.Value)
	// 月干
	pSiZhu.MonthZhu.Gan.ShiShen.Value = GetShiShenFromGan(nGan, pSiZhu.MonthZhu.Gan.Value)
	// 日干
	pSiZhu.DayZhu.Gan.ShiShen.Value = GetShiShenFromGan(nGan, pSiZhu.DayZhu.Gan.Value)
	// 时干
	pSiZhu.HourZhu.Gan.ShiShen.Value = GetShiShenFromGan(nGan, pSiZhu.HourZhu.Gan.Value)

	// 地支藏干
	pSiZhu.YearZhu.CangGan = GetCangGanFromZhi(pSiZhu.YearZhu.Zhi.Value)
	pSiZhu.MonthZhu.CangGan = GetCangGanFromZhi(pSiZhu.MonthZhu.Zhi.Value)
	pSiZhu.DayZhu.CangGan = GetCangGanFromZhi(pSiZhu.DayZhu.Zhi.Value)
	pSiZhu.HourZhu.CangGan = GetCangGanFromZhi(pSiZhu.HourZhu.Zhi.Value)

	// 地支藏转十神
	// 年
	for i := 0; i < 3; i++ {
		var nCangGan = pSiZhu.YearZhu.CangGan[i]
		if nCangGan >= 0 {
			pSiZhu.YearZhu.CangGanSS[i] = GetShiShenFromGan(nGan, nCangGan)
			pSiZhu.YearZhu.CangGanStr[i] = GetTianGanFromNumber(nCangGan)
			pSiZhu.YearZhu.CangGanSSStr[i] = GetShiShenFromNumber(pSiZhu.YearZhu.CangGanSS[i])
			pSiZhu.YearZhu.CangGan5X[i] = Get5XingFromGan(nCangGan)
			pSiZhu.YearZhu.CangGan5XStr[i] = GetWuXingFromNumber(pSiZhu.YearZhu.CangGan5X[i])
		} else {
			pSiZhu.YearZhu.CangGanSS[i] = -1
		}
	}
	// 月
	for i := 0; i < 3; i++ {
		var nCangGan = pSiZhu.MonthZhu.CangGan[i]
		if nCangGan >= 0 {
			pSiZhu.MonthZhu.CangGanSS[i] = GetShiShenFromGan(nGan, nCangGan)
			pSiZhu.MonthZhu.CangGanStr[i] = GetTianGanFromNumber(nCangGan)
			pSiZhu.MonthZhu.CangGanSSStr[i] = GetShiShenFromNumber(pSiZhu.MonthZhu.CangGanSS[i])
			pSiZhu.MonthZhu.CangGan5X[i] = Get5XingFromGan(nCangGan)
			pSiZhu.MonthZhu.CangGan5XStr[i] = GetWuXingFromNumber(pSiZhu.MonthZhu.CangGan5X[i])
		} else {
			pSiZhu.MonthZhu.CangGanSS[i] = -1
		}
	}
	// 日
	for i := 0; i < 3; i++ {
		var nCangGan = pSiZhu.DayZhu.CangGan[i]
		if nCangGan >= 0 {
			pSiZhu.DayZhu.CangGanSS[i] = GetShiShenFromGan(nGan, nCangGan)
			pSiZhu.DayZhu.CangGanStr[i] = GetTianGanFromNumber(nCangGan)
			pSiZhu.DayZhu.CangGanSSStr[i] = GetShiShenFromNumber(pSiZhu.DayZhu.CangGanSS[i])
			pSiZhu.DayZhu.CangGan5X[i] = Get5XingFromGan(nCangGan)
			pSiZhu.DayZhu.CangGan5XStr[i] = GetWuXingFromNumber(pSiZhu.DayZhu.CangGan5X[i])
		} else {
			pSiZhu.DayZhu.CangGanSS[i] = -1
		}
	}
	// 时
	for i := 0; i < 3; i++ {
		var nCangGan = pSiZhu.HourZhu.CangGan[i]
		if nCangGan >= 0 {
			pSiZhu.HourZhu.CangGanSS[i] = GetShiShenFromGan(nGan, nCangGan)
			pSiZhu.HourZhu.CangGanStr[i] = GetTianGanFromNumber(nCangGan)
			pSiZhu.HourZhu.CangGanSSStr[i] = GetShiShenFromNumber(pSiZhu.HourZhu.CangGanSS[i])
			pSiZhu.HourZhu.CangGan5X[i] = Get5XingFromGan(nCangGan)
			pSiZhu.HourZhu.CangGan5XStr[i] = GetWuXingFromNumber(pSiZhu.HourZhu.CangGan5X[i])
		} else {
			pSiZhu.HourZhu.CangGanSS[i] = -1
		}
	}

	// 转字符串
	pSiZhu.YearZhu.Gan.ShiShen.Str = GetShiShenFromNumber(pSiZhu.YearZhu.Gan.ShiShen.Value)
	pSiZhu.MonthZhu.Gan.ShiShen.Str = GetShiShenFromNumber(pSiZhu.MonthZhu.Gan.ShiShen.Value)
	pSiZhu.DayZhu.Gan.ShiShen.Str = "主"
	pSiZhu.HourZhu.Gan.ShiShen.Str = GetShiShenFromNumber(pSiZhu.HourZhu.Gan.ShiShen.Value)
}
