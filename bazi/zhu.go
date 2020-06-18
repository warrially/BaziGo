package bazi

import "fmt"

// TZhu 柱
type TZhu struct {
	pGanZhi  *TGanZhi  // 干支
	pGan     *TGan     // 天干
	pZhi     *TZhi     // 地支
	pCangGan *TCangGan // 藏干
}

// NewZhu 新建柱子
func NewZhu() *TZhu {
	return &TZhu{}
}

func (self *TZhu) String() string {
	return fmt.Sprintf("%v", self.pGanZhi)
}

// 生成年干支
func (self *TZhu) genYearGanZhi(nYear int, nDayGan int) *TZhu {
	// 通过年获取干支
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	self.pGanZhi = NewGanZhiFromYear(nYear)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	self.pGan, self.pZhi = self.pGanZhi.ExtractGanZhi()

	// 在这里计算藏干
	self.pCangGan = NewCangGan(nDayGan, self.pZhi)
	return self
}

//
func (self *TZhu) genMonthGanZhi(nMonth int, nYearGan, nDayGan int) *TZhu {
	// 根据口诀从本年干数计算本年首月的干数
	switch nYearGan {
	case 0, 5:
		// 甲己 丙佐首
		nYearGan = 2
	case 1, 6:
		// 乙庚 戊为头
		nYearGan = 4
	case 2, 7:
		// 丙辛 寻庚起
		nYearGan = 6
	case 3, 8:
		// 丁壬 壬位流
		nYearGan = 8
	case 4, 9:
		// 戊癸 甲好求
		nYearGan = 0
	}

	// 计算本月干数
	nYearGan += ((nMonth - 1) % 10)

	// 拆干
	self.pGan = NewGan(nYearGan % 10)
	self.pZhi = NewZhi((nMonth - 1 + 2) % 12)

	// 组合干支
	self.pGanZhi = CombineGanZhi(self.pGan, self.pZhi)
	// 在这里计算藏干
	self.pCangGan = NewCangGan(nDayGan, self.pZhi)
	return self
}

func (self *TZhu) genDayGanZhi(nAllDays int) *TZhu {

	// 通过总天数来获取
	// 获得八字年的干支，0-59 对应 甲子到癸亥
	self.pGanZhi = NewGanZhiFromDay(nAllDays)
	// 拆分干支
	// 获得八字年的干0-9 对应 甲到癸
	// 获得八字年的支0-11 对应 子到亥
	self.pGan, self.pZhi = self.pGanZhi.ExtractGanZhi()

	// 在这里计算藏干
	self.pCangGan = NewCangGan(self.pGan.Value(), self.pZhi)

	return self
}

func (self *TZhu) genHourGanZhi(nHour, nDayGan int) *TZhu {

	nHour %= 24
	if nHour < 0 {
		nHour += 24
	}

	nZhi := 0
	if nHour == 23 {
		// 次日子时
		nDayGan = (nDayGan + 1) % 10
	} else {
		nZhi = (nHour + 1) / 2
	}

	// Gan 此时是本日干数，根据规则换算成本日首时辰干数
	if nDayGan >= 5 {
		nDayGan -= 5
	}

	// 计算此时辰干数
	nDayGan = (2*nDayGan + nZhi) % 10

	self.pGan = NewGan(nDayGan)
	self.pZhi = NewZhi(nZhi)

	// 组合干支
	self.pGanZhi = CombineGanZhi(self.pGan, self.pZhi)

	// 在这里计算藏干
	self.pCangGan = NewCangGan(nDayGan, self.pZhi)
	return self
}

// Gan 获取干
func (self *TZhu) Gan() *TGan {
	return self.pGan
}

// Zhi 获取支
func (self *TZhu) Zhi() *TZhi {
	return self.pZhi
}

// GanZhi 获取干支
func (self *TZhu) GanZhi() *TGanZhi {
	return self.pGanZhi
}

// ToYinYang 从柱里获取阴阳 (阴 == 0,  阳 == 1)
func (self *TZhu) ToYinYang() int {
	// 甲丙戊庚壬 0, 2, 4, 6, 8 阳 (1)
	// 乙丁己辛癸 1, 3, 5, 7, 9 阴 (0)
	return (self.Gan().Value() + 1) % 2
}

// CangGan 获取藏干
func (self *TZhu) CangGan() *TCangGan {
	return self.pCangGan
}
