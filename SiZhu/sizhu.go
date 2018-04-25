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

// 从月 和 年干 获得月柱
func GetZhuFromMonth(nMonth int, nGan int) TZhu {
	var zhu TZhu
	// 根据口诀从本年干数计算本年首月的干数
	log.Println(nMonth)
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
	log.Println(nGan)
	// 计算本月干数
	nGan += ((nMonth - 1) % 10)

	zhu.Gan = nGan % 10
	zhu.Zhi = (nMonth - 1 + 2) % 12
	zhu.GanZhi = CombineGanZhi(zhu.Gan, zhu.Zhi)
	// 转换字符串
	zhu.GanZhiStr = GetGanZhiFromNumber(zhu.GanZhi)
	zhu.GanStr = GetTianGanFromNumber(zhu.Gan)
	zhu.ZhiStr = GetDiZhiFromNumber(zhu.Zhi)
	return zhu
}

// 将天干地支组合成干支，0-9 0-11 转换成 0-59
func CombineGanZhi(nGan, nZhi int) int {
	if (nGan <= 9) && (nZhi <= 11) {
		for i := 0; i <= 6; i++ {
			if (i*10+nGan)%12 == nZhi {
				return i*10 + nGan
			}
		}
	}
	return -1
}
