package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
)

// 计算十神
func CalcNaYin(pSiZhu *TSiZhu) {
	// 从干支转纳音
	pSiZhu.YearZhu.NaYin = GetNaYinFromGanZhi(pSiZhu.YearZhu.GanZhi)
	pSiZhu.MonthZhu.NaYin = GetNaYinFromGanZhi(pSiZhu.MonthZhu.GanZhi)
	pSiZhu.DayZhu.NaYin = GetNaYinFromGanZhi(pSiZhu.DayZhu.GanZhi)
	pSiZhu.HourZhu.NaYin = GetNaYinFromGanZhi(pSiZhu.HourZhu.GanZhi)

	// 纳音转字符串
	pSiZhu.YearZhu.NaYinStr = GetNaYinFromNumber(pSiZhu.YearZhu.NaYin)
	pSiZhu.MonthZhu.NaYinStr = GetNaYinFromNumber(pSiZhu.MonthZhu.NaYin)
	pSiZhu.DayZhu.NaYinStr = GetNaYinFromNumber(pSiZhu.DayZhu.NaYin)
	pSiZhu.HourZhu.NaYinStr = GetNaYinFromNumber(pSiZhu.HourZhu.NaYin)
}

// 纳音刚好是干支的一半, 0 1 -> 0 甲子乙丑海中金
func GetNaYinFromGanZhi(nGanZhi int) int {
	nNaYin := nGanZhi / 2
	return nNaYin
}
