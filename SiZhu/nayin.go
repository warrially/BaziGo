package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
)

// 计算十神
func CalcNaYin(pSiZhu *TSiZhu) {
	// 从干支转纳音
	GetNaYinFromGanZhi2(&pSiZhu.YearZhu.GanZhi)
	GetNaYinFromGanZhi2(&pSiZhu.MonthZhu.GanZhi)
	GetNaYinFromGanZhi2(&pSiZhu.DayZhu.GanZhi)
	GetNaYinFromGanZhi2(&pSiZhu.HourZhu.GanZhi)
}

// 纳音刚好是干支的一半, 0 1 -> 0 甲子乙丑海中金
func GetNaYinFromGanZhi(nGanZhi int) int {
	var nNaYin = nGanZhi / 2
	return nNaYin
}

func GetNaYinFromGanZhi2(pGanZhi *TGanZhi) TNaYin {
	pGanZhi.NaYin.Value = GetNaYinFromGanZhi(pGanZhi.Value)
	return pGanZhi.NaYin
}
