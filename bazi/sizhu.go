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
	pYearZhu   *TZhu       // 年柱
	pMonthZhu  *TZhu       // 月柱
	pDayZhu    *TZhu       // 日柱
	pHourZhu   *TZhu       // 时柱
	pSolarDate *TSolarDate // 新历日期
	pBaziDate  *TBaziDate  // 八字历日期
}

func (self *TSiZhu) init() *TSiZhu {
	// 通过八字年来获取年柱
	self.pYearZhu.genYearGanZhi(self.pBaziDate.Year())
	// 通过年干支和八字月
	self.pMonthZhu.genMonthGanZhi(self.pBaziDate.Month(), self.pYearZhu.Gan().Value())
	// 通过公历 年月日计算日柱
	self.pDayZhu.genDayGanZhi(self.pSolarDate.GetAllDays())
	// 通过小时 获取时柱
	self.pHourZhu.genHourGanZhi(self.pSolarDate.Hour(), self.pDayZhu.Gan().Value())

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
		self.pYearZhu.Zhi().ToCangGan(self.pDayZhu.Gan().Value()),
		self.pMonthZhu.Zhi().ToCangGan(self.pDayZhu.Gan().Value()),
		self.pDayZhu.Zhi().ToCangGan(self.pDayZhu.Gan().Value()),
		self.pHourZhu.Zhi().ToCangGan(self.pDayZhu.Gan().Value()),
		self.pYearZhu.GanZhi().ToNaYin(),
		self.pMonthZhu.GanZhi().ToNaYin(),
		self.pDayZhu.GanZhi().ToNaYin(),
		self.pHourZhu.GanZhi().ToNaYin(),
	)
}
