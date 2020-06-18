package bazi

// NewDaYun 新大运
func NewDaYun(pSiZhu *TSiZhu, nSex int) *TDaYun {
	p := &TDaYun{}
	p.init(pSiZhu, nSex)
	return p
}

// TDaYun 大运
type TDaYun struct {
	pZhuList   [12]*TZhu   // 12个大运柱列表
	pQiYunDate *TSolarDate // 起运时间XX年XX月开始起运
	isShunNi   bool        //  顺转还是逆转(true 顺,  false 逆)
}

func (self *TDaYun) init(pSiZhu *TSiZhu, nSex int) *TDaYun {
	for i := 0; i < 12; i++ {
		self.pZhuList[i] = NewZhu() // 新建12个柱
	}

	isYang = pSiZhu.



	return self
}
