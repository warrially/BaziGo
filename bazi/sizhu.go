package bazi

import "fmt"

// NewSiZhu 新四柱
func NewSiZhu(pSolarDate *TSolarDate, pBaziDate *TBaziDate) *TSiZhu {
	pSiZhuDate := &TSiZhu{}
	// 通过八字年来获取年柱
	pSiZhuDate.genZhuFromYear(pBaziDate.Year)
	// 通过年干支和八字月
	pSiZhuDate.genZhuFromMonth(pBaziDate.Month)

	// 通过公历 年月日计算日柱
	pSiZhuDate.genZhuFromDay(pSolarDate.GetAllDays())

	//
	pSiZhuDate.genZhuFromHour(pSolarDate.Hour)
	return pSiZhuDate
}

// TSiZhu 四柱
type TSiZhu struct {
	pYearZhu  *TZhu // 年柱
	pMonthZhu *TZhu // 月柱
	pDayZhu   *TZhu // 日柱
	pHourZhu  *TZhu // 时柱
}

// genZhuFromYear 从八字年获得年柱
func (self *TSiZhu) genZhuFromYear(nYear int) {
	pZhu := NewZhu()

	// 通过年获取干支
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	pZhu.pGanZhi = NewGanZhiFromYear(nYear)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	pZhu.pGan, pZhu.pZhi = pZhu.pGanZhi.ExtractGanZhi()

	self.pYearZhu = pZhu
}

// genZhuFromMonth  从八字月 和 年干 获得月柱
func (self *TSiZhu) genZhuFromMonth(nMonth int) {
	nGan := self.pYearZhu.pGan.Value()

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

	pZhu := NewZhu()

	// 拆干
	pZhu.pGan = NewGan(nGan % 10)
	pZhu.pZhi = NewZhi((nMonth - 1 + 2) % 12)

	// 组合干支
	pZhu.pGanZhi = CombineGanZhi(pZhu.pGan, pZhu.pZhi)

	self.pMonthZhu = pZhu
}

// genZhuFromDay 从总天数获取日柱
func (self *TSiZhu) genZhuFromDay(nAllDays int) {

	pZhu := NewZhu()

	// 通过总天数来获取
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	pZhu.pGanZhi = NewGanZhiFromYear(nAllDays)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	pZhu.pGan, pZhu.pZhi = pZhu.pGanZhi.ExtractGanZhi()

	self.pDayZhu = pZhu
}

// genZhuFromHour 从小事获取时柱
func (self *TSiZhu) genZhuFromHour(nHour int) {
	nGan := self.pDayZhu.pGan.Value()

	nHour %= 24
	if nHour < 0 {
		nHour += 24
	}

	nZhi := 0
	if nHour == 23 {
		// 次日子时
		nGan = (nGan + 1) % 10
	} else {
		nZhi = (nHour + 1) / 2
	}

	// Gan 此时是本日干数，根据规则换算成本日首时辰干数
	if nGan >= 5 {
		nGan -= 5
	}

	// 计算此时辰干数
	nGan = (2*nGan + nZhi) % 10

	pZhu := NewZhu()

	pZhu.pGan = NewGan(nGan)
	pZhu.pZhi = NewZhi(nZhi)

	// 组合干支
	pZhu.pGanZhi = CombineGanZhi(pZhu.pGan, pZhu.pZhi)

	self.pHourZhu = pZhu

}

func (self *TSiZhu) String() string {
	return fmt.Sprintf("年柱:%v\n月柱:%v\n日柱:%v\n时柱:%v\n",
		self.pYearZhu,
		self.pMonthZhu,
		self.pDayZhu,
		self.pHourZhu)
}
