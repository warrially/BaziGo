package common

// 天干

// GetTianGanFromNumber 从数字获得天干名, 0-9
func GetTianGanFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "甲"
	case 1:
		return  "乙"
	case 2:
		return  "丙"
	case 3:
		return  "丁"
	case 4:
		return  "戊"
	case 5:
		return 	"己"
	case 6:
		return  "庚"
	case 7:
		return  "辛"
	case 8:
		return  "壬"
	case 9:
		return  "癸"
	}

	return ""
}

// NewGan 创建天干
func NewGan(nValue int) * TGan {
	gan := TGan(nValue)
	return &gan
}

// TGan 天干
type TGan int 

// ToString 转换成可阅读的字符串
func (self *TGan) ToString() string {
	return self.String()
}

// ToInt 转换成int
func (self *TGan) ToInt() int {
	return self.Value()
}

// ToWuXing 天干转化成五行
func (self *TGan) ToWuXing() *TWuXing {
	// todo 
}

// ToShiShen 天干转化成食神


// Value 转换成int
func (self *TGan) Value() int {
	return (int)(*self)
}

// String 转换成可阅读的字符串
func (self* TGan) String() string{
	return GetTianGanFromNumber(self.Value())
}