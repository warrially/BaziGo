package bazi

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
