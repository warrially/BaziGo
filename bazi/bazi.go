package bazi

// NewBazi 新建八字
func NewBazi(pSolarDate *TSolarDate, nSex int) *TBazi {
	//
	pBazi := &TBazi{
		pSolarDate: pSolarDate,
	}
	
	return pBazi.init()
}

// GetBazi 旧版八字接口, 八字入口
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) *TBazi {
	// 先解决时间问题. 然后开始处理八字问题
	pSolarDate := NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	if pSolarDate == nil {
		return nil
	}

	return NewBazi(pSolarDate, nSex)
}

// TBazi 八字大类
type TBazi struct {
	pSolarDate *TSolarDate // 新历的日期
	pBaziDate *TBaziDate // 八字历
}


// 八字初始化
func (self * TBazi) init() *TBazi{
	// 1. 拿到新历的情况下, 需要计算八字历
	self.pBaziDate = self.pSolarDate.ToBaziDate()

	// 2. 


	return self
}


func (self * TBazi) String() string {
	return self.pSolarDate.String() + "\n"+ self.pBaziDate.String()
}