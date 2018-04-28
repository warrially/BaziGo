package Common

// 日期
type TDate struct {
	Year   int // 年
	Month  int // 月
	Day    int // 日
	Hour   int // 时
	Minute int // 分
	Second int // 秒
	JieQi  int // 节气
}

// 五行属性
type TWuXing struct {
	Value int    // 五行
	Str   string // 五行字符串
}

// 十神属性
type TShiShen struct {
	Value int    // 五行
	Str   string // 五行字符串
}

// 纳音
type TNaYin struct {
	Value int    // 纳音五行
	Str   string // 纳音五行字符串
}

// 干支属性
type TGanZhi struct {
	Value int    // 干支0-59 对应 甲子到癸亥
	Str   string // 干支字符串
	NaYin TNaYin // 纳音
}

// 干属性
type TGan struct {
	Value   int      // 天干
	Str     string   // 天干实际字符串
	WuXing  TWuXing  // 天干五行
	ShiShen TShiShen // 天干十神
}

// 支属性
type TZhi struct {
	Value   int     // 地支
	Str     string  // 地支实际字符串
	WuXing  TWuXing // 地支五行
	CangGan [3]TGan // 藏干
}

// 柱子
type TZhu struct {
	GanZhi TGanZhi // 干支
	Gan    TGan    // 天干
	Zhi    TZhi    // 地支

}

// 四柱
type TSiZhu struct {
	YearZhu  TZhu // 年柱
	MonthZhu TZhu // 月柱
	DayZhu   TZhu // 日柱
	HourZhu  TZhu // 时柱
}

type TXiYong struct {
}

var JIE_QI_STR = [24]string{
	"立春", // 节气  Beginning of Spring   0
	"雨水", // 中气  Rain Water            1
	"惊蛰", // 节气  Waking of Insects     2
	"春分", // 中气  March Equinox         3
	"清明", // 节气  Pure Brightness       4
	"谷雨", // 中气  Grain Rain            5
	"立夏", // 节气  Beginning of Summer   6
	"小满", // 中气  Grain Full            7
	"芒种", // 节气  Grain in Ear          8
	"夏至", // 中气  Summer Solstice       9
	"小暑", // 节气  Slight Heat           10
	"大暑", // 中气  Great Heat            11
	"立秋", // 节气  Beginning of Autumn   12
	"处暑", // 中气  Limit of Heat         13
	"白露", // 节气  White Dew             14
	"秋分", // 中气  September Equinox     15
	"寒露", // 节气  Cold Dew              16
	"霜降", // 中气  Descent of Frost      17
	"立冬", // 节气  Beginning of Winter   18
	"小雪", // 中气  Slight Snow           19
	"大雪", // 节气  Great Snow            20
	"冬至", // 中气  Winter Solstice       21
	"小寒", // 节气  Slight Cold           22，这是一公历年中的第一个节气
	"大寒"} // 中气  Great Cold            23

// 60 干支
var GAN_ZHI_STR = [60]string{
	"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
	"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
	"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
	"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
	"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
	"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥"}

// 从数字获得天干地支名, 0-59
func GetGanZhiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 60) {
		return GAN_ZHI_STR[nValue]
	}
	return "未知"
}

// {* 天干字符串，Heavenly Stems}
var TIAN_GAN_STR = [10]string{
	"甲", "乙", "丙", "丁", "戊",
	"己", "庚", "辛", "壬", "癸"}

// 从数字获得天干名, 0-9
func GetTianGanFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 10) {
		return TIAN_GAN_STR[nValue]
	}
	return ""
}

// {* 地支字符串，Earthly Branches}
var DI_ZHI_STR = [12]string{
	"子", "丑", "寅", "卯",
	"辰", "巳", "午", "未",
	"申", "酉", "戌", "亥"}

// 从数字获得地支名, 0-11
func GetDiZhiFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 12) {
		return DI_ZHI_STR[nValue]
	}
	return ""
}

// {* 五行字符串，以通常的金木水火土为顺序 }
// 这里没用五行相生或者相克来排列
var WU_XING_STR = [5]string{
	"金", "木", "水", "火", "土"}

// 从数字获得五行名, 0-4
func GetWuXingFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 5) {
		return WU_XING_STR[nValue]
	}
	return ""
}

// 十神字符串
var SHI_SHEN_STR = [10]string{
	"比", "劫", "食", "伤", "才",
	"财", "杀", "官", "卩", "印"}

func GetShiShenFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 10) {
		return SHI_SHEN_STR[nValue]
	}
	return ""
}

//  {* 纳音五行，与相邻一对六十干支对应}
// 甲子乙丑海中金丙寅丁卯炉中火戊辰己巳大林木
// 庚午辛未路旁土壬申癸酉剑锋金甲戌乙亥山头火
// 丙子丁丑涧下水戊寅己卯城头土庚辰辛巳白蜡金
// 壬午癸未杨柳木 甲申乙酉井泉水丙戌丁亥屋上土
// 戊子己丑霹雳火庚寅辛卯松柏木壬辰癸巳长流水
// 甲午乙未砂中金丙申丁酉山下火戊戌己亥平地木
// 庚子辛丑壁上土壬寅癸卯金箔金甲辰乙巳覆灯火
// 丙午丁未天河水戊申己酉大驿土庚戌辛亥钗钏金
// 壬子癸丑桑柘木甲寅乙卯大溪水丙辰丁巳砂中土
// 戊午己未天上火庚申辛酉石榴木壬戌癸亥大海水
var NA_YIN_STR = [30]string{
	"海中金", "炉中火", "大林木",
	"路旁土", "剑锋金", "山头火",

	"涧下水", "城墙土", "白蜡金",
	"杨柳木", "泉中水", "屋上土",

	"霹雷火", "松柏木", "长流水",
	"沙中金", "山下火", "平地木",

	"壁上土", "金箔金", "佛灯火",
	"天河水", "大驿土", "钗钏金",

	"桑柘木", "大溪水", "沙中土",
	"天上火", "石榴木", "大海水"}

func GetNaYinFromNumber(nValue int) string {
	if (nValue >= 0) && (nValue < 30) {
		return NA_YIN_STR[nValue]
	}
	return ""
}
