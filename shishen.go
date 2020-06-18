package bazi

// GetShiShenFromNumber 从数字获得十神名, 0-9
func GetShiShenFromNumber(nValue int) string {
	switch nValue {
	case 0:
		return "比"
	case 1:
		return "劫"
	case 2:
		return "食"
	case 3:
		return "伤"
	case 4:
		return "才"
	case 5:
		return "财"
	case 6:
		return "杀"
	case 7:
		return "官"
	case 8:
		return "卩"
	case 9:
		return "印"
	}

	return ""
}

// NewShiShen 新建十神
func NewShiShen(nValue int) *TShiShen {
	nValue %= 10
	pShishen := TShiShen(nValue)
	return &pShishen

}

// 十神
var shishenlist = [...][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, // 甲
	{1, 0, 3, 2, 5, 4, 7, 6, 9, 8}, // 乙
	{8, 9, 0, 1, 2, 3, 4, 5, 6, 7}, // 丙
	{9, 8, 1, 0, 3, 2, 5, 4, 7, 6}, // 丁
	{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, // 戊
	{7, 6, 9, 8, 1, 0, 3, 2, 5, 4}, // 己
	{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, // 庚
	{5, 4, 7, 6, 9, 8, 1, 0, 3, 2}, // 辛
	{2, 3, 4, 5, 6, 7, 8, 9, 0, 1}, // 壬
	{3, 2, 5, 4, 7, 6, 9, 8, 1, 0}} // 癸

// GetShiShenFromGan 从日干和目标干获取十神
// nDayGan 日干 Gan2 目标干
func GetShiShenFromGan(nDayGan int, pGan2 *TGan) *TShiShen {
	return NewShiShen(shishenlist[nDayGan][pGan2.Value()])
}

// ToString 转换成可阅读的字符串
func (self *TShiShen) ToString() string {
	return self.String()
}

// ToInt 转换成int
func (self *TShiShen) ToInt() int {
	return self.Value()
}

// TShiShen 十神
type TShiShen int

// Value 转换成int
func (self *TShiShen) Value() int {
	return (int)(*self)
}

// String 转换成可阅读的字符串
func (self *TShiShen) String() string {
	return GetShiShenFromNumber(self.Value())
}
