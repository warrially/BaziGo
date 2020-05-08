package bazi

import "fmt"

// NewSiZhu 新四柱
func NewSiZhu(pSolarDate *TSolarDate, pBaziDate *TBaziDate) *TSiZhu {
	pSiZhu := &TSiZhu{}
	// 通过八字年来获取年柱
	pSiZhu.genZhuFromYear(pBaziDate.Year())
	// 通过年干支和八字月
	pSiZhu.genZhuFromMonth(pBaziDate.Month())

	// 通过公历 年月日计算日柱
	pSiZhu.genZhuFromDay(pSolarDate.GetAllDays())

	// 通过小时 获取时柱
	pSiZhu.genZhuFromHour(pSolarDate.Hour())

	return pSiZhu
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
	pZhu.pGanZhi = NewGanZhiFromDay(nAllDays)
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

//  genShiShen 计算十神

func (self *TSiZhu) String() string {
	return fmt.Sprintf("四柱:%v %v %v %v\n命盘解析:\n%v(%v)[%v]\t%v(%v)[%v]\t%v(%v)[%v]\t%v(%v)[%v]\t\n%v(%v)    \t%v(%v)    \t%v(%v)    \t%v(%v)\n%v %v %v %v",
		self.pYearZhu.pGanZhi,
		self.pMonthZhu.pGanZhi,
		self.pDayZhu.pGanZhi,
		self.pHourZhu.pGanZhi,
		// ----------------------------------------------------------------------------------
		self.pYearZhu.pGan, self.pYearZhu.pGan.ToWuXing(), self.pYearZhu.pGan.ToShiShen(self.pDayZhu.pGan.Value()),
		self.pMonthZhu.pGan, self.pMonthZhu.pGan.ToWuXing(), self.pMonthZhu.pGan.ToShiShen(self.pDayZhu.pGan.Value()),
		self.pDayZhu.pGan, self.pDayZhu.pGan.ToWuXing(), "主",
		self.pHourZhu.pGan, self.pHourZhu.pGan.ToWuXing(), self.pHourZhu.pGan.ToShiShen(self.pDayZhu.pGan.Value()),
		// ----------------------------------------------------------------------------------
		self.pYearZhu.pZhi, self.pYearZhu.pZhi.ToWuXing(),
		self.pMonthZhu.pZhi, self.pMonthZhu.pZhi.ToWuXing(),
		self.pDayZhu.pZhi, self.pDayZhu.pZhi.ToWuXing(),
		self.pHourZhu.pZhi, self.pHourZhu.pZhi.ToWuXing(),
		// ----------------------------------------------------------------------------------
		self.pYearZhu.pZhi.ToCangGan(self.pDayZhu.pGan.Value()),
		self.pMonthZhu.pZhi.ToCangGan(self.pDayZhu.pGan.Value()),
		self.pDayZhu.pZhi.ToCangGan(self.pDayZhu.pGan.Value()),
		self.pHourZhu.pZhi.ToCangGan(self.pDayZhu.pGan.Value()),
	)
}
