package bazi

import "fmt"

// NewBazi 新建八字
func NewBazi(pSolarDate *TSolarDate, nSex int) *TBazi {
	pBazi, _ := NewBaziE(pSolarDate, nSex)
	return pBazi
}

// NewBaziE 新建八字并返回明确错误。
func NewBaziE(pSolarDate *TSolarDate, nSex int) (*TBazi, error) {
	if pSolarDate == nil {
		return nil, ErrInvalidDate
	}
	if nSex != SEX_FEMALE && nSex != SEX_MALE {
		return nil, ErrInvalidSex
	}
	if pSolarDate.Year() < MIN_BAZI_YEAR || pSolarDate.Year() > MAX_BAZI_YEAR {
		return nil, ErrDateOutOfRange
	}
	pBazi := &TBazi{
		pSolarDate: pSolarDate,
		nSex:       nSex,
	}
	return pBazi.initE()
}

// NewBaziFromLunarDate 新建八字 从农历
func NewBaziFromLunarDate(pLunarDate *TLunarDate, nSex int) *TBazi {
	pBazi, _ := NewBaziFromLunarDateE(pLunarDate, nSex)
	return pBazi
}

// NewBaziFromLunarDateE 从农历新建八字并返回明确错误。
func NewBaziFromLunarDateE(pLunarDate *TLunarDate, nSex int) (*TBazi, error) {
	if pLunarDate == nil {
		return nil, ErrInvalidLunarDate
	}
	if nSex != SEX_FEMALE && nSex != SEX_MALE {
		return nil, ErrInvalidSex
	}
	pBazi := &TBazi{
		pLunarDate: pLunarDate,
		nSex:       nSex,
	}

	return pBazi.initE()
}

// GetBazi 旧版八字接口, 八字入口
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) *TBazi {
	pBazi, _ := GetBaziE(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex)
	return pBazi
}

// GetBaziE 带错误信息的安全八字入口。
// GetBazi 为兼容旧代码而保留，新代码应优先使用本函数。
func GetBaziE(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) (*TBazi, error) {
	if err := CheckBaziInput(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex); err != nil {
		return nil, err
	}
	// 先解决时间问题. 然后开始处理八字问题
	pSolarDate := NewSolarDate(nYear, nMonth, nDay, nHour, nMinute, nSecond)
	if pSolarDate == nil {
		return nil, ErrInvalidDate
	}

	return NewBaziE(pSolarDate, nSex)
}

// TBazi 八字大类
type TBazi struct {
	pSolarDate *TSolarDate // 新历的日期
	pLunarDate *TLunarDate // 农历日期
	pBaziDate  *TBaziDate  // 八字历
	pSiZhu     *TSiZhu     // 四柱嗯
	nSex       int         // 性别1男其他女
	pDaYun     *TDaYun     // 大运
	pQiYunDate *TSolarDate // 起运时间XX年XX月开始起运
}

// 八字初始化
func (m *TBazi) init() *TBazi {
	pBazi, _ := m.initE()
	return pBazi
}

func (m *TBazi) initE() (*TBazi, error) {
	// 1. 新农互转
	if m.pSolarDate == nil {
		if m.pLunarDate == nil {
			return nil, ErrInvalidBaziSource
		}
		// 农转新
		m.pSolarDate = m.pLunarDate.ToSolarDate()
	} else {
		// 新转农
		m.pLunarDate = m.pSolarDate.ToLunarDate()
	}
	if m.pSolarDate == nil || m.pLunarDate == nil {
		return nil, ErrInvalidDate
	}

	// 1. 拿到新历的情况下, 需要计算八字历
	m.pBaziDate = m.pSolarDate.ToBaziDate()
	if m.pBaziDate == nil || m.pBaziDate.PreviousJie() == nil || m.pBaziDate.NextJie() == nil {
		return nil, ErrDateOutOfRange
	}

	// 2. 根据八字历, 准备计算四柱了
	m.pSiZhu = NewSiZhu(m.pSolarDate, m.pBaziDate)

	// 3. 计算大运
	m.pDaYun = NewDaYun(m.pSiZhu, m.nSex)

	// 4. 计算起运时间
	m.pQiYunDate = NewQiYun(m.pDaYun.ShunNi(), m.pBaziDate.PreviousJie().ToSolarDate(), m.pBaziDate.NextJie().ToSolarDate(), m.pSolarDate)

	// 5. 起运时间融入到大运中
	nAge := m.QiYunDate().Year() - m.Date().Year()
	for i := 0; i < 12; i++ {
		m.pDaYun.nAge[i] = nAge + 10*i
	}

	return m, nil
}

// String 打印用
func (m *TBazi) String() string {
	return fmt.Sprintf("%v\n %v\n %v\n%v\n%v \n起运时间%v", m.pSolarDate, m.pLunarDate, m.pBaziDate, m.pSiZhu, m.pDaYun, m.pQiYunDate)
}

// SiZhu 四柱
func (m *TBazi) SiZhu() *TSiZhu {
	return m.pSiZhu
}

// Date 获取日期， 默认就是新历
func (m *TBazi) Date() *TSolarDate {
	return m.pSolarDate
}

// SolarData 获取新历日期
func (m *TBazi) SolarData() *TSolarDate {
	return m.Date()
}

// LunarDate 获取农历日期
func (m *TBazi) LunarDate() *TLunarDate {
	return m.pLunarDate
}

// DaYun 获取大运
func (m *TBazi) DaYun() *TDaYun {
	return m.pDaYun
}

// QiYunDate 起运时间
func (m *TBazi) QiYunDate() *TSolarDate {
	return m.pQiYunDate
}
