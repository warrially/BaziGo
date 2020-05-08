package bazi

import "fmt"

// TZhu 柱
type TZhu struct {
	pGanZhi *TGanZhi // 干支
	pGan    *TGan    // 天干
	pZhi    *TZhi    // 地支
}

// NewZhu 新建柱子
func NewZhu() *TZhu {
	return &TZhu{}
}

func (self *TZhu) String() string {
	return fmt.Sprintf("干支:%v 干:%v 支:%v", self.pGanZhi, self.pGan, self.pZhi)
}
