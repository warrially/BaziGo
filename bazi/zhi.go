package bazi

// 地支

// GetTianZhiFromNumber 从数字获得地支名, 0-9
func GetTianZhiFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "甲"
	case 1:
		return "乙"
	case 2:
		return "丙"
	case 3:
		return "丁"
	case 4:
		return "戊"
	case 5:
		return "己"
	case 6:
		return "庚"
	case 7:
		return "辛"
	case 8:
		return "壬"
	case 9:
		return "癸"
	}

	return ""
}

// NewZhi 创建地支
func NewZhi(nValue int) *TZhi {
	Zhi := TZhi(nValue)
	return &Zhi
}

// TZhi 地支
type TZhi int

// ToString 转换成可阅读的字符串
func (self *TZhi) ToString() string {
	return self.String()
}

// ToInt 转换成int
func (self *TZhi) ToInt() int {
	return self.Value()
}

// ToWuXing 地支转化成五行
func (self *TZhi) ToWuXing() *TWuXing {
	switch self.Value() {
	case 8, 9:
		return NewWuXing(0)
	case 2, 3:
		return NewWuXing(1)
	case 0, 11:
		return NewWuXing(2)
	case 5, 6:
		return NewWuXing(3)
	case 1, 4, 7, 10:
		return NewWuXing(4) // 土
	}
	return nil
}

// ToShiShen 地支转化成食神

// Value 转换成int
func (self *TZhi) Value() int {
	return (int)(*self)
}

// String 转换成可阅读的字符串
func (self *TZhi) String() string {
	return GetTianZhiFromNumber(self.Value())
}
