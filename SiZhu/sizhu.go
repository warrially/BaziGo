package SiZhu

import (
	. "github.com/warrially/BaziGo/Common"
	"log"
)

func wary() {
	log.Println(1)
}

// 从年获得年柱
func GetZhuFromYear(nYear int) TZhu {
	var zhu TZhu
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	zhu.GanZhi = GetGanZhiFromYear(nYear)
	zhu.GanZhiStr = GetGanZhiFromNumber(zhu.GanZhi)
	// 获得八字年的干0-9 对应 甲到癸
	zhu.Gan = GetGanFromYear(nYear)
	zhu.GanStr = GetTianGanFromNumber(zhu.Gan)
	// 获得八字年的支0-11 对应 子到亥
	zhu.Zhi = GetZhiFromYear(nYear)
	zhu.ZhiStr = GetDiZhiFromNumber(zhu.Zhi)
	return zhu
}
