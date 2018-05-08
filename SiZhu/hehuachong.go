package SiZhu

// 合化冲

import (
	. "github.com/warrially/BaziGo/Common"
)

// 天干五合
// 五合是专指天干而言的，地支为六合。
// 命局中干支五行力量的对比与变化，会充分表现在天干的五行上，
// 因此天干五合情况的出现，也会进一步影响或改变命局中的五行力量强弱关系。

// 甲己合化土， 乙庚合化金， 丙辛合化水， 丁壬合化木， 戊癸合化火。
func quickCheckTianGan(nGan1, nGan2 int) int {
	if nGan1 == nGan2 {
		return -1 // 相同不合
	}

	nGan1 %= 5
	nGan2 %= 5

	if nGan1 == nGan2 {
		switch nGan1 {
		case 0:
			return 4 // 甲 阳木 + 己 阴土 = 合化土
		case 1:
			return 0 // 庚 阳金 + 乙 阴木 = 合化金
		case 2:
			return 2 // 丙 阳火 + 辛 阴金 = 合化水
		case 3:
			return 1 // 壬 阳水 + 丁 阴火 = 合化木
		case 4:
			return 3 // 戊 阳土 + 癸 阴水 = 合化火
		}
	}

	return -1
}

// 计算合
func GetTianGanHe(nGan1, nGan2, n5x int) string {
	if nGan1%2 == 0 {
		return "天干" + GetTianGanFromNumber(nGan1) + GetTianGanFromNumber(nGan2) + "合" + GetWuXingFromNumber(n5x)
	} else {
		return "天干" + GetTianGanFromNumber(nGan2) + GetTianGanFromNumber(nGan1) + "合" + GetWuXingFromNumber(n5x)
	}
}

// 查找天干五合
func CheckTianGanWuHe(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	// 查找天干五合是否存在
	// var WU_XING_STR = [5]string{
	// "金", "木", "水", "火", "土"}

	var nTgYear, nTgMonth, nTgDay, nTgHour int //
	nTgYear = pSiZhu.YearZhu.Gan.Value
	nTgMonth = pSiZhu.MonthZhu.Gan.Value
	nTgDay = pSiZhu.DayZhu.Gan.Value
	nTgHour = pSiZhu.HourZhu.Gan.Value

	var nCount int = 0

	// 检查年月
	if nHe := quickCheckTianGan(nTgYear, nTgMonth); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgYear
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgYear, nTgMonth, nHe)
		nCount++
	}
	// 检查月日
	if nHe := quickCheckTianGan(nTgMonth, nTgDay); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgDay
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgMonth, nTgDay, nHe)
		nCount++
	}
	// 检查日时
	if nHe := quickCheckTianGan(nTgDay, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgDay
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgHour
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgDay, nTgHour, nHe)
		nCount++
	}

	// 检查年日
	if nHe := quickCheckTianGan(nTgYear, nTgDay); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgYear
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgYear, nTgDay, nHe)
		nCount++
	} else
	// 检查月时
	if nHe := quickCheckTianGan(nTgMonth, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgDay
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgMonth, nTgHour, nHe)
		nCount++
	}
	// 检查年时
	if nHe := quickCheckTianGan(nTgYear, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgDay
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgHour
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = GetTianGanHe(nTgYear, nTgHour, nHe)
		nCount++
	}
}

// 地支三会局
// 1、寅卯辰三会木局
// 在四柱、大运、流年中，只要有寅卯辰三个地支，为寅卯辰三会木局。在寅、卯、辰、未、亥、子六个月令中，三会木局为旺相（称之为三会局旺，下同）；在巳、午、申、酉、戌、丑六个月令中，三会木局为衰弱（称之为三会局衰，下同）。寅卯辰三会木局，表示寅卯木力量增加，辰土的性质不复存在而变化为木。若日主喜木，则代表日主运气好和身体健康的信息；若是日主忌木，则代表日主运气不好，而且肝经、胆经、脾胃经、肾经等有疾。
// 2、巳午未三会火局
// 在四柱、大运、流年中，只要有巳午未三个地支，为巳午未三会火局。在寅、卯、巳、午、未、戌六个月令中，三会火局为旺相；在辰、申、酉、亥、子、丑六个月令中，三会火局为衰弱。巳午未三会火局，表示巳午火力量增加，未土的性质不复存在而变化为火。若日主喜火，则代表日主运气好和身体健康；若日主忌火，则代表日主运气不好，而且肝胆经、肺经、脾胃经、肾经等有病。
// 3、申酉戌三会金局
// 在四柱、大运、流年中，只要有申酉戌三个地支，为申酉戌三会金局。在辰、未、申、酉、戌、丑六个月令中，三会金局为旺相；在寅、卯、巳、午、亥、子六个月令中，三会金局为衰弱。申酉戌三会金局，表示申酉金力量增加，戌土的性质不复存在而变化为金。若日主喜金，则代表日主运气好，而且身体健康；若日主忌金，则代表日主运气不好，而且肝胆经、脾胃经、肺经等有疾。
// 4、亥子丑三会水局
// 在四柱、大运、流年中，只要有亥子丑三个地支，为亥子丑三会水局。在辰、申、酉、亥、子、丑六个月令中，三会水局为旺相；在寅、卯、巳、午、未、戌六个月令中，三会水局为衰弱。亥子丑三会水局，表示亥子水力量增加，丑土的性质不复存在而变化为水。若日主喜水，则代表日主运气好和身体健康的信息；若日主忌水，表示日主运气不好，或心经、肝胆经、脾胃经等有疾。
func quickCheckDiZhiSanHui(nZhi1, nZhi2, nZhi3 int) int {
	if nZhi1 == nZhi2 || nZhi2 == nZhi3 || nZhi1 == nZhi3 {
		return -1 // 相同不会
	}

	// 简单冒泡排序
	if nZhi1 > nZhi2 {
		nZhi1, nZhi2 = nZhi2, nZhi1
	}
	if nZhi2 > nZhi3 {
		nZhi2, nZhi3 = nZhi3, nZhi2
	}
	if nZhi1 > nZhi2 {
		nZhi1, nZhi2 = nZhi2, nZhi1
	}

	// 1、寅卯辰三会木局
	if nZhi1 == 2 && nZhi2 == 3 && nZhi3 == 4 {
		return 1
	}

	// 2、巳午未三会火局
	if nZhi1 == 5 && nZhi2 == 6 && nZhi3 == 7 {
		return 3
	}

	// 3、申酉戌三会金局
	if nZhi1 == 8 && nZhi2 == 9 && nZhi3 == 10 {
		return 0
	}

	// 4、亥子丑三会水局
	if nZhi1 == 0 && nZhi2 == 1 && nZhi3 == 11 {
		return 2
	}
}

// 检查地支三会
func CheckDiZhiSanHui(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	var nDzYear, nDzMonth, nDzDay, nDzHour int // 地支
	nDzYear = pSiZhu.YearZhu.Zhi.Value
	nDzMonth = pSiZhu.MonthZhu.Zhi.Value
	nDzDay = pSiZhu.DayZhu.Zhi.Value
	nDzHour = pSiZhu.HourZhu.Zhi.Value

	// 年月日
	if nHui := quickCheckDiZhiSanHui(nDzYear, nDzMonth, nDzDay); nHui > 0 {

	}
	// 年月时
	if nHui := quickCheckDiZhiSanHui(nDzYear, nDzMonth, nDzHour); nHui > 0 {

	}
	// 年日时
	if nHui := quickCheckDiZhiSanHui(nDzYear, nDzDay, nDzHour); nHui > 0 {

	}
	// 月日时
	if nHui := quickCheckDiZhiSanHui(nDzMonth, nDzDay, nDzHour); nHui > 0 {

	}

}
