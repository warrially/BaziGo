package bazi

import (
	"fmt"
	"log"
)

// NewBaziDate 从新历转成八字历
func NewBaziDate(pSolarDate * TSolarDate) *TBaziDate{
	pBaziDate := &TBaziDate{}
	pBaziDate.Year = GetLiChunYear(pSolarDate) // 拿到八字年, 根据立春来的
	date1, date2 := GetJieQiDate(pSolarDate) // 拿到前后两个的日期
	log.Println("前后", date1, date2)
	

	return pBaziDate
}



// TBaziDate 八字历法
// 八字历法的年  和 新历的 和 农历的都不一样. 八字历法是按照立春为1年. 然后每个节气为月
type TBaziDate struct {
	Year int // 年. 立春
	Month int // 月.  节气
}

func (self *TBaziDate) String() string {
	return fmt.Sprintf("八字历:%d年%02d月",
	self.Year,self.Month)
}