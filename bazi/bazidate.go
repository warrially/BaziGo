package bazi

import (
	"fmt"
)

// NewBaziDate 从新历转成八字历
func NewBaziDate(pSolarDate *TSolarDate) *TBaziDate {
	pBaziDate := &TBaziDate{}
	pBaziDate.Year = GetLiChunYear(pSolarDate)                          // 拿到八字年, 根据立春来的
	pBaziDate.PreviousJie, pBaziDate.NextJie = GetJieQiDate(pSolarDate) // 拿到前后两个的日期
	nJieQi := pBaziDate.PreviousJie.JieQi
	pBaziDate.JieQi = &nJieQi
	pBaziDate.Month = pBaziDate.JieQi.ToMonth()

	return pBaziDate
}

// TBaziDate 八字历法
// 八字历法的年  和 新历的 和 农历的都不一样. 八字历法是按照立春为1年. 然后每个节气为月
type TBaziDate struct {
	Year  int // 年. 立春
	Month int // 月.
	Day   int // 天
	Hour  int // xiaohsi

	JieQi       *TJieQi     // 节气名称
	PreviousJie *TJieQiDate // 上一个节(气)
	NextJie     *TJieQiDate // 下一个节(气)
}

func (self *TBaziDate) String() string {
	return fmt.Sprintf("八字历:%d年%02d月 节气:%s\n上一个:%v\n下一个:%v",
		self.Year, self.Month, self.JieQi, self.PreviousJie, self.NextJie)
}
