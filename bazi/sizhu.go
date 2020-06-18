package bazi

import "fmt"

// NewSiZhu 新四柱
func NewSiZhu(pSolarDate *TSolarDate, pBaziDate *TBaziDate) *TSiZhu {
	p := &TSiZhu{
		pYearZhu:   NewZhu(),
		pMonthZhu:  NewZhu(),
		pDayZhu:    NewZhu(),
		pHourZhu:   NewZhu(),
		pSolarDate: pSolarDate,
		pBaziDate:  pBaziDate,
	}
	p.init()
	return p
}

// TSiZhu 四柱
type TSiZhu struct {
	pYearZhu    *TZhu        // 年柱
	pMonthZhu   *TZhu        // 月柱
	pDayZhu     *TZhu        // 日柱
	pHourZhu    *TZhu        // 时柱
	pHeHuaChong *THeHuaChong // 荷花冲
	pSolarDate  *TSolarDate  // 新历日期
	pBaziDate   *TBaziDate   // 八字历日期
}

func (self *TSiZhu) init() *TSiZhu {

	// 通过公历 年月日计算日柱
	nDayGan := self.pDayZhu.genDayGanZhi(self.pSolarDate.GetAllDays()).Gan().Value() // 获取日干(日主)
	// 通过小时 获取时柱
	self.pHourZhu.genHourGanZhi(self.pSolarDate.Hour(), nDayGan)
	// 通过八字年来获取年柱
	nYearGan := self.pYearZhu.genYearGanZhi(self.pBaziDate.Year(), nDayGan).Gan().Value()
	// 通过年干支和八字月
	self.pMonthZhu.genMonthGanZhi(self.pBaziDate.Month(), nYearGan, nDayGan)

	return self
}

//  genShiShen 计算十神

func (self *TSiZhu) String() string {
	return fmt.Sprintf("四柱:%v %v %v %v\n命盘解析:\n%v(%v)[%v]\t%v(%v)[%v]\t%v(%v)[%v]\t%v(%v)[%v]\t\n%v(%v)    \t%v(%v)    \t%v(%v)    \t%v(%v)\n%v %v %v %v\n%v %v %v %v",
		self.pYearZhu.GanZhi(),
		self.pMonthZhu.GanZhi(),
		self.pDayZhu.GanZhi(),
		self.pHourZhu.GanZhi(),
		// ----------------------------------------------------------------------------------
		self.pYearZhu.Gan(), self.pYearZhu.Gan().ToWuXing(), self.pYearZhu.Gan().ToShiShen(self.pDayZhu.Gan().Value()),
		self.pMonthZhu.Gan(), self.pMonthZhu.Gan().ToWuXing(), self.pMonthZhu.Gan().ToShiShen(self.pDayZhu.Gan().Value()),
		self.pDayZhu.Gan(), self.pDayZhu.Gan().ToWuXing(), "主",
		self.pHourZhu.Gan(), self.pHourZhu.Gan().ToWuXing(), self.pHourZhu.Gan().ToShiShen(self.pDayZhu.Gan().Value()),
		// ----------------------------------------------------------------------------------
		self.pYearZhu.Zhi(), self.pYearZhu.Zhi().ToWuXing(),
		self.pMonthZhu.Zhi(), self.pMonthZhu.Zhi().ToWuXing(),
		self.pDayZhu.Zhi(), self.pDayZhu.Zhi().ToWuXing(),
		self.pHourZhu.Zhi(), self.pHourZhu.Zhi().ToWuXing(),
		// ----------------------------------------------------------------------------------
		self.pYearZhu.CangGan(),
		self.pMonthZhu.CangGan(),
		self.pDayZhu.CangGan(),
		self.pHourZhu.CangGan(),
		self.pYearZhu.GanZhi().ToNaYin(),
		self.pMonthZhu.GanZhi().ToNaYin(),
		self.pDayZhu.GanZhi().ToNaYin(),
		self.pHourZhu.GanZhi().ToNaYin(),
	)
}

// YearZhu 返回年柱
func (self *TSiZhu) YearZhu() *TZhu {
	return self.pYearZhu
}

// MonthZhu 返回月柱
func (self *TSiZhu) MonthZhu() *TZhu {
	return self.pMonthZhu
}

// DayZhu 返回日柱
func (self *TSiZhu) DayZhu() *TZhu {
	return self.pDayZhu
}

// HourZhu 返回时柱
func (self *TSiZhu) HourZhu() *TZhu {
	return self.pHourZhu
}
