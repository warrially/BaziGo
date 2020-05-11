package bazi

import "fmt"

// NewBazi 新建八字
func NewBazi(pSolarDate *TSolarDate, nSex int) *TBazi {
	//
	pBazi := &TBazi{
		pSolarDate: pSolarDate,
		nSex:       nSex,
	}

	fmt.Println("pSolarDate ", pSolarDate)
	fmt.Println("pBazi", pBazi)
	return pBazi.init()
}

// NewBaziFromLunarDate 新建八字 从农历
func NewBaziFromLunarDate(pLunarDate *TLunarDate, nSex int) *TBazi {
	pBazi := &TBazi{
		pLunarDate: pLunarDate,
		nSex:       nSex,
	}

	return pBazi.init()
}

// GetBazi 旧版八字接口, 八字入口
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) *TBazi {
	// 先解决时间问题. 然后开始处理八字问题
	pSolarDate := NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	fmt.Println(pSolarDate)
	if pSolarDate == nil {
		return nil
	}

	return NewBazi(pSolarDate, nSex)
}

// TBazi 八字大类
type TBazi struct {
	pSolarDate *TSolarDate // 新历的日期
	pLunarDate *TLunarDate // 农历日期
	pBaziDate  *TBaziDate  // 八字历
	pSiZhu     *TSiZhu     // 四柱嗯
	nSex       int         // 性别1男其他女
}

// 八字初始化
func (self *TBazi) init() *TBazi {
	// 1. 新农互转
	if self.pSolarDate == nil {
		if self.pLunarDate == nil {
			return nil
		}

		// todo 这里进行新农互转
		// self.pSolarDate = self.pLunarDate
	} else {
		// todo 这里进行新农互转
		self.pLunarDate = self.pSolarDate.ToLunarDate()
	}

	// 1. 拿到新历的情况下, 需要计算八字历
	self.pBaziDate = self.pSolarDate.ToBaziDate()

	// 2. 根据八字历, 准备计算四柱了
	self.pSiZhu = NewSiZhu(self.pSolarDate, self.pBaziDate)

	// 3.

	return self
}

func (self *TBazi) String() string {
	return fmt.Sprintf("%v\n %v\n %v\n%v", self.pSolarDate, self.pLunarDate, self.pBaziDate, self.pSiZhu)
}
