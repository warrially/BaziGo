package bazi

// NewBazi 新建八字
func NewBazi(pDate *TDate, nSex int) *TBazi {
	//
	pBazi := &TBazi{
		pDate: pDate,
	}

	return pBazi
}

// GetBazi 旧版八字接口, 八字入口
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) *TBazi {
	// 先解决时间问题. 然后开始处理八字问题
	pDate := NewDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	if pDate == nil {
		return nil
	}

	return NewBazi(pDate, nSex)
}

// TBazi 八字大类
type TBazi struct {
	pDate *TDate // 八字日期相关
}
