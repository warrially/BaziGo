package SiZhu

// 合化冲

import (
	. "github.com/warrially/BaziGo/Common"
)

//
func CheckHeHuaChong(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	CheckTianGanWuHe(pSiZhu, pHeHuaChong)
	CheckDiZhiSanHui(pSiZhu, pHeHuaChong)
	CheckDiZhiSanHe(pSiZhu, pHeHuaChong)
	CheckDiZhiLiuHai(pSiZhu, pHeHuaChong)
	CheckDiZhiLiuChong(pSiZhu, pHeHuaChong)
}

// 天干五合
// 五合是专指天干而言的，地支为六合。
// 命局中干支五行力量的对比与变化，会充分表现在天干的五行上，
// 因此天干五合情况的出现，也会进一步影响或改变命局中的五行力量强弱关系。

// 甲己合化土， 乙庚合化金， 丙辛合化水， 丁壬合化木， 戊癸合化火。
func quickCheckTianGan(nGan1, nGan2 int) (int, string) {
	if nGan1 == nGan2 {
		return -1, "" // 相同不合
	}

	nGan1 %= 5
	nGan2 %= 5

	if nGan1 == nGan2 {
		switch nGan1 {
		case 0:
			return 4, "甲己合化土" // 甲 阳木 + 己 阴土 = 合化土
		case 1:
			return 0, "庚乙合化金" // 庚 阳金 + 乙 阴木 = 合化金
		case 2:
			return 2, "丙辛合化水" // 丙 阳火 + 辛 阴金 = 合化水
		case 3:
			return 1, "壬丁合化木" // 壬 阳水 + 丁 阴火 = 合化木
		case 4:
			return 3, "戊癸合化火" // 戊 阳土 + 癸 阴水 = 合化火
		}
	}

	return -1, ""
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
	if nHe, str := quickCheckTianGan(nTgYear, nTgMonth); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgYear
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = str
		nCount++
	}
	// 检查月日
	if nHe, str := quickCheckTianGan(nTgMonth, nTgDay); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgDay
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = str
		nCount++
	}
	// 检查日时
	if nHe, str := quickCheckTianGan(nTgDay, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgDay
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgHour
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = str
		nCount++
	}

	// 检查年日
	if nHe, str := quickCheckTianGan(nTgYear, nTgDay); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgYear
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = str
		nCount++
	} else
	// 检查月时
	if nHe, str := quickCheckTianGan(nTgMonth, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgMonth
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgDay
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = str
		nCount++
	}
	// 检查年时
	if nHe, str := quickCheckTianGan(nTgYear, nTgHour); nHe >= 0 {
		pHeHuaChong.TgWuHe[nCount].Gan1 = nTgDay
		pHeHuaChong.TgWuHe[nCount].Gan2 = nTgHour
		pHeHuaChong.TgWuHe[nCount].He = nHe
		pHeHuaChong.TgWuHe[nCount].Str = str
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
func quickCheckDiZhiSanHui(nZhi1, nZhi2, nZhi3 int) (int, string) {
	if nZhi1 == nZhi2 || nZhi2 == nZhi3 || nZhi1 == nZhi3 {
		return -1, "" // 相同不会
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
		return 1, "寅卯辰三会木局"
	}

	// 2、巳午未三会火局
	if nZhi1 == 5 && nZhi2 == 6 && nZhi3 == 7 {
		return 3, "巳午未三会火局"
	}

	// 3、申酉戌三会金局
	if nZhi1 == 8 && nZhi2 == 9 && nZhi3 == 10 {
		return 0, "申酉戌三会金局"
	}

	// 4、亥子丑三会水局
	if nZhi1 == 0 && nZhi2 == 1 && nZhi3 == 11 {
		return 2, "亥子丑三会水局"
	}

	return -1, ""
}

// 检查地支三会
func CheckDiZhiSanHui(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	var nDzYear, nDzMonth, nDzDay, nDzHour int // 地支
	nDzYear = pSiZhu.YearZhu.Zhi.Value
	nDzMonth = pSiZhu.MonthZhu.Zhi.Value
	nDzDay = pSiZhu.DayZhu.Zhi.Value
	nDzHour = pSiZhu.HourZhu.Zhi.Value

	var nCount int = 0
	// 年月日
	if nHui, str := quickCheckDiZhiSanHui(nDzYear, nDzMonth, nDzDay); nHui > 0 {
		pHeHuaChong.DzSanHui[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzSanHui[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzSanHui[nCount].Zhi3 = nDzDay
		pHeHuaChong.DzSanHui[nCount].Hui = nHui
		pHeHuaChong.DzSanHui[nCount].Str = str
		nCount++
	}
	// 年月时
	if nHui, str := quickCheckDiZhiSanHui(nDzYear, nDzMonth, nDzHour); nHui > 0 {
		pHeHuaChong.DzSanHui[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzSanHui[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzSanHui[nCount].Zhi3 = nDzHour
		pHeHuaChong.DzSanHui[nCount].Hui = nHui
		pHeHuaChong.DzSanHui[nCount].Str = str
		nCount++
	}
	// 年日时
	if nHui, str := quickCheckDiZhiSanHui(nDzYear, nDzDay, nDzHour); nHui > 0 {
		pHeHuaChong.DzSanHui[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzSanHui[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzSanHui[nCount].Zhi3 = nDzHour
		pHeHuaChong.DzSanHui[nCount].Hui = nHui
		pHeHuaChong.DzSanHui[nCount].Str = str
		nCount++
	}
	// 月日时
	if nHui, str := quickCheckDiZhiSanHui(nDzMonth, nDzDay, nDzHour); nHui > 0 {
		pHeHuaChong.DzSanHui[nCount].Zhi1 = nDzMonth
		pHeHuaChong.DzSanHui[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzSanHui[nCount].Zhi3 = nDzHour
		pHeHuaChong.DzSanHui[nCount].Hui = nHui
		pHeHuaChong.DzSanHui[nCount].Str = str
		nCount++
	}
}

// 地支三合局
// 申子辰合水，申生长刚开始,到子水最帝旺,辰土墓.
// 寅午戌合火，寅生长刚开始,到午火最帝旺,戌土墓.
// 巳酉丑合金，巳生长刚开始,到酉金最帝旺,丑土墓.
// 亥卯未合木，亥生长刚开始,到卯木最帝旺,未土墓
func quickCheckDiZhiSanHe(nZhi1, nZhi2, nZhi3 int) (int, string) {
	if nZhi1 == nZhi2 || nZhi2 == nZhi3 || nZhi1 == nZhi3 {
		return -1, "" // 相同不和
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

	// 申子辰三合水局
	if nZhi1 == 0 && nZhi2 == 4 && nZhi3 == 8 {
		return 2, "申子辰三合水局"
	}

	// 巳酉丑三合金局
	if nZhi1 == 1 && nZhi2 == 5 && nZhi3 == 9 {
		return 0, "巳酉丑三合金局"
	}

	// 寅午戌三合火局
	if nZhi1 == 2 && nZhi2 == 6 && nZhi3 == 10 {
		return 3, "寅午戌三合火局"
	}

	// 亥卯未三合木局
	if nZhi1 == 3 && nZhi2 == 7 && nZhi3 == 11 {
		return 1, "亥卯未三合木局"
	}

	return -1, ""
}

// 检查地支三会
func CheckDiZhiSanHe(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	var nDzYear, nDzMonth, nDzDay, nDzHour int // 地支
	nDzYear = pSiZhu.YearZhu.Zhi.Value
	nDzMonth = pSiZhu.MonthZhu.Zhi.Value
	nDzDay = pSiZhu.DayZhu.Zhi.Value
	nDzHour = pSiZhu.HourZhu.Zhi.Value

	var nCount int = 0
	// 年月日
	if nHe, str := quickCheckDiZhiSanHe(nDzYear, nDzMonth, nDzDay); nHe > 0 {
		pHeHuaChong.DzSanHe[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzSanHe[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzSanHe[nCount].Zhi3 = nDzDay
		pHeHuaChong.DzSanHe[nCount].He = nHe
		pHeHuaChong.DzSanHe[nCount].Str = str
		nCount++
	}
	// 年月时
	if nHe, str := quickCheckDiZhiSanHe(nDzYear, nDzMonth, nDzHour); nHe > 0 {
		pHeHuaChong.DzSanHe[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzSanHe[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzSanHe[nCount].Zhi3 = nDzHour
		pHeHuaChong.DzSanHe[nCount].He = nHe
		pHeHuaChong.DzSanHe[nCount].Str = str
		nCount++
	}
	// 年日时
	if nHe, str := quickCheckDiZhiSanHe(nDzYear, nDzDay, nDzHour); nHe > 0 {
		pHeHuaChong.DzSanHe[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzSanHe[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzSanHe[nCount].Zhi3 = nDzHour
		pHeHuaChong.DzSanHe[nCount].He = nHe
		pHeHuaChong.DzSanHe[nCount].Str = str
		nCount++
	}
	// 月日时
	if nHe, str := quickCheckDiZhiSanHe(nDzMonth, nDzDay, nDzHour); nHe > 0 {
		pHeHuaChong.DzSanHe[nCount].Zhi1 = nDzMonth
		pHeHuaChong.DzSanHe[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzSanHe[nCount].Zhi3 = nDzHour
		pHeHuaChong.DzSanHe[nCount].He = nHe
		pHeHuaChong.DzSanHe[nCount].Str = str
		nCount++
	}
}

// 地支相冲
// 子午相冲
// 丑未相冲
// 寅申相冲
// 卯酉相冲
// 辰戌相冲
// 巳亥相冲
// 地支取七位为冲，犹天干取七位为煞之义。如子午对冲，子至午七数，甲逢庚为煞，甲至庚七数。数中六则合，七则过， 故相冲击为煞也。相冲者，十二支战击之神，大概为凶。
func quickCheckDiZhiLiuChong(nGan1, nGan2 int) (int, string) {
	if nGan1 == nGan2 {
		return -1, "" // 相同不合
	}

	nGan1 %= 6
	nGan2 %= 6

	if nGan1 == nGan2 {
		switch nGan1 {
		case 0:
			return 0, "子午相冲"
		case 1:
			return 0, "丑未相冲"
		case 2:
			return 0, "寅申相冲"
		case 3:
			return 0, "卯酉相冲"
		case 4:
			return 0, "辰戌相冲"
		case 5:
			return 0, "巳亥相冲"
		}
	}

	return -1, ""
}

// 查找地支六冲
func CheckDiZhiLiuChong(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	var nDzYear, nDzMonth, nDzDay, nDzHour int // 地支
	nDzYear = pSiZhu.YearZhu.Zhi.Value
	nDzMonth = pSiZhu.MonthZhu.Zhi.Value
	nDzDay = pSiZhu.DayZhu.Zhi.Value
	nDzHour = pSiZhu.HourZhu.Zhi.Value

	var nCount int = 0

	// 检查年月
	if nChong, str := quickCheckDiZhiLiuChong(nDzYear, nDzMonth); nChong >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
	// 检查月日
	if nChong, str := quickCheckDiZhiLiuChong(nDzMonth, nDzDay); nChong >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
	// 检查日时
	if nChong, str := quickCheckDiZhiLiuChong(nDzDay, nDzHour); nChong >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzHour
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}

	// 检查年日
	if nChong, str := quickCheckDiZhiLiuChong(nDzYear, nDzDay); nChong >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	} else
	// 检查月时
	if nChong, str := quickCheckDiZhiLiuChong(nDzMonth, nDzHour); nChong >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
	// 检查年时
	if nChong, str := quickCheckDiZhiLiuChong(nDzYear, nDzHour); nChong >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzHour
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
}

// 地支六害
// 子未相害
// 丑午相害
// 寅巳相害
// 卯辰相害
// 申亥相害
// 酉戌相害
func quickCheckDiZhiLiuHai(nZhi1, nZhi2 int) (int, string) {
	// 简单冒泡排序
	if nZhi1 > nZhi2 {
		nZhi1, nZhi2 = nZhi2, nZhi1
	}

	if nZhi1 == 0 && nZhi2 == 7 {
		return 0, "子未相害"
	}
	if nZhi1 == 1 && nZhi2 == 6 {
		return 0, "丑午相害"
	}
	if nZhi1 == 2 && nZhi2 == 5 {
		return 0, "寅巳相害"
	}
	if nZhi1 == 3 && nZhi2 == 4 {
		return 0, "卯辰相害"
	}
	if nZhi1 == 8 && nZhi2 == 11 {
		return 0, "申亥相害"
	}
	if nZhi1 == 9 && nZhi2 == 10 {
		return 0, "酉戌相害"
	}

	return -1, ""
}

// 查找地支六害
func CheckDiZhiLiuHai(pSiZhu *TSiZhu, pHeHuaChong *THeHuaChong) {
	var nDzYear, nDzMonth, nDzDay, nDzHour int // 地支
	nDzYear = pSiZhu.YearZhu.Zhi.Value
	nDzMonth = pSiZhu.MonthZhu.Zhi.Value
	nDzDay = pSiZhu.DayZhu.Zhi.Value
	nDzHour = pSiZhu.HourZhu.Zhi.Value

	var nCount int = 0

	// 检查年月
	if nHai, str := quickCheckDiZhiLiuHai(nDzYear, nDzMonth); nHai >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
	// 检查月日
	if nHai, str := quickCheckDiZhiLiuHai(nDzMonth, nDzDay); nHai >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
	// 检查日时
	if nHai, str := quickCheckDiZhiLiuHai(nDzDay, nDzHour); nHai >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzHour
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}

	// 检查年日
	if nHai, str := quickCheckDiZhiLiuHai(nDzYear, nDzDay); nHai >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzYear
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	} else
	// 检查月时
	if nHai, str := quickCheckDiZhiLiuHai(nDzMonth, nDzHour); nHai >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzMonth
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
	// 检查年时
	if nHai, str := quickCheckDiZhiLiuHai(nDzYear, nDzHour); nHai >= 0 {
		pHeHuaChong.DzLiuChong[nCount].Zhi1 = nDzDay
		pHeHuaChong.DzLiuChong[nCount].Zhi2 = nDzHour
		pHeHuaChong.DzLiuChong[nCount].Str = str
		nCount++
	}
}
