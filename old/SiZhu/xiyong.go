package SiZhu

// 喜用神
// 喜用神是中国传统八字命理学上的术语， 喜用神是喜神与用神的合称。
// 八字，即把人出生的年、月、日、时分作四柱，每柱配有一天干和地支，合共八字。
// 八字不同的排列，包含不同的阴阳五行信息，构成各种不同的八字命局。
// 命局中有“不及”和“太过”等情况，称作“病”，而“用神”正是针对不同的“病”所下的“药”。
// “喜神”则是对“用神”能够起到生扶作用的阴阳五行元素。
// 四柱命局以用神为核心， 用神健全有力与否，影响人一生的命；
// 一生补救与否, 影响人一生的运。凡用神之力不足，四柱中有生助用神者，
// 或四柱刑冲克害用神而能化凶神，制凶神者，就是喜神。 四柱没有用神，就得靠行运流年来补。
// 对于命局五行较为平衡，用神不太紧缺的四柱，其一生较为平顺，无大起大落。
import (
	. "github.com/warrially/BaziGo/Common"
	"log"
)

// 计算喜用神
func CalcXiYong(pSiZhu *TSiZhu) TXiYong {
	var xiyong TXiYong

	// 1. 通过四柱计算天干地支强度
	var wuxing = [5]int{0, 0, 0, 0, 0} // 金木水火土

	// 2. 拿到四柱的月支
	var nMonthZhi = pSiZhu.MonthZhu().Zhi().Value()
	// log.Println("月支是", nMonthZhi, pSiZhu.MonthZhu.Zhi.Str)

	// 3. 根据四柱天干, 换算强度
	wuxing[pSiZhu.YearZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.YearZhu.Gan.Value]
	wuxing[pSiZhu.MonthZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.MonthZhu.Gan.Value]
	wuxing[pSiZhu.DayZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.DayZhu.Gan.Value]
	wuxing[pSiZhu.HourZhu.Gan.WuXing.Value] += TIAN_GAN_QIANG_DU_LIST[nMonthZhi][pSiZhu.HourZhu.Gan.Value]

	// log.Println("计算完毕天干后的五行权值是:", wuxing)

	// 4. 根据四柱地支, 换算强度
	for i := 0; i < 3; i++ {
		// 年
		var nCangGan = pSiZhu.YearZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			wuxing[pSiZhu.YearZhu.Zhi.CangGan[i].WuXing.Value] += DI_ZHI_QIANG_DU_LIST[nMonthZhi][pSiZhu.YearZhu.Zhi.CangGan[i].Value]
		}

		// 月
		nCangGan = pSiZhu.MonthZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			wuxing[pSiZhu.MonthZhu.Zhi.CangGan[i].WuXing.Value] += DI_ZHI_QIANG_DU_LIST[nMonthZhi][pSiZhu.MonthZhu.Zhi.CangGan[i].Value]
		}

		// 日
		nCangGan = pSiZhu.DayZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			wuxing[pSiZhu.DayZhu.Zhi.CangGan[i].WuXing.Value] += DI_ZHI_QIANG_DU_LIST[nMonthZhi][pSiZhu.DayZhu.Zhi.CangGan[i].Value]
		}

		// 时
		nCangGan = pSiZhu.HourZhu.Zhi.CangGan[i].Value
		if nCangGan >= 0 {
			wuxing[pSiZhu.HourZhu.Zhi.CangGan[i].WuXing.Value] += DI_ZHI_QIANG_DU_LIST[nMonthZhi][pSiZhu.HourZhu.Zhi.CangGan[i].Value]
		}
	}
	log.Println("计算完毕天干后的五行权值是:", wuxing)

	// 5. 根据日干五行, 计算出同类和异类
	var nDayWuXing = pSiZhu.DayZhu.Gan.WuXing.Value
	xiyong.Same, xiyong.Diff = CalcWuXingQiangRuo(nDayWuXing, wuxing)
	// log.Println("五行同类", xiyong.Same)
	// log.Println("五行异类", xiyong.Diff)
	if xiyong.Same >= xiyong.Diff {
		log.Printf("身强 %d, %.2f%%\n", xiyong.Same-xiyong.Diff, float64(100*xiyong.Same)/float64(xiyong.Diff+xiyong.Same))
	} else {
		log.Printf("身弱 %d, %.2f%%\n", xiyong.Diff-xiyong.Same, float64(100*xiyong.Diff)/float64(xiyong.Diff+xiyong.Same))
	}
	// 月支
	xiyong.MonthZhi = nMonthZhi
	// 日五行
	xiyong.DayWuXing = nDayWuXing

	return xiyong
}

// 计算五行强弱
func CalcWuXingQiangRuo(nDayWuXing int, wuxing [5]int) (int, int) {
	var nSame = 0 // 同类
	var nDiff = 0 // 异类

	// 自己
	nSame += wuxing[nDayWuXing]

	switch nDayWuXing {
	case 0: // 金 同类土
		nSame += wuxing[4]
		nDiff += wuxing[1] + wuxing[2] + wuxing[3]
	case 1: // 木 同类水
		nSame += wuxing[2]
		nDiff += wuxing[0] + wuxing[3] + wuxing[4]
	case 2: // 水 同类金
		nSame += wuxing[0]
		nDiff += wuxing[1] + wuxing[3] + wuxing[4]
	case 3: // 火 同类木
		nSame += wuxing[1]
		nDiff += wuxing[0] + wuxing[2] + wuxing[4]
	case 4: // 土 同类火
		nSame += wuxing[3]
		nDiff += wuxing[0] + wuxing[1] + wuxing[2]
	}
	return nSame, nDiff
}
