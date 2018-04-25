package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
)

// 从年获得年柱
func GetZhuFromYear(nYear int) TZhu {
	var zhu TZhu

	// 获得八字年的干支，0-59 对应 甲子到癸亥
	zhu.GanZhi = GetGanFromYear(nYear)
	zhu.GanZhiStr = GetGanZhiFromNumber(zhu.GanZhi)

	zhu.Gan = GetGanFromYear(nYear)
	zhu.Zhi = GetZhiFromYear(nYear)

	return zhu
}
