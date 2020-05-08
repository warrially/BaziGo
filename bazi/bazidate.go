package bazi

import (
	"fmt"
)

// NewBaziDate 从新历转成八字历
func NewBaziDate(pSolarDate *TSolarDate) *TBaziDate {
	pBaziDate := &TBaziDate{}
	pBaziDate.nYear = GetLiChunYear(pSolarDate)                           // 拿到八字年, 根据立春来的
	pBaziDate.pPreviousJie, pBaziDate.pNextJie = GetJieQiDate(pSolarDate) // 拿到前后两个的日期
	nJieQi := pBaziDate.pPreviousJie.JieQi
	pBaziDate.pJieQi = &nJieQi
	pBaziDate.nMonth = pBaziDate.pJieQi.ToMonth()

	return pBaziDate
}

// TBaziDate 八字历法
// 八字历法的年  和 新历的 和 农历的都不一样. 八字历法是按照立春为1年. 然后每个节气为月
type TBaziDate struct {
	nYear  int // 年. 立春
	nMonth int // 月.
	nDay   int // 天
	nHour  int // xiaohsi

	pJieQi       *TJieQi     // 节气名称
	pPreviousJie *TJieQiDate // 上一个节(气)
	pNextJie     *TJieQiDate // 下一个节(气)
}

func (self *TBaziDate) String() string {
	return fmt.Sprintf("八字历: %4d 年 %02d 月 \n上一个:%v\n下一个:%v",
		self.nYear, self.nMonth, self.pPreviousJie, self.pNextJie)
}

// Year  年. 立春
func (self *TBaziDate) Year() int {
	return self.nYear
}

// Month  月.
func (self *TBaziDate) Month() int {
	return self.nMonth
}

// Day  天
func (self *TBaziDate) Day() int {
	return self.nDay
}

// Hour 小时
func (self *TBaziDate) Hour() int {
	return self.nHour
}
