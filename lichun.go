package bazi

import "github.com/warrially/BaziGo/internal/jieqidata"

// GetLiChunYear 获取某个日期的立春年。
// 立春数据直接来自统一的嵌入式节气数据，不再维护重复数据表。
func GetLiChunYear(pSolarDate *TSolarDate) int {
	if pSolarDate == nil {
		return 0
	}
	nYear := pSolarDate.Year()
	nMinYear, nMaxYear, err := jieqidata.YearRange()
	if err != nil || nYear < nMinYear || nYear > nMaxYear {
		return nYear
	}

	// 每年 24 条；年初的小寒、大寒之后第 3 条记录为立春。
	nIndex := (nYear-nMinYear)*24 + 2
	nLiChunSeconds, nJieQi, err := jieqidata.Record(nIndex)
	if err != nil || TJieQi(nJieQi) != TJieQi(0) {
		return nYear
	}
	if pSolarDate.Get64TimeStamp() > nLiChunSeconds {
		return nYear
	}
	return nYear - 1
}
